//go:build linux && amd64

package nodejs_syscall

import "syscall"

const (
	//334
	SYS_RSEQ = 334
	// 435
	SYS_CLONE3 = 435
)

var ALLOW_SYSCALLS = []int{
	syscall.SYS_OPEN, syscall.SYS_WRITE, syscall.SYS_CLOSE, syscall.SYS_READ,
	syscall.SYS_OPENAT, syscall.SYS_NEWFSTATAT, syscall.SYS_IOCTL, syscall.SYS_LSEEK,
	syscall.SYS_FSTAT,
	syscall.SYS_MPROTECT, syscall.SYS_MMAP, syscall.SYS_MUNMAP,
	sysca