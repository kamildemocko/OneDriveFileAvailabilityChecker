package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("error: usage: <app> PATH")
		return
	}

	joinedPath := strings.Join(args, " ")
	cleanPath := filepath.Clean(joinedPath)
	path, err := filepath.Abs(cleanPath)
	if err != nil {
		fmt.Printf("error: the file path '%s' is not correct\n", cleanPath)
		return
	}

	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("error: the file '%s' does not exist\n", path)
		return
	}

	psScript := fmt.Sprintf(`
		$folder = [System.IO.Path]::GetDirectoryName("%s")
		$file = [System.IO.Path]::GetFileName("%s")
		$shell = New-Object -ComObject Shell.Application
		$folder = $shell.Namespace($folder)
		$item = $folder.ParseName($file)
		$syncStatus = $folder.GetDetailsOf($item, 303)
		Write-Output $syncStatus
	`, path, path)

	cmd := exec.Command("powershell", "-Command", psScript)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		fmt.Printf("error: error running command: %s", err)
	}

	status := strings.TrimSpace(out.String())

	switch status {
	case "Available offline", "Always available on this device", "Available on this device", "Available when online":
		fmt.Printf("available: %s\n", status)
	case "Sync pending":
		fmt.Printf("pending: %s\n", status)
	case "Error":
		fmt.Printf("not_available: %s\n", status)
	default:
		fmt.Printf("unknown_status: %s\n", status)
	}
}