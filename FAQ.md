## FAQ

### 1. Why does my Python code throw an exception like "xxx.so: cannot open shared object file: No such file or directory"?

This occurs because the `dify-sandbox` implementation generates a temporary file in the `/var/sandbox/sandbox-python/` directory to save and execute your Python code. Before running your Python code, it uses `syscall.Chroot` to restrict the current process's root to the `/var/sandbox/sandbox-python/` directory. This directory structure, visible to the Python process, determines all the Python modules/packages that can be imported, including modules based on C code.

- Root: `/var/sandbox/sandbox-python/` is the root directory from the Python process perspective. Its subdirectories depend on the `python_lib_path` configuration in your `config.yaml`. Usually, it includes:
  - `etc/` directory
  - `python.so` shared object, compiled and built by