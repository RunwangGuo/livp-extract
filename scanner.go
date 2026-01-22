package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func ScanLivpFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if strings.ToLower(filepath.Ext(path)) == ".livp" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
