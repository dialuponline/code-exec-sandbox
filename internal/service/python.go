package service

import (
	"time"

	"github.com/langgenius/dify-sandbox/internal/core/runner/python"
	runner_types "github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/static"
	"github.com/langgenius/dify-sandbox/internal/types"
)

type RunCodeResponse struct {
	Stderr string `json:"error"`
	Stdout string `json:"stdout"`
}

func RunPython3Code(code string, preload string, options *runner_types.RunnerOptions) *types.DifySandboxResponse {
	if err := checkOptions(options); err != nil {
		return types.ErrorResponse(-400, err.Error())
	}

	timeout := time.Duration(
		static.GetDifySandboxGlobalConfigu