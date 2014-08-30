package main

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"
)

func CloneCmd(s string, directory string) (cmd *exec.Cmd) {
	cmd = exec.Command("git", "clone", s, directory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}

func UpdateCmd(directory string) (cmd *exec.Cmd) {
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
