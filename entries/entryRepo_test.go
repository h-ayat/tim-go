package entries

import (
	"os"
	"testing"

	"h-ayat.github.io/tim/cal"
	"h-ayat.github.io/tim/util"
)

func TestEntryCRUD(t *testing.T) {

	allEntries := Read(cal.CurrentDate())
	if len(allEntries) != 0 {
		t.Error("Expected no entries")
	}

	AddNewEntry("test1")
	AddNewEntry("test2")
	AddNewEntry("test3")

	allEntries = Read(cal.CurrentDate())

	if len(allEntries) != 3 {
		t.Error("Expected 3 elements")
	}

}

func TestMain(m *testing.M) {
	os.Setenv(util.BasePathENV, "/tmp/tim/")
	deleteTodayFile()
	exitVal := m.Run()
	deleteTodayFile()
	os.Exit(exitVal)
}

func deleteTodayFile() {
	deleteFile(cal.CurrentDate())
}

func deleteFile(date cal.Date) {
	util.Delete(entriesDatePath(date))
}
