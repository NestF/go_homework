package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify the temperature in °C")
		return
	}

	celsius, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("error: input is not a number")
		return
	}

	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.1f °F\n", fahrenheit)
}
