package cal

import (
	"fmt"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

func (date Date) DateAddress() string {
	return fmt.Sprintf("%d/%d/%d.njson", date.year, date.month, date.day)
}

func (date Date) Ago(days int) Date {
	d := time.Date(date.year, time.Month(date.month), date.day, 0, 0, 0, 0, time.Local)
	return ToDate(d.AddDate(0, 0, -1*days))
}

type Time struct {
	hour   int
	minute int
}

func (t Time) Pretty() string {
	return fmt.Sprintf("%d:%d", t.hour, t.minute)
}

func CurrentDate() Date {
	return ToDate(time.Now())
}

func ToDate(date time.Time) Date {
	return Date{date.Year(), int(date.Month()), date.Day()}
}
func CurrentTime() Time {
	now := time.Now()
	return Time{now.Hour(), now.Minute()}
}
