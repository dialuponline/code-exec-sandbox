package nodejs

import (
	"embed"
	"fmt"
	"os"
	"path"

	"github.com/langgenius/dify-sandbox/internal/utils/log"
)

const (
	LIB_PATH     = "/var/sandbox/sandbox-nodejs"
	LIB_NAME     = "nodejs.so"
	PROJECT_NAME = "nodejs-project"
)

//go:embed nodejs.so
var nodejs_lib []byte

//go:embed dependens
var nodejs_dependens embed.FS // it's a directory

func init() {
	releaseLibBinary()
}

func releaseLibBinary() {
	log.Info("initializing nodejs runner environment...")
	os.RemoveAll(LIB_PATH)
	os.Remove(LIB_PATH)

	err := os.MkdirAll(LIB_PATH, 0755)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to create %s", LIB_PATH))
	}
	err = os.WriteFile(path.Join(LIB_PATH, LIB_NAME), nodejs_lib, 0755)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to write %s", path.Join(LIB_PATH, 