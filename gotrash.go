package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
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

	processFiles(os.Args[1:])
}

func getTrashPath(uid string) (string, error) {
	u, err := user.LookupId(uid)
	if err != nil {
		return "", err
	}
	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome != "" {
		return filepath.Join(xdgDataHome, "Trash", "files"), nil
	}
	return filepath.Join(u.HomeDir, ".local", "share", "Trash", "files"), nil
}

func validateCommand(command string) error {
	_, err := exec.LookPath(command)
	return err
}

func processFiles(files []string) {
	for _, file := range files {
		if err := moveFileToTrash(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error moving %s to trash: %v\n", file, err)
		}
	}
}

func moveFileToTrash(filename string) error {
	source, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	fileInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("file %s does not exist", filename)
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to get file owner info")
	}

	uid := strconv.Itoa(int(stat.Uid))
	trashPath, err := getTrashPath(uid)
	if err != nil {
		return fmt.Errorf("cannot determine trash path: %v", err)
	}

	if err := os.MkdirAll(trashPath, 0755); err != nil {
		return fmt.Errorf("failed to create trash directory: %v", err)
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
