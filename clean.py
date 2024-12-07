#!/bin/python3

import os
import subprocess
from os.path import join

if __name__ == "__main__":
    start_path = os.path.dirname(os.path.realpath(__file__))
    for root, dirs, files in os.walk(start_path):
        for name in dirs:
            subprocess.run(["go", "clean"], cwd=join(root,name))

