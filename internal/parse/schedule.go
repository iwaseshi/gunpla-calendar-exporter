package parse

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	datePattern1 = regexp.MustCompile(`(\d{4})\s(\d{1,2})月(\d{1,2})日`)
	datePattern2 = regexp.MustCompile(`(\d{1,2})月(\d{1,2})日`)
)

func Schedule(ctx context.Context, url string) (map[time.Time][]string, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var result = make(map[time.Time][]string)
	fmt.Println(url + " より再販情報を取得します。")
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var outerLIs []string
			err := chromedp.Evaluate(`Array.from(document.querySelectorAll('div.toc-content ol.toc-list.open > li')).map(li => li.innerText)`, &outerLIs).Do(ctx)
			if err != nil {
				return err
			}
			for _, outerLI := range outerLIs {
				lines := strings.Split(outerLI, "\n")
				// 最初の要素は日付の情報が文字列で格納されている
				date, err := convertDate(lines[0])
				if err != nil {
					continue
				}
				result[*date] = lines[1:]
			}
			return nil
		}),
	); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("スケジュールが取得できませんでした。")
	}
	for k, v := range result {
		fmt.Printf("Key: %s \n Value: %v\n\n", k, v)
	}
	return result, nil
}

func convertDate(date string) (*time.Time, error) {
	matches := datePattern1.FindStringSubmatch(date)
	matches2 := datePattern2.FindStringSubmatch(date)
	if matches != nil {
		dateStr := fmt.Sprintf("%s-%02s-%02s", matches[1], matches[2], matches[3])
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("日付の変換に失敗しました:", err)
			return nil, err
		}
		return &date, nil
	} else if matches2 != nil {
		dateStr := fmt.Sprintf("%d-%02s-%02s", time.Now().Year(), matches2[1], matches2[2])
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("日付の変換に失敗しました:", err)
			return nil, err
		}
		return &date, nil
	} else {
		return nil, fmt.Errorf("日付の形式に合致する文字列が見つかりませんでした。")
	}
}
