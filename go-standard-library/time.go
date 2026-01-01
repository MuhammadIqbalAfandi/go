//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	local := time.Now().Local()
	fmt.Println(local)
	fmt.Println(local.Format("2006-01-02 15:04:05"))

	utc := time.Date(1998, time.April, 30, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2025-12-21 11:50:00"
	valueTime, err := time.Parse(formatter, value) 
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Weekday())

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowLoc := time.Now().In(loc)
	fmt.Println(nowLoc)
}