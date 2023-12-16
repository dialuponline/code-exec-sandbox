package nodejs

import (
	"embed"
	"fmt"
	"os"
	"path"

	"github.com/langgenius/dify-sandbox/internal/utils/log"
)

const (
	LIB_PATH     = "/var/sandbox/s