package main

import (
	"fmt"

	"voz-archive-plus/utilities"
	"voz-archive-plus/vozscrape"
)

// this is the console application
func main() {
	fmt.Println(Box())
}

// Box will scape forum boxes
func Box() string {
	s := vozscrape.NewBox()
	res2B, _ := utilities.JSONMarshal(s, true)
	return string(res2B)
}
