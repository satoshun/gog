package main

import (
	"fmt"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

func actionUpdate(c *cli.Context) {
	repository := c.Args().First()
	if repository == "" {
		// all update
		var wg sync.WaitGroup
		for _, d := range GitDiretories(basePath(c)) {
			fmt.Println("update", d)
			wg.Add(1)
			go func(d string) {
				git := git.NewGit(d)
				if git.HasRemote() {
					git.Update().Run()
				}

				wg.Done()
			}(d)
		}

		wg.Wait()
		return
	}

	directory := projectDir(c, repository)
	git := git.NewGit(directory)
	err := git.Update().Run()
	if err != nil {
		fmt.Println("fail command:", err)
		return
	}
}
