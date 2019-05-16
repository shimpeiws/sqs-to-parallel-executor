package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	body := os.Args[1]
	fmt.Print(body)

	rand.Seed(time.Now().UnixNano())
	random := rand.Int()
	if random%5 == 0 {
		fmt.Print("Error From examples")
		os.Exit(1)
	}
}
