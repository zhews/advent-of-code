from collections import defaultdict


def part_02(input_str: str) -> int:
    rows = [l for l in input_str.split("\n") if l != ""]
    checks = [
        (-1, 0),  # left
        (1, 0),  # right
        (0, -1),  # above
        (0, 1),  # below
        (-1, -1),  # up left
        (1, -1),  # up right
        (-1, 1),  # down left
        (1, 1),  # down right
    ]
    numbers = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]

    gears = defaultdict(list)
    for y in range(len(rows)):
        current_number = ""
        is_part_of_gear = False
        touched_gears = []
        for x in range(len(rows[y])):
            current_char = rows[y][x]
            is_number = current_char in numbers

            if is_number:
                current_number += current_char

                for cx, cy in checks:
                    check_y = y + cy
                    check_x = x + cx

                    if (
                        check_x < 0
                        or check_x > len(rows[y]) - 1
                        or check_y < 0
                        or check_y > len(rows) - 1
                    ):
                        continue

                    check_char = rows[check_y][check_x]
                    if check_char == "*":
                        is_part_of_gear = True
                        touched_gears.append((check_y, check_x))

            elif current_number != "":
                if is_part_of_gear:
                    for gear_y, gear_x in touched_gears:
                        gears[f"{gear_y},{gear_x}"].append(int(current_number))
                current_number = ""
                is_part_of_gear = False
                touched_gears = []

        if current_number != "":
            if is_part_of_gear:
                for gear_y, gear_x in touched_gears:
                    gears[f"{gear_y},{gear_x}"].append(int(current_number))
            current_number = ""
            is_part_of_gear = False
            touched_gears = []

    result = 0
    for gear in gears.values():
        gear = set(gear)
        gear = list(gear)
        if len(gear) != 2:
            continue
        result += gear[0] * gear[1]
    return result
