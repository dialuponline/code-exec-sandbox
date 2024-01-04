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
	