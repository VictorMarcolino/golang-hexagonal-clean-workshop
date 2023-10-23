package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func GetEnvWithDefault(variableName, defaultValue string) string {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetPathRelativeToFolder(subfolder string) string {
	projectFolderName := GetEnvWithDefault("PROJECT_FOLDER", "golang-hexagonal-clean-workshop")
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(fmt.Sprintf(`(.*%v)`, projectFolderName))
	matches := r.FindStringSubmatch(cwd)
	if len(matches) != 2 {
		panic("directory not found in path")
	}
	basePath := matches[1]

	return filepath.Join(basePath, subfolder)
}
