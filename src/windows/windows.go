package windows

import (
	"os"
	"os/exec"
)

func installChoco() error {
	// Check if Chocolatey is already installed
	_, err := exec.LookPath("choco")
	if err == nil {
		return nil // Chocolatey is already installed
	}

	// Download and run the Chocolatey installation script
	cmd := exec.Command("powershell.exe", "-Command", "(New-Object Net.WebClient).DownloadString('https://chocolatey.org/install.ps1') | iex")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err // Failed to install Chocolatey
	}
	return nil // Successfully installed Chocolatey
}
