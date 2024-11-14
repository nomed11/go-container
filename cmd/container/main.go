package main

import (
	"fmt"
	"os"
	"path/filepath"

	"go-container/internal/container"
	"go-container/internal/fs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: container run <command>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func run() {
	// Create new container with command
	cont, err := container.NewContainer(os.Args[2:])
	if err != nil {
		fmt.Printf("Error creating container: %v\n", err)
		os.Exit(1)
	}

	// Setup root filesystem
	rootfs := fs.NewRootFS(filepath.Join(".", "rootfs"))
	if !rootfs.Exists() {
		fmt.Println("Root filesystem not found. Please run 'make setup-rootfs' first")
		os.Exit(1)
	}

	// Run the container
	if err := cont.Run(); err != nil {
		fmt.Printf("Error running container: %v\n", err)
		os.Exit(1)
	}
}

func child() {
	cont, err := container.NewContainer(os.Args[2:])
	if err != nil {
		fmt.Printf("Error creating container: %v\n", err)
		os.Exit(1)
	}

	if err := cont.Child(); err != nil {
		fmt.Printf("Error executing child process: %v\n", err)
		os.Exit(1)
	}
}