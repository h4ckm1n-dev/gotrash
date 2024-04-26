package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func getTrashPath() string {
	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome != "" {
		return filepath.Join(xdgDataHome, "Trash", "files")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".local", "share", "Trash", "files")
}

func main() {
	command := "rm"
	_, err := exec.LookPath(command)
	if err != nil {
		os.Args = append([]string{command}, os.Args...)
	}

	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		printHelp()
		os.Exit(0)
	}

	files := os.Args[1:]

	trashPath := getTrashPath()

	err = os.MkdirAll(trashPath, 0755)
	if err != nil {
		os.Exit(1)
	}

	for _, file := range files {
		err := moveFileToTrash(file, trashPath)
		if err != nil {
			fmt.Printf("Error moving %s to trash: %v\n", file, err)
		}
	}
}

func moveFileToTrash(filename, trashPath string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	dest := filepath.Join(trashPath, filepath.Base(filename))

	err := os.Rename(filename, dest)
	if err != nil {
		return err
	}

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		return fmt.Errorf("file %s not found in trash", filename)
	}

	return nil
}

func printHelp() {
	fmt.Println("🗑️ Usage: gotrash [OPTION]... [FILE]...")
	fmt.Println("Remove (move to trash) the FILE(s).")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("📋 -h, --help         display this help and exit")
	fmt.Println("")
	fmt.Println("Note: To use gotrash instead of rm, consider aliasing rm to gotrash by running:")
	fmt.Println("  🔄 alias rm=\"gotrash\"")
	fmt.Println("You can use gotrash as a proxy for the rm command with all options 🚀")
	fmt.Println("")
}
