package runner

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/langgenius/dify-sandbox/internal/utils/log"
)

type OutputCaptureRunner struct {
	stdout chan []byte
	stderr chan []byte
	done   chan bool

	timeout time.Duration

	after_exit_hook func()
}

func NewOutputCaptureRunner() *OutputCaptureRunner {
	return &OutputCaptureRunner{
		stdout: make(chan []byte),
		stderr: make(chan []byte),
		done:   make(chan bool),
	}
}

func (s *OutputCaptureRunner) WriteError(data []byte) {
	if s.stderr != nil {
		s.stderr <- data
	}
}

func (s *OutputCaptureRunner) WriteOutput(data []byte) {
	if s.stdout != nil {
		s.stdout <- data
	}
}

func (s *OutputCaptureRunner) SetAfterExitHook(hook func()) {
	s.after_exit_hook = hook
}

func (s *OutputCaptureRunner) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *OutputCaptureRunner) CaptureOutput(cmd *exec.Cmd) error {
	// start a timer for the timeout
	timeout := s.timeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		if cmd != nil && cmd.Process != nil {
			// write the error
			s.WriteError([]byte("error: timeout\n"))
			// send a signal to the process
			cmd.Process.Kill()
		}
	}()

	// create a pipe for the stdout
	stdout_reader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// create a pipe for the stderr
	stderr_reader, err := cmd.StderrPipe