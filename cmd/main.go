package main

import (
	"gunpla-calendar-exporter/internal/generate"
	"gunpla-calendar-exporter/internal/parse"
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
	monthLower := strings.ToLower(now.Month().String())
	schedule, err := parse.Schedule(baseUrl + strconv.Itoa(year) + monthLower)
	if err != nil {
		panic(err)
	}
	if err = generate.Ics(monthLower, schedule); err != nil {
		panic(err)
	}
}
