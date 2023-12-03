from part_01 import part_01
from part_02 import part_02


with open("input.txt") as input_file:
    input_str = input_file.read()
    print(part_01(input_str))
    print(part_02(input_str))
