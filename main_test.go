package main

import (
	"testing"
)

func TestSplitRepo(t *testing.T) {
	// git protocol
	host, path, base := SplitRepo("git@github.com:satoshun/pythonjs.git")
	if host != "github.com" {
		t.Errorf("invalid host github.com", host)
	}
	if path != "/satoshun/pythonjs" {
		t.Errorf("invalid path /satoshun/pythonjs", path)
	}
	if base != "pythonjs" {
		t.Errorf("invalid base pythonjs", base)
	}

	// http, https protocol
	host, path, base = SplitRepo("https://github.com/satoshun/pythonjs.git")
	if host != "github.com" {
		t.Errorf("invalid host github.com", host)
	}
	if base != "pythonjs" {
		t.Errorf("invalid base pythonjs", base)
	}
}
