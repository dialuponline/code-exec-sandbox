# Code-Exec-Sandbox
## Overview
Code-Exec-Sandbox offers a feasible method to execute untrusted code in a high-security environment. It's fashioned for use in an environment with multiple tenants, giving multiple users the ability to submit code for execution. Execution is conducted in a confined environment, limiting the system calls and resources the code can access.

## Guide
### Prerequisites
Code-Exec-Sandbox at the moment only extends support for Linux, as it's designed to work in docker containers. It requires the existence of the following dependencies:
- libsec