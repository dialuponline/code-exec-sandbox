package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", ".test.py")
	cmd.Env = []string{}
	reader, _ := cmd.StdoutPipe()

	stderr_reader, _ := cmd.StderrPipe()

	cmd.Start()

	go func() {
		for {
			buf := make([]byte, 1024)
	