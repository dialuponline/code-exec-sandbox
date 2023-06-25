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
	_, err := exec.Command("python3", "cmd/test/fuzz_python_amd64/test.py").Output()
	if err == nil {
	} else {
		fmt.Println("failed")
	}
}

func find_syscall(syscall int, syscalls []int) int {
	for i, s := range syscalls {
		if s == syscall {
			return i
		}
	}
	return -1
}

func main() {
	or