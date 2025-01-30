package main

import (
	"fmt"
	"os"
	"UptimeKumaProbe/helpers"
)

func main() {
	args := os.Args

	//if (ArgsMatch(args, []string{"kprobe"})) {

	for i, arg := range args {
		fmt.Printf("[%d]: %s\n", i, arg)
	}
}