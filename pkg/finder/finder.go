package finder

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func FindProject(name string) (string, error) {
	from, err := resolveFolders()
	if err != nil {
		return "", err
	}

	for _, root := range from {
		projectPath := ""
		rootDepth := depth(root)

		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if !d.IsDir() {
				return nil
			}

			if depth(path)-rootDepth > 2 {
				return filepath.SkipDir
			}

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

func depth(path string) int {
	return len(strings.Split(filepath.Clean(path), string(os.PathSeparator)))
}
