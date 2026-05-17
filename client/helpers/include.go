package helpers

import (
	"fmt"
	"path/filepath"
)

func Include(path string) []string {

	FileList, err := filepath.Glob("client/views/templates/*.html")
	if err != nil {
		fmt.Println(err)
	}

	PathFiles, err := filepath.Glob("client/views/" + path + "/*.html")
	if err != nil {
		fmt.Println(err)
	}

	FileList = append(FileList, PathFiles...)

	return FileList
}
