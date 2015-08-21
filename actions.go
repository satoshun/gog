package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

func actionGet(c *cli.Context) {
	rURL := c.Args().First()
	if rURL == "" {
		log.Fatal("please set repository url")
	}

	wd := projectDir(c, rURL)
	cmd := git.NewGit(wd).Clone(rURL)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	host, p, base := splitRepo(rURL)
	cmd = hookCmd(map[string]string{
		"Directory":   wd,
		"Repository":  rURL,
		"Host":        host,
		"Path":        p,
		"ProjectName": base,
	})

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	link := c.String("link")
	if link != "" {
		oldname := wd
		newPath := linkPath(c, link)
		if err := os.MkdirAll(newPath, 0755); err != nil {
			log.Fatal(err)
		}
		if err := os.Symlink(oldname, path.Join(newPath, base)); err != nil {
			log.Fatal(err)
		}
	}
}

func actionUpdate(c *cli.Context) {
	rURL := c.Args().First()
	// no specified repository url then all update
	if rURL == "" {
		var wg sync.WaitGroup
		makeWorkerPool(func(d string) {
			git := git.NewGit(d)
			if git.HasRemote() {
				git.Update().Run()
			}

			wg.Done()
		})
		for _, d := range retriveGitDirs(basePath(c)) {
			log.Println("update", d)
			wg.Add(1)
			go func(url string) {
				channel := <-workerPool
				channel <- url
			}(d)
		}

		wg.Wait()
		return
	}

	directory := projectDir(c, rURL)
	git := git.NewGit(directory)
	if err := git.Update().Run(); err != nil {
		log.Fatal(err)
	}
}

func actionList(c *cli.Context) {
	var paths []map[string]string
	cwd := basePath(c) + "/"
	maxLen := 0

	for _, d := range retriveGitDirs(cwd) {
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
