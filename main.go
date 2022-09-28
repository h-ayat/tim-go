package main

import (
	"fmt"
	"os"
	"strings"

	"h-ayat.github.io/tim/entries"
)

const version = "0.0.1"

func addEntry(args []string) {

	entries.AddNewEntry(strings.Join(args, " "))
}

func printHelp() {
	message := fmt.Sprintf(`
	Tim %s

	Usage:
	(No switches)            Add entry (tim do something)
	`, version)

	fmt.Println(message)
}

func handleSwitches(args []string) {

}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
	} else {
		if strings.HasPrefix(args[0], "-") {
			handleSwitches(args)
		} else {
			addEntry(args)
		}
	}

}
