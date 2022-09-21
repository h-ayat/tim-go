package config

import (
	"log"
	"os"
)

func BaseAddress() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return ""
	} else {
		return dirname + "/.tim/"
	}
}
