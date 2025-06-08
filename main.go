package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("More commands needed")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		run()
	}
}

func run() {
	fmt.Printf("Funning %v as PID %d\n", os.Args[2:], os.Getpid())
}
