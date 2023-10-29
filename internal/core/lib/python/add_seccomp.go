package python

import (
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/langgenius/dify-sandbox/internal/core/lib"
	"github.com/langgenius/dify-sandbox/internal/static/python_syscall"
)

//var allow_syscalls = []int{}

func InitSeccomp(uid int, gid int, enable_network bool) error {
	err := syscall.Chroot(".")
	if err != nil {
		return err
	