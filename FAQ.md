## FAQ

### 1. Why does my Python code throw an exception like "xxx.so: cannot open shared object file: No such file or directory"?

This occurs because the `dify-sandbox` implementation generates a temporary file in the `/var/sandbox/sandbox-python/` directory to save and execute your Python code. Before running your Python code, it uses `syscall.Chroot` to restrict the current process's root to the `/var/sandbox/sandbox-python/` directory. This directory structure, visible to the Python process, determines all the Python modules/packages that can be imported, including modules based on C code.

- Root: `/var/sandbox/sandbox-python/` is the root directory from the Python process perspective. Its subdirectories depend on the `python_lib_path` configuration in your `config.yaml`. Usually, it includes:
  - `etc/` directory
  - `python.so` shared object, compiled and built by `dify-sandbox`
  - `usr/lib` directory
  - `usr/local/`

If you haven't configured `python_lib_path`, `dify-sandbox` will default to the following settings (see code [internal/static/config_default_amd64.go](https://github.com/langgenius/dify-sandbox/blob/main/internal/static/config_default_amd64.go); for ARM systems, see `config_default_arm64.go`):

```go
var DEFAULT_PYTHON_LIB_REQUIREMENTS = []string{
    "/usr/local/lib/python3.10", // Usually your Python installation directory; if using conda, modify this to the conda virtual environment root directory, e.g., /root/anaconda3/envs/{env_name}
    "/usr/lib/python3.10",
    "/usr/lib/python3",
    "/usr/lib/x86_64-linux-gnu/libssl.so.3", // Your Python code's shared object dependency; it will be copied to /var/sandbox/sandbox-python/usr/lib/x86_64-linux-gnu/, and your Python process will load it from /usr/lib/x86_64-linux-gnu/
    "/usr/lib/x86_64-linux-gnu/libcry