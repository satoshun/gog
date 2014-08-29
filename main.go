package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

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

func CloneCmd(s string, directory string) (cmd *exec.Cmd) {
	cmd = exec.Command("git", "clone", s, directory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}

func UpdateCmd(s string, directory string) (cmd *exec.Cmd) {
	cmd = exec.Command("git", "pull")
	cmd.Dir = directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}

func HookCmd(maps map[string]string) (cmd *exec.Cmd) {
	s := os.Getenv("GO_GIT_HOOK_CMD")
	if s == "" {
		return nil
	}
	var doc bytes.Buffer
	tmpl, _ := template.New("hook").Parse(s)
	tmpl.Execute(&doc, maps)

	cmd = exec.Command("/bin/sh", "-c", doc.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return
}

func CloneDirectory(d, s string) string {
	if s == "" {
		return "."
	}
	host, p, _ := SplitRepo(d)
	return path.Join(s, "src", host, p)
}

func main() {
	app := cli.NewApp()
	app.Name = "go-git"
	app.Version = "0.0.3"
	app.Usage = "use directory like Go"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "repository, r",
			Usage: "repository url",
		},
		cli.StringFlag{
			Name:  "base, b",
			Usage: "define git path",
		},
		cli.BoolFlag{
			Name:  "update, u",
			Usage: "update repository",
		},
	}

	app.Action = func(c *cli.Context) {
		repository := c.String("repository")
		if repository == "" {
			fmt.Println("please set repository option: -r or -repository")
			return
		}
		var s string
		for _, ca := range [...]string{c.String("base"), os.Getenv("GO_GIT_PATH"), os.Getenv("GOPATH")} {
			if ca != "" {
				s = ca
				break
			}
		}
		directory := CloneDirectory(repository, s)

		var cmd *exec.Cmd
		if c.Bool("update") {
			cmd = UpdateCmd(repository, directory)
		} else {
			cmd = CloneCmd(repository, directory)
		}
		err := cmd.Run()
		if err != nil {
			fmt.Println("fail command:", err)
			return
		}

		if !c.Bool("update") {
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
		}
	}

	app.Run(os.Args)
}
