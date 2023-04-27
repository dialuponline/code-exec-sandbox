package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/langgenius/dify-sandbox/internal/static/nodejs_syscall"
)

const (
	SYSCALL_NUMS = 400
)

func run(allowed_syscalls []int) {
	os.Chdir("/tmp/sandbox-f5faacb9-f441-43ec-a077-a09a8b7cc7f0/var/sandbox/sandbox-nodejs/nodejs-project/node_temp/node_temp")

	nums := []string{}
	for _, syscall := range allowed_syscalls {
		nums = append(nums, strconv.Itoa(syscall))
	}
	os.Setenv("ALLOWED_SYSCALLS", strings.J