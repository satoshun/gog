package main

import (
	"log"
	"os"
	"path"

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
