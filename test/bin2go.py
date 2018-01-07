#!/usr/bin/env python

import os
import sys

BYTES_PER_LINE = 16

asm_file = os.path.dirname(__file__) + "/cpudiag.bin"
with open(asm_file, "rb") as fin:
    line_num = 0
    sys.stdout.write('package main\n')
    sys.stdout.write('\n')
    sys.stdout.write('var TestRom []uint8 = []uint8{\n')
    sys.stdout.write('    ')
    while True:
        byte = fin.read(1)
        if byte == "":
            break
        sys.stdout.write('0x{:02x}, '.format(ord(byte)))
        line_num += 1
        if (line_num % BYTES_PER_LINE) == 0:
            sys.stdout.write('\n    ')
    sys.stdout.write('\n}\n')
