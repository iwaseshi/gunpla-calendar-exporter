package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	ical "github.com/arran4/golang-ical"
)

func Ics(filename string, source map[time.Time][]string) error {
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
	file, err := os.Create(filepath.Join("./gen", filename+".ics"))
	if err != nil {
		fmt.Println("Error creating ICS file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(cal.Serialize())
	if err != nil {
		fmt.Println("Error writing to ICS file:", err)
		return err
	}
	fmt.Println("ICS file created successfully")
	return nil
}
