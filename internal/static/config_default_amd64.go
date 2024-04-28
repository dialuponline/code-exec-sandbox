//go:build linux && amd64

package static

var DEFAULT_PYTHON_LIB_REQUIREMENTS = []string{
	"/usr/local/lib/python3.10",
	"/usr/lib/python3.10",
	"/usr/lib/python3",
	"/usr/lib/x86_64-linux-gnu/libssl.so.3",
	"/usr/lib/x86_64-linux-gnu/libcrypto.so.3",