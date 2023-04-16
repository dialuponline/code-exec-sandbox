rm -f internal/core/runner/python/python.so
rm -f internal/core/runner/nodejs/nodejs.so
rm -f /tmp/sandbox-python/python.so
rm -f /tmp/sandbox-nodejs/nodejs.so
echo "Building Python lib"
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o internal/core/runner/python/python.so -buildmode=c-shared -ldflags="-s -w" cmd/lib/python/main.go &&
