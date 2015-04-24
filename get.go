package main

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

func actionGet(c *cli.Context) {
	rURL := c.Args().First()
	if rURL == "" {
		log.Fatal("please set repository url")
	}

	cwd := projectDir(c, rURL)
	cmd := git.NewGit(cwd).Clone(rURL)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	host, p, base := splitRepo(rURL)
	cmd = hookCmd(map[string]string{
		"Directory":   cwd,
		"Repository":  rURL,
		"Host":        host,
		"Path":        p,
		"ProjectName": base,
	})

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
