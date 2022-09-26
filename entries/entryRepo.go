package entries

import (
	"encoding/json"
	"log"
	"time"

	"h-ayat.github.io/tim/cal"
	"h-ayat.github.io/tim/util"
)

type Entry struct {
	Time    string
	Tags    []string
	Message string
}

const EndOfDay = "EOD"

var v = time.Now()

func entriesPath() string {
	return util.BaseAddress() + "entries/"
}
func entriesDatePath(date cal.Date) string {
	return util.BaseAddress() + "entries/" + date.DateAddress()
}

func AddNewEntry(message string) {
	entry := Entry{cal.CurrentTime().Pretty(), []string{}, message}
	date := cal.CurrentDate()

	s, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}

	util.Append(entriesDatePath(date), string(s))
}

func Read(date cal.Date) []Entry {
	lines := util.ReadAll(entriesDatePath(date))
	data := []Entry{}

	for _, l := range lines {
		var e Entry
		err := json.Unmarshal([]byte(l), &e)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, e)
	}
	return data
}

func DeleteTodayFile() {
	DeleteFile(cal.CurrentDate())
}

func DeleteFile(date cal.Date) {
	util.Delete(date.DateAddress())
}
