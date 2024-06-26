package main

import (
	"context"
	"flag"
	"fmt"
	"gunpla-calendar-exporter/internal/generate"
	"gunpla-calendar-exporter/internal/parse"
	"gunpla-calendar-exporter/internal/upload"
	"log"
	"strings"
	"time"
)

const (
	resaleUrl     = "https://kaigoshinootakunaburogu.com/gunpla-resale-calendar-%d%s"
	newReleaseUrl = "https://kaigoshinootakunaburogu.com/gunpla-newrelease-calendar-%d%s"
)

var (
	toUpload = flag.Bool("upload", true, "flag")
)

func main() {
	ctx := context.Background()
	flag.Parse()
	now := time.Now()
	// 30日や31日といった月末日付のずれを考慮して20日後を指定する。
	monthLower := strings.ToLower(now.AddDate(0, 0, 20).Month().String())
	// 再販スケジュールが非公開になることになったため、一旦は新製品スケジュールのみを取得する
	schedule, err := parse.Schedule(ctx, fmt.Sprintf(newReleaseUrl, now.Year(), monthLower))
	if err != nil {
		log.Fatal(err)
	}
	path, err := generate.Ics(monthLower, schedule)
	if err != nil {
		log.Fatal(err)
	}
	if *toUpload {
		fmt.Println("GCSへのアップロードを行います。")
		if err = upload.CloudStorage(ctx, *path); err != nil {
			log.Fatal(err)
		}
	}
}
