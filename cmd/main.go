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
	"sync"
	"time"
)

const (
	resaleUrl     = "https://kaigoshinootakunaburogu.com/gunpla-resale-calendar-%d%s"
	newReleaseUrl = "https://kaigoshinootakunaburogu.com/gunpla-newrelease-calendar-%d%s"
)

var (
	toUpload = flag.Bool("upload", true, "flag")
)

func init() {
	flag.Parse()
}

func main() {
	ctx := context.Background()
	now := time.Now()
	// 30日や31日といった月末日付のずれを考慮して7日後を指定する。
	monthLower := strings.ToLower(now.AddDate(0, 0, 7).Month().String())
	var (
		wg         sync.WaitGroup
		newRelease map[time.Time][]string
		resale     map[time.Time][]string
		err1, err2 error
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		resale, err2 = parse.Schedule(ctx, fmt.Sprintf(resaleUrl, now.Year(), monthLower))
	}()
	go func() {
		defer wg.Done()
		newRelease, err1 = parse.Schedule(ctx, fmt.Sprintf(newReleaseUrl, now.Year(), monthLower))
	}()
	wg.Wait()

	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	path, err := generate.Ics(monthLower, mergeSchedules(newRelease, resale))
	if err != nil {
		log.Fatal(err)
	}
	if !*toUpload {
		return
	}
	fmt.Println("GCSへのアップロードを行います。")
	if err = upload.CloudStorage(ctx, *path); err != nil {
		log.Fatal(err)
	}
}

func mergeSchedules(result1, result2 map[time.Time][]string) map[time.Time][]string {
	mergedResult := make(map[time.Time][]string)

	addToMap := func(date time.Time, items []string) {
		if _, exists := mergedResult[date]; !exists {
			mergedResult[date] = []string{}
		}
		itemSet := make(map[string]bool)
		for _, item := range mergedResult[date] {
			itemSet[item] = true
		}
		for _, item := range items {
			if !itemSet[item] {
				mergedResult[date] = append(mergedResult[date], item)
				itemSet[item] = true
			}
		}
	}

	for date, items := range result1 {
		addToMap(date, items)
	}
	for date, items := range result2 {
		addToMap(date, items)
	}
	for k, v := range mergedResult {
		fmt.Printf("Key: %s \n Value: %v\n\n", k, v)
	}

	return mergedResult
}
