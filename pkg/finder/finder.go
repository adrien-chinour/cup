package finder

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindProject(name string) (string, error) {
	from, err := resolveFolders()
	if err != nil {
		return "", err
	}

	fmt.Println(from)

	for _, root := range from {
		projectPath := ""
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if d.Name() == name && d.IsDir() {
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

func resolveFolders() ([]string, error) {
	env := os.Getenv("CUPDIR")
	if env == "" {
		return nil, errors.New("CUPDIR is not defined")
	}

	return strings.Split(env, ":"), nil
}
