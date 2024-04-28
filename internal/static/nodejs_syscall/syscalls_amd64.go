//go:build linux && amd64

package nodejs_syscall

import "syscall"

const (
	//334
	SYS_RSEQ = 334
	// 43