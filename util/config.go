package util

import (
	"log"
	"os"
)

const BasePathENV = "TIM_BASE_PATH"

func BaseAddress() string {
	var b = os.Getenv(BasePathENV)
	if b == "" {
		b = defaultBasePath()
	}
	return b
}

func defaultBasePath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return ""
	} else {
		return dirname + "/.tim/"
	}
}
