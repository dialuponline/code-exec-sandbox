package integrationtests_test

import (
	"strings"
	"testing"

	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/service"
)

func TestNodejsBase64(t *testing.T) {
	// Test case for base64
	runMultipleTestings(t, 30, func(t *testing.T) {
		resp := service.RunNodeJsCode(`
const base64 = Buffer.from("hello world").toString("base64");
console.log(Buffer.from(base64, "base64").toString());
		`, "", &types.RunnerOptions{
			EnableNetwork: true,
		})
		if resp.Code != 0 {
			t.Fatal(resp)
		}

		if resp.Data.