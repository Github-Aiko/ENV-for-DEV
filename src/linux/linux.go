package linux

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func IsLinux() bool {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			id := strings.TrimPrefix(line, "ID=")
			if id == "ubuntu" || id == "centos" || id == "debian" {
				return true
			}
		}
	}

	return false
}

func UpdatePackage() {
	if IsLinux() { // Check if Linux
		if _, err := os.Stat("/usr/bin/apt-get"); err == nil { // Check if apt-get exists
			// Update package
			cmd := exec.Command("apt-get", "update")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else if _, err := os.Stat("/usr/bin/yum"); err == nil { // Check if yum exists
			// Update package
			cmd := exec.Command("yum", "update")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}
