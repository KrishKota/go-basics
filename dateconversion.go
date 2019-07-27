package main

import (
	"fmt"
	"time"
)

func main() {
	var val string
	fmt.Println("Enter time in string:")
	fmt.Scanln(&val)
	convertedDate := timeConversion(val)
	fmt.Println(convertedDate)
}

//MM-DD-YYYY
//2006-01-02 - yyyy-mm-dd
//03:04:05 - HH:MM:SS
func timeConversion(val string) string {
	layout1 := "03:04PM"
	layout2 := "15:04"
	//fmt.Printf("Value Passed :%v\n", val)
	t, err := time.Parse(layout1, val)
	if err != nil {
		fmt.Println("Parsing is not done correclty", err)
	}
	return t.Format(layout2)
}
