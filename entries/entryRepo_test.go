package entries

import (
	"fmt"
	"os"
	"testing"

	"h-ayat.github.io/tim/cal"
	"h-ayat.github.io/tim/util"
)

func TestEntryCRUD(t *testing.T) {

	DeleteTodayFile()

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

	DeleteTodayFile()

}

func TestMain(m *testing.M) {
	os.Setenv(util.BasePathENV, "/tmp/tim/")
	fmt.Println("Set up stuff for tests here")
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}
