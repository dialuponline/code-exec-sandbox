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
	syscall.SYS_MREMAP,
	syscall.SYS_BRK,
	syscall.SYS_RT_SIGACTION, syscall.SYS_RT_SIGPROCMASK,
	syscall.SYS_MADVISE, syscall.SYS_GETPID, syscall.SYS_GETUID,
	syscall.SYS_FCNTL, syscall.SYS_SIGALTSTACK, syscall.SYS_RT_SIGRETURN,
	syscall.SYS_FUTEX,
	syscall.SYS_EXIT_GROUP,
	syscall.SYS_EPOLL_CTL,
	syscall.SYS_EPOLL_PWAIT,
	syscall.SYS_SCHED_YIELD, syscall.SYS_EXIT,
	syscall.SYS_SCHED_GETAFFINITY, syscall.SYS_SET_ROBUST_LIST,
	SYS_RSEQ,

	syscall.SYS_SETUID, syscall.SYS_SETGID, syscall.SYS_GETTID,

	syscall.SYS_CLOCK_GETTIME, syscall.SYS_GETTIMEOFDAY, syscall.SYS_NANOSLEEP,
	syscall.SYS_TIME,

	syscall.SYS_TGKILL