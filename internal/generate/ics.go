package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	ical "github.com/arran4/golang-ical"
)

func Ics(filename string, source map[time.Time][]string) (*string, error) {
	cal := ical.NewCalendar()
	cal.SetMethod(ical.MethodRequest)

	for date, titles := range source {
		for _, title := range titles {
			event := cal.AddEvent(fmt.Sprintf("%s-%d", date.Format("20060102"), time.Now().UnixNano()))
			event.SetStartAt(date)
			event.SetEndAt(date.Add(time.Hour * 1))
			event.SetSummary(title)

		}
	}
	projectRoot := os.Getenv("GUNPLA_CALENDAR_EXPORTER_ROOT")
	if projectRoot == "" {
		return nil, fmt.Errorf("GUNPLA_CALENDAR_EXPORTER_ROOT environment variable is not set.")
	}
	outputoFilePath := filepath.Join(projectRoot, "gen", filename+".ics")
	file, err := os.Create(outputoFilePath)
	if err != nil {
		fmt.Println("Error creating ICS file:", err)
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(cal.Serialize())
	if err != nil {
		fmt.Println("Error writing to ICS file:", err)
		return nil, err
	}
	fmt.Println("ICS file created successfully")
	return &outputoFilePath, nil
}
