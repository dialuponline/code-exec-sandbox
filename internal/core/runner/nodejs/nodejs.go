package nodejs

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/langgenius/dify-sandbox/internal/core/runner"
	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/static"
)

type NodeJsRunner struct {
	runner.TempDirRunner
}

//go:embed prescript.js
var nodejs_sandbox_fs []byte

var (
	REQUIRED_FS = []string{
		path.Join(LIB_PATH, PROJECT_NAME, "node_temp"),
		path.Join(LIB_PATH, LIB_NAME),
		"/etc/ssl/certs/ca-certificates.crt",
		"/etc/nsswitch.conf",
		"/etc/resolv.conf",
		"/run/systemd/resolve/stub-resolv.conf",
		"/etc/hosts",
	}
)

func (p *NodeJsRunner) Run(
	code string,
	timeout time.Duration,
	stdin []byte,
	preload string,
	options *types.RunnerOptions,
) (chan []byte, chan []byte, chan bool, error) {
	configuration := static.GetDifySandboxGlobalConfigurations()

	// capture the output
	output_handler := runner.NewOutputCaptureRunner()
	output_handler.SetTimeout(timeout)

	err := p.WithTempDir("/", REQUIRED_FS, func(root_path string) error {
		output_handler.SetAfterExitHook(func() {
			os.RemoveAll(root_path)
			os.Remove(root_path)
		})

		// initialize the environment
		script_path, err := p.InitializeEnvironment(code, preload, root_path)
		if err != nil {
			return err
		}

		// create a new process
		cmd := exec.Command(
			static.GetDifySandboxGlobalConfigurations().NodejsPath,
		