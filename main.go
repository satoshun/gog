package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/satoshun/go-git"
)

// SplitRepo split url to host, path, basename
func SplitRepo(u string) (host, p, basename string) {
	if i := strings.Index(u, "@"); i >= 0 {
		u = u[i+1:]
		tmp := strings.SplitN(u, ":", 2)
		host, p = tmp[0], tmp[1]
	} else {
		u, _ := url.Parse(u)
		host, p = u.Host, u.Path
	}

	// start
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	// remove .git
	if strings.HasSuffix(p, ".git") {
		p = p[:len(p)-4]
	}
	basename = path.Base(p)

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
	return BasePath(c)
}

func main() {
	app := cli.NewApp()
	app.Name = "gog"
	app.Version = "0.2.2"
	app.Usage = "structure directory like Go"
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
				cmd := git.NewGit(directory).Clone(repository)
				err := cmd.Run()
				if err != nil {
					fmt.Println("fail command:", err)
					return
				}

				host, p, base := SplitRepo(repository)
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
							git := git.NewGit(d)
							git.Update().Run()
							wg.Done()
						}(d)
					}

					wg.Wait()
					return
				}

				directory := ProjectDir(c, repository)
				git := git.NewGit(directory)
				err := git.Update().Run()
				if err != nil {
					fmt.Println("fail command:", err)
					return
				}
			},
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "list cloned repository",
			Action: func(c *cli.Context) {
				var paths []map[string]string
				srcPath := SrcPath(c) + "/"
				maxLen := 0

				for _, d := range GitDiretories(srcPath) {
					path := strings.TrimPrefix(d, srcPath)
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
			},
		},
	}

	app.Run(os.Args)
}
