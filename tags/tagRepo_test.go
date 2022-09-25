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
	DeleteTagFile()
	AddTag("testA")
	AddTag("testB")
	allTags = LoadAllTags()
	DeleteTagFile()

	if len(allTags) != 2 {
		t.Error("Expected two tags")
	}
}

func TestMain(m *testing.M) {
	os.Setenv(util.BasePathENV, "/tmp/tim/")
	fmt.Println("Set up stuff for tests here")
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}
