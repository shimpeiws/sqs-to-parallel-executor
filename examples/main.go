package main

import (
	"fmt"
	"os"
)

func main() {
	body := os.Args[1]
	fmt.Printf("input body = %s\n", body)
}
