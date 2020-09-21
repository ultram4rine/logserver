package cmd

import (
	"os"
	"os/exec"
)

var (
	// InstallWEBDependenciesCmd install dependencies of WEB app.
	InstallWEBDependenciesCmd = &exec.Cmd{
		Path:   "npm",
		Args:   append([]string{"npm"}, "install"),
		Dir:    "./ui",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	// BuildWEBAppCmd build WEB app.
	BuildWEBAppCmd = &exec.Cmd{
		Path:   "npm",
		Args:   append([]string{"npm"}, "run", "build"),
		Dir:    "./ui",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
)
