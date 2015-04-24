package main

import (
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/codegangsta/cli"
)

// splitRepo split url to host, path, basename
func splitRepo(u string) (host, p, basename string) {
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

func projectDir(c *cli.Context, d string) string {
	host, p, _ := splitRepo(d)
	return path.Join(srcPath(c), host, p)
}

func basePath(c *cli.Context) string {
	for _, ca := range [...]string{c.String("base"), os.Getenv("GOG_PATH"), os.Getenv("GOPATH")} {
		if ca != "" {
			return ca
		}
	}

	return "./"
}

func srcPath(c *cli.Context) string {
	return basePath(c)
}

func main() {
	app := cli.NewApp()
	app.Name = "gog"
	app.Version = "0.3.0"
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
			Action:    actionGet,
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "update repository",
			Action:    actionUpdate,
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "list cloned repository",
			Action:    actionList,
		},
	}

	app.Run(os.Args)
}
