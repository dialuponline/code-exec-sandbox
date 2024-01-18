package dependencies

import (
	"sync"

	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
)

var preload_script_map = map[string]string{}
var preload_script_map_lock = &sync.RWMutex{}

func SetupDependency(package_name string, version string) {
	preload_script_map_lock.Lock()
	d