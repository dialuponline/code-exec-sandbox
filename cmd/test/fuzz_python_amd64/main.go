package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/langgenius/dify-sandbox/internal/static/python_syscall"
)

const (
	SYSCALL_NUMS = 400
)

func run(allowed_syscalls []int) {
	nums := []string{}
	for _, syscall := range allowed_syscalls {
		nums = append(nums, strconv.Itoa(syscall))
	}
	os.Setenv("ALLOWED_SYSCALLS", strings.Join(nums, ","))
	_, err := exec.Command("