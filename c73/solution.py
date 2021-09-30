from typing import List

def distance_to_lobby(row_len, row, column, lobby_pos):
	if lobby_pos == "R":
		return row_len - column
	if lobby_pos == "L":
		return int(column)
	if lobby_pos == "U" and row == 0 or lobby_pos == "D" and row == 1:
		return min(row_len - column, column)
	else:
		return 1 + min(row_len - column, column)

def solution(num_houses: int, lobby: str, row_1: List[int], row_2: List[int]) -> int:

	distances = {}
	# {"row pos": "distance_from_tekgar_to_trash distance_from_trash_to_lobby"}

	for i in range(len(row_1)):
		distances[f"0 {i}"] = distance_to_lobby(num_houses, lobby, 0, int(i + 1))

	print(row_1, "\n" + str(row_2))
	print(distances)

	return 0

solution(
	5,
	"R",
	[1, 0, 0, 2, 1],
	[1, 0, 1, 0, 1]
)