package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	if err := validateCommand("rm"); err != nil {
		fmt.Fprintf(os.Stderr, "Required command 'rm' is not available: %v\n", err)
		os.Exit(1)
	}

	trashPath, err := getTrashPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot determine trash path: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(trashPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create trash directory: %v\n", err)
		os.Exit(1)
	}

	processFiles(os.Args[1:], trashPath)
}

func getTrashPath() (string, error) {
	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome != "" {
		return filepath.Join(xdgDataHome, "Trash", "files"), nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".local", "share", "Trash", "files"), nil
}

func validateCommand(command string) error {
	_, err := exec.LookPath(command)
	return err
}

func processFiles(files []string, trashPath string) {
	for _, file := range files {
		if err := moveFileToTrash(file, trashPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error moving %s to trash: %v\n", file, err)
		}
	}
}

func moveFileToTrash(filename, trashPath string) error {
	source, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	dest := filepath.Join(trashPath, filepath.Base(source))
	return os.Rename(source, dest)
}

func printHelp() {
	fmt.Println("🗑️ Usage: gotrash [OPTION]... [FILE]...")
	fmt.Println("Remove (move to trash) the FILE(s).")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("📋 -h, --help         display this help and exit")
	fmt.Println("")
	fmt.Println("Note: Consider aliasing rm to gotrash:")
	fmt.Println("🔄 alias rm='gotrash'")
	fmt.Println("You can use gotrash as a proxy for the rm command with all its options 🚀")
	fmt.Println("")
}
