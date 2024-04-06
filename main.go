package main

import (
	"gunpla-calendar-exporter/internal/generate"
	"gunpla-calendar-exporter/internal/parse"
	"gunpla-calendar-exporter/internal/upload"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	baseUrl = "https://kaigoshinootakunaburogu.com/gunpla-resale-calendar-"
)

func main() {
	now := time.Now()
	year := now.Year()
	// 30日や31日といった月末日付のずれを考慮して20日後を指定する。
	monthLower := strings.ToLower(now.AddDate(0, 0, 20).Month().String())
	schedule, err := parse.Schedule(baseUrl + strconv.Itoa(year) + monthLower)
	if err != nil {
		log.Fatal(err)
	}
	path, err := generate.Ics(monthLower, schedule)
	if err != nil {
		log.Fatal(err)
	}
	if err = upload.CloudStorage(*path); err != nil {
		log.Fatal(err)
	}
}
