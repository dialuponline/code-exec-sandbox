import ctypes
import os
import sys
import json
import traceback

# setup sys.excepthook
def excepthook(type, value, tb):
    sys.stderr.write("".join(traceback.format_exception(type, value, tb)))
    sys.stderr.flush()
    sys.exit(-1)

sys.excepthook = excepthook

lib = ctypes.CDLL("/var/sandbox/sandbox-python/python.so")
lib.DifySeccomp.argtypes = [ctypes.c_uint32, ctypes.c_uint32, ctypes.c_bool]
lib.DifySeccomp.restype = None


import json
import os

import json
import sys
import traceback
import os

os.chdir("/var/sandbox/sandbox-python")

lib.DifySeccomp(65537, 1001, 1)

# declare main function here
def main() -> dict:
    original_strings_with_empty = ["apple", "", "cherry", "date", "", "fig", "grape", "honeydew", "kiwi", "", "mango", "nectarine", "orange", "papaya", "quince", "raspberry", "strawberry"