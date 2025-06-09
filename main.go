package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("More commands needed")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func run() {
	fmt.Printf("Running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // Hostname
			syscall.CLONE_NEWPID | // Process IDs
			syscall.CLONE_NEWNS | // Mount points
			syscall.CLONE_NEWNET | // Network
			syscall.CLONE_NEWIPC, // IPC
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}
