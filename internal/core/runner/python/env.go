package python

import (
	_ "embed"
	"os"
	"os/exec"
	"path"

	"github.com/langgenius/dify-sandbox/internal/core/runner"
	"github.com/langgenius/dify-sandbox/internal/static"
	"github.com/langgenius/dify-sandbox/internal/utils/log"
)

//go:embed env.sh
var env_script string

func PreparePythonDependenciesEnv() error {
	config := static.GetDifySandboxGlobalConfigurations()

	runner := runner.TempDirRunner{}
	err := runner.WithTempDir("/", []string{}, func(root_path string) error {
		err