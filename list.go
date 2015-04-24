package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

func actionList(c *cli.Context) {
	var paths []map[string]string
	cwd := basePath(c) + "/"
	maxLen := 0

	for _, d := range GitDiretories(cwd) {
		path := strings.TrimPrefix(d, cwd)
		if len(path) > maxLen {
			maxLen = len(path)
		}
		paths = append(paths, map[string]string{
			"path": path,
			"full": d,
		})
	}

	f := "%-" + strconv.Itoa(maxLen+2) + "s"
	for _, d := range paths {
		fmt.Printf(f, d["path"])
		git := git.NewGit(d["full"])
		git.LogOneline().Run()
	}
}
