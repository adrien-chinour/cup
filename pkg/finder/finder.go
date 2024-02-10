package finder

import (
	"errors"
	"os"
	"path/filepath"
)

func FindProject(name string) (string, error) {
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
