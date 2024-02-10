package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

func main() {
	action, project, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path, err := findProject(project)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = runCommand(path, action)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Finish")
}

// Parse user args to resolve action and project name
func readArgs(args []string) (string, string, error) {
	action := "start"
	project := ""

	// FIXME find a way to not write "else" block
	if len(args) == 3 {
		action = args[1]
		project = args[2]
	} else if len(args) == 2 {
		project = args[1]
	} else {
		return "", "", errors.New("invalid number of args")
	}

	if !slices.Contains([]string{"start", "stop", "kill"}, action) {
		return "", "", errors.New("invalid command")
	}

	return action, project, nil
}

// Find project path by name
// If 2 projects has same name, the first one will be chosen
func findProject(project string) (string, error) {
	// TODO use env variable or custom dotfile for this
	// TODO maybe a recursive version but seems more difficult (don't want to stuck in a node_module folder..)
	from := []string{
		"/home/adrien/code/adrien-chinour",
		"/home/adrien/code/phpimages",
		"/home/adrien/code/boxop",
	}

	for _, root := range from {
		projectPath := ""
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if d.Name() == project && d.IsDir() {
				projectPath = path
			}

			return nil
		})

		if projectPath != "" {
			return projectPath, nil
		}

		if err != nil {
			return "", err
		}
	}

	return "", errors.New("project not found")
}

// Run docker compose command based on action and projectPath
func runCommand(projectPath string, action string) error {
	args := []string{"compose"}

	switch action {
	case "start":
		args = append(args, "up", "-d")
	case "stop":
		args = append(args, "stop")
	case "kill":
		args = append(args, "down")
	default:
		return errors.New("not implemented action")
	}

	cmd := exec.Command("docker", args...)
	cmd.Dir = projectPath

	fmt.Println(fmt.Sprintf("Run %v action for project %v", action, projectPath))

	return cmd.Run()
}
