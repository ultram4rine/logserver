package cmd

import (
	"os"
	"os/exec"
)

var (
	// InstallWEBDependenciesCmd install dependencies of WEB app.
	InstallWEBDependenciesCmd = exec.Command("npm", "install")

	// BuildWEBAppCmd build WEB app.
	BuildWEBAppCmd = exec.Command("npm", "run", "build")
)

func init() {
	InstallWEBDependenciesCmd.Dir = "./ui"
	InstallWEBDependenciesCmd.Stdout = os.Stdout
	InstallWEBDependenciesCmd.Stderr = os.Stderr

	BuildWEBAppCmd.Dir = "./ui"
	BuildWEBAppCmd.Stdout = os.Stdout
	BuildWEBAppCmd.Stderr = os.Stderr
}
