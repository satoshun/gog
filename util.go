package main

import (
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

func Exists(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}
	return false

}

func GitDiretories(root string) []string {
	l := make([]string, 0)
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() || !Exists(path.Join(p, ".git")) {
			return nil
		}

		l = append(l, p)
		return filepath.SkipDir
	})

	return l
}

func projectDir(c *cli.Context, d string) string {
	host, p, _ := splitRepo(d)
	return path.Join(basePath(c), host, p)
}

func basePath(c *cli.Context) string {
	for _, ca := range [...]string{c.String("base"), os.Getenv("GOG_PATH"), os.Getenv("GOPATH")} {
		if ca != "" {
			return ca
		}
	}

	return "./"
}

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
