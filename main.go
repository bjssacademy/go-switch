package main

import (
	"fmt"
	"time"
)

func main(){

		currentTime := time.Now()
		hour := currentTime.Hour()

		if hour < 12 {
			fmt.Println("Good morning")
		} else if hour < 18 {
			fmt.Println("Good afternoon")
		} else {
			fmt.Println("Good evening")
		}

		switch {
		case hour < 12:
			fmt.Println("Good morning")
		case hour < 18:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good evening")
		}

}