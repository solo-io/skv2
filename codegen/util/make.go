package util

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunMake(target string, opts ...func(*exec.Cmd)) error {
	cmd := exec.Command("make", "-B", "-C", filepath.Dir(GoModPath()), target)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd.Run()
}
