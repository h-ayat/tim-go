package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
)

func ReadAll(path string) []string {
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

func Append(path string, content string) {

	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content + "\n"); err != nil {
		panic(err)
	}
}

func Overwrite(path string, content string) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func Delete(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return
	} else {
		os.Remove(path)
	}

}
