package main

import (
	"github.com/langgenius/dify-sandbox/internal/core/lib/python"
)
import "C"

//export DifySeccomp
func DifySeccomp(uid int, g