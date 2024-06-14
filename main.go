package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("error: usage: <app> PATH")
		return
	}

	path := strings.Join(args, " ")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf("error: the file '%s' does not exist\n", path))
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

	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("error: error running command: %s", err))
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