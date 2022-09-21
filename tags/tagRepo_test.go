package tags

import (
	"fmt"
	"testing"
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
