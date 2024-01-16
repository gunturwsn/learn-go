package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2001, time.March, 31, 10, 11, 43, 0, time.UTC)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2010-10-10 10:10:10"
	parse, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(parse)
	}

	fmt.Println(parse.Year())
	fmt.Println(parse.Month())
	fmt.Println(parse.Day())
	fmt.Println(parse.Hour())
	fmt.Println(parse.Minute())
	fmt.Println(parse.Second())
}
