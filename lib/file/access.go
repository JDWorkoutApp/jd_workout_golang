package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func AccessFromCurrentDir(path string) string {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	absPath, err := filepath.Abs(filepath.Join(wd, path))

	if err != nil {
		panic(err)
	}

	return absPath
}

func AccessFromRootDir(path string) string {
	absPath := filepath.Join(GetRootDir(), path)

	return absPath
}

func GetRootDir() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get current directory")
		panic(err)
	}

	modPath := findFile(wd, ".env")
	if modPath != "" {
		return filepath.Dir(modPath)
	}

	panic("root not found")
}

func findFile(dir string, filename string) string {
	for {
		modPath := filepath.Join(dir, filename)
		_, err := os.Stat(modPath)
		if err == nil {
			return modPath
		}

		if dir == filepath.Dir(dir) {
			break
		}

		dir = filepath.Dir(dir)
	}

	panic("file not found")
}
