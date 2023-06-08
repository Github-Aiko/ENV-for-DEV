package macos

import (
	"fmt"
	"os/exec"
	"runtime"
)

func InstallHomebrew() error {
	var cmd *exec.Cmd
	if runtime.GOARCH == "arm64" {
		cmd = exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	} else {
		cmd = exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install Homebrew: %w", err)
	}
	return nil
}
