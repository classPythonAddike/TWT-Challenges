import sys

sys.setrecursionlimit(10**6)


def nearby_water(matrix, x, y) -> list:
    waters = []
    neighbours = [
        (y - 1, x),
        (y + 1, x),
        (y, x - 1),
        (y, x + 1)
    ]

    if matrix[y][x] == 0:
        return []

    for Y, X in neighbours:
        if matrix[Y][X] == 1:
            waters.append((Y, X))

    return waters


def get_river_count(pos_x, pos_y, matrix, passed_positions):
    for water_y, water_x in nearby_water(matrix, pos_x, pos_y):
        if (water_y, water_x) not in passed_positions:
            passed_positions.append((water_y, water_x))
            passed_positions = get_river_count(
                water_x,
                water_y,
                matrix,
                passed_positions
            )
    return passed_positions


for i in range(int(input())):
    width, height = map(int, input().split())

    matrix = [[0] * (width + 2)]

    for i in range(height):
        matrix += [[0] + [*map(int, input().split())] + [0]]

    matrix += [[0] * (width + 2)]

    global_seen = []
    river_lengths = []

    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            if matrix[y][x] and (y, x) not in global_seen:
                river = get_river_count(x, y, matrix, [(y, x)])
                global_seen += river
                river_lengths += [len(river)]

    print(" ".join(map(str, sorted(river_lengths))))

