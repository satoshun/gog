package main

import (
	"os"
	"path"
	"path/filepath"
)

func Exists(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}
	return false

}

func GitDiretories(root string) []string {
	l := make([]string, 0)
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() || !Exists(path.Join(p, ".git")) {
			return nil
		}

		l = append(l, p)
		return filepath.SkipDir
	})

	return l
}
