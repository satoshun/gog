package main

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"
)

func hookCmd(maps map[string]string) (cmd *exec.Cmd) {
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
