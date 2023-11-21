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
		"/etc/nsswitch.