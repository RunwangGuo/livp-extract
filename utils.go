package main

import "os"

func IsImageExt(ext string) bool {
	switch ext {
	case ".heic", ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

func IsTargetExt(ext string) bool {
	return ext == ".mov" || IsImageExt(ext)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}
