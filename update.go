package main

import (
	"log"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

func actionUpdate(c *cli.Context) {
	rURL := c.Args().First()
	// if no specified repository url then all update
	if rURL == "" {
		var wg sync.WaitGroup
		for _, d := range GitDiretories(basePath(c)) {
			log.Println("update", d)
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

	directory := projectDir(c, rURL)
	git := git.NewGit(directory)
	if err := git.Update().Run(); err != nil {
		log.Fatal(err)
	}
}
