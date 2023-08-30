package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// Create date
	utc := time.Date(2023, 10, 06, 07, 46, 12, 30, time.Local)
	fmt.Println(utc)

	// Passing string
	parse, _ := time.Parse(time.RFC1123, "Mon, 02 Jan 2006 15:04:05 MST")
	fmt.Println(parse)

	// Full Docs: https://pkg.go.dev/time

}
