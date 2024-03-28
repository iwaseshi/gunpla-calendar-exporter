package main

import (
	"gunpla-calendar-exporter/internal/generate"
	"gunpla-calendar-exporter/internal/parse"
)

func main() {
	month := "april"
	schedule, err := parse.Schedule("https://kaigoshinootakunaburogu.com/gunpla-resale-calendar-2024" + month)
	if err != nil {
		panic(err)
	}
	_ = generate.Ics(month, schedule)
}
