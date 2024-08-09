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
	fmt.Println(url + " より発売スケジュールを取得します。")
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
				filterd := filterGunplaSchedules(lines[1:])
				if len(filterd) > 0 {
					result[*date] = filterd
				}
			}
			return nil
		}),
	); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		fmt.Println("スケジュールが取得できませんでした。")
		return result, nil
	}
	fmt.Println(url + " から発売スケジュールが取得できました。")
	return result, nil
}

func convertDate(date string) (*time.Time, error) {
	matches := datePattern1.FindStringSubmatch(date)
	matches2 := datePattern2.FindStringSubmatch(date)
	if matches != nil {
		dateStr := fmt.Sprintf("%s-%02s-%02s", matches[1], matches[2], matches[3])
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("日付の変換に失敗しました: %w", err)
		}
		return &date, nil
	} else if matches2 != nil {
		dateStr := fmt.Sprintf("%d-%02s-%02s", time.Now().Year(), matches2[1], matches2[2])
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("日付の変換に失敗しました: %w", err)
		}
		return &date, nil
	} else {
		return nil, fmt.Errorf("日付の形式に合致する文字列が見つかりませんでした。")
	}
}

func filterGunplaSchedules(input []string) []string {
	var result []string
	for _, str := range input {
		if strings.Contains(str, "HG") || strings.Contains(str, "MG") || strings.Contains(str, "RG") || strings.Contains(str, "EG") {
			result = append(result, str)
		}
	}
	return result
}
