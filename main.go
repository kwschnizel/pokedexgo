package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	res := strings.Fields(text)
	return res
}
