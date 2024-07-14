package integrationtests_test

import (
	"strings"
	"testing"

	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/service"
)

func TestPythonBase64(t *testing.T) {
	// Test case for base64
	runMultipleTestings(t, 50, func(t *testing.T) {
		resp := service.RunPython3Code(`
import base64
print(base6