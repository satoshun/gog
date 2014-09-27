package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func CloneCmd(s string, directory string) (cmd *exec.Cmd) {
	args := []string{"clone", s, directory}
	cmd = gitCmd(args)
	return
}

func UpdateCmd(directory string) (cmd *exec.Cmd) {
	args := []string{"pull"}
	cmd = gitCmd(args)
	cmd.Dir = directory
	return
}

func OriginUrl(path string) string {
	args := []string{"config", "--get", "remote.origin.url"}
	cmd := exec.Command("git", args...)
	output := new(bytes.Buffer)
	cmd.Stdout = output
	cmd.Stderr = output
	cmd.Dir = path

	err := cmd.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(output.String())
}

func LogCmd(directory string) (cmd *exec.Cmd) {
	args := []string{"--no-pager", "log", "-1", "--oneline"}
	cmd = gitCmd(args)
	cmd.Dir = directory
	return
}

func gitCmd(args []string) (cmd *exec.Cmd) {
	cmd = exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}

func HookCmd(maps map[string]string) (cmd *exec.Cmd) {
	s := os.Getenv("GOG_HOOK_CMD")
	if s == "" {
		return nil
	}
	var doc bytes.Buffer
	tmpl, _ := template.New("hook").Parse(s)
	tmpl.Execute(&doc, maps)

	cmd = shellCmd(doc.String())
	return
}

func shellCmd(script string) (cmd *exec.Cmd) {
	cmd = exec.Command("/bin/sh", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return
}
