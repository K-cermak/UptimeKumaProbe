package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for i, arg := range args {
		fmt.Printf("[%d]: %s\n", i, arg)
	}
}

func getHelp() {
	fmt.Println("Usage: kprobe")
}