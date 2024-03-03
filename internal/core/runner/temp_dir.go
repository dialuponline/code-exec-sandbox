package runner

import (
	"os"
	"os/exec"
	"path"

	"github.com/google/uuid"
)

type TempDirRunner struct{}

func (s *TempDirRunner) WithTempDir(basedir string, paths []string, closures func(path string) error) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// create a tmp dir
	tmp_dir := path.Join(basedir, "tmp", "sandbo