package main

import (
	"fmt"

	"h-ayat.github.io/tim/cal"
	"h-ayat.github.io/tim/entries"
)

func main() {

	entries.AddNewEntry("khar")
	entries.AddNewEntry("boz")

	fmt.Println(entries.Read(cal.CurrentDate().Ago(32)))
}
