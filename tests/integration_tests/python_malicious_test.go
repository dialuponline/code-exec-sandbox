package integrationtests_test

import (
	"strings"
	"testing"

	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/service"
)

func TestSysFork(t *testing.T) {
	// Test case for sys_fork
	resp := service.RunPython3Code(`
import os
print(os.fork())
print(123)
	`, "", &ty