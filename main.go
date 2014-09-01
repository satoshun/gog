package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/codegangsta/cli"
)

func SplitRepo(s string) (host, p, base string) {
	if strings.HasPrefix(s, "git@") {
		s = s[4:]
		tmp := strings.SplitN(s, ":", 2)
		host, p = tmp[0], tmp[1]
	} else {
		u, _ := url.Parse(s)
		host, p = u.Host, u.Path
	}

	// start /
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	// remove .git
	if strings.HasSuffix(p, ".git") {
		p = p[:len(p)-4]
	}
	base = path.Base(p)

	return
}

func ProjectDir(c *cli.Context, d string) string {
	host, p, _ := SplitRepo(d)
	return path.Join(SrcPath(c), host, p)
}

func BasePath(c *cli.Context) string {
	for _, ca := range [...]string{c.String("base"), os.Getenv("GOG_PATH"), os.Getenv("GOPATH")} {
		if ca != "" {
			return ca
		}
	}

	return "./"
}

func SrcPath(c *cli.Context) string {
	return path.Join(BasePath(c), "src")
}

func main() {
	app := cli.NewApp()
	app.Name = "gog"
	app.Version = "0.2.1"
	app.Usage = "use directory like Go"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "base, b",
			Usage: "define git path",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "get",
			ShortName: "g",
			Usage:     "clone repository",
			Action: func(c *cli.Context) {
				repository := c.Args().First()
				if repository == "" {
					fmt.Println("please set repository url")
					return
				}

				directory := ProjectDir(c, repository)
				cmd := CloneCmd(repository, directory)
				err := cmd.Run()
				if err != nil {
					fmt.Println("fail command:", err)
					return
				}

				host, p, base := SplitRepo(repository)
				cmd = HookCmd(map[string]string{
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
			},
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "update repository",
			Action: func(c *cli.Context) {
				repository := c.Args().First()
				if repository == "" {
					// all update
					var wg sync.WaitGroup
					for _, d := range GitDiretories(SrcPath(c)) {
						fmt.Println("update", d)
						wg.Add(1)
						go func(d string) {
							cmd := UpdateCmd(d)
							cmd.Run()
							wg.Done()
						}(d)
					}

					wg.Wait()
					return
				}

				directory := ProjectDir(c, repository)
				cmd := UpdateCmd(directory)
				err := cmd.Run()
				if err != nil {
					fmt.Println("fail command:", err)
					return
				}
			},
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "list clone repository",
			Action: func(c *cli.Context) {
				srcPath := SrcPath(c) + "/"
				for _, d := range GitDiretories(SrcPath(c)) {
					fmt.Printf("%s\t", strings.TrimPrefix(d, srcPath))
					LogCmd(d).Run()
				}
			},
		},
	}

	app.Run(os.Args)
}
