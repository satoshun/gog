package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

func actionGet(c *cli.Context) {
	repository := c.Args().First()
	if repository == "" {
		fmt.Println("please set repository url")
		return
	}

	directory := projectDir(c, repository)
	cmd := git.NewGit(directory).Clone(repository)
	err := cmd.Run()
	if err != nil {
		fmt.Println("fail command:", err)
		return
	}

	host, p, base := splitRepo(repository)
	cmd = hookCmd(map[string]string{
		"Directory":   directory,
		"Repository":  repository,
		"Host":        host,
		"Path":        p,
		"ProjectName": base,
	})

	if cmd != nil {
		err = cmd.Run()
		if err != nil {
			fmt.Println("fail hook command:", cmd, err)
			return
		}
	}
}
