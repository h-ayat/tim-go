package tags

import (
	"bufio"
	"errors"
	"log"
	"os"

	"h-ayat.github.io/tim/config"
)

var path = config.BaseAddress() + "tags.txt"

func LoadAllTags() []string {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return []string{}
	} else {
		f, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		data := []string{}

		for scanner.Scan() {

			data = append(data, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return data
	}
}

func AddTag(tag string) {
	os.MkdirAll(config.BaseAddress(), os.ModePerm)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(tag + "\n"); err != nil {
		panic(err)
	}
}

func DeleteTagFile() {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return
	} else {
		os.Remove(path)
	}
}
