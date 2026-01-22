package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ExtractLivp(livpPath, dstDir string, onlyMov, onlyImage bool) error {
	r, err := zip.OpenReader(livpPath)
	if err != nil {
		return err
	}
	defer r.Close()

	baseName := strings.TrimSuffix(
		filepath.Base(livpPath),
		filepath.Ext(livpPath),
	)

	for _, f := range r.File {
		ext := strings.ToLower(filepath.Ext(f.Name))

		if onlyMov {
			if ext != ".mov" {
				continue
			}
		} else if onlyImage {
			if !IsImageExt(ext) {
				continue
			}
		} else {
			if !IsTargetExt(ext) {
				continue
			}
		}

		dstPath := filepath.Join(dstDir, baseName+ext)
		if err := extractZipFile(f, dstPath); err != nil {
			return err
		}
	}

	return nil
}

func extractZipFile(f *zip.File, dstPath string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return err
	}

	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, rc)
	return err
}
