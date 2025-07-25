package runner

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"achinour/cup/pkg/finder"
)

func Run(action string, project string) {
	path, err := finder.FindProject(project)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = runCommand(path, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func runCommand(projectPath string, action string) error {
	args := []string{"compose"}

	switch action {
	case "up":
		args = append(args, "up", "-d")
	case "stop":
		args = append(args, "stop")
	case "down":
		args = append(args, "down")
	default:
		return errors.New("not implemented action")
	}

	cmd := exec.Command("docker", args...)
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
