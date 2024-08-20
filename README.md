# Code-Exec-Sandbox
## Overview
Code-Exec-Sandbox offers a feasible method to execute untrusted code in a high-security environment. It's fashioned for use in an environment with multiple tenants, giving multiple users the ability to submit code for execution. Execution is conducted in a confined environment, limiting the system calls and resources the code can access.

## Guide
### Prerequisites
Code-Exec-Sandbox at the moment only extends support for Linux, as it's designed to work in docker containers. It requires the existence of the following dependencies:
- libseccomp
- pkg-config
- gcc
- golang 1.20.6

### Procedure
1. Clone the repository using `git clone https://github.com/dialuponline/code-exec-sandbox` and navigate to the project directory.
2. Execute ./install.sh to install the required dependencies.
3. Use ./build/build_[amd64|arm64].sh command to compile the sandbox binary.
4. E