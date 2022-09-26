package tags

import (
	"fmt"
	"os"
	"testing"

	"h-ayat.github.io/tim/util"
)

func TestTags(t *testing.T) {
	allTags := LoadAllTags()
	if len(allTags) != 0 {
		fmt.Println("--------")
		fmt.Println(allTags)
		t.Error("Expected no tags")
	}
	AddTag("testA")
	AddTag("testB")
	allTags = LoadAllTags()

	if len(allTags) != 2 {
		t.Error("Expected two tags")
	}
}

func TestMain(m *testing.M) {
	os.Setenv(util.BasePathENV, "/tmp/tim/")
	deleteTagFile()
	exitVal := m.Run()
	deleteTagFile()
	os.Exit(exitVal)
}

func deleteTagFile() {
	util.Delete(path())
}
