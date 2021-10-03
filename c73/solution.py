def distance_to_lobby(row, pos, lobby_dir, row_length):
	return abs(
		0 if lobby_dir in "RL" else {"U": abs(1 - row), "D": 2 - row}[lobby_dir]
	) + (
		min(row_length - pos - 1, pos) if lobby_dir in "UD" else {"R": row_length - pos - 1, "L": pos}[lobby_dir]
	)

def distance_to_trash(trash_row, trash_pos, tek_row, tek_pos, row_length):
	return abs(
		trash_row - tek_row
	) + [
		abs(trash_pos - tek_pos),
		min(
			trash_pos + tek_pos,
			(2 * row_length) - trash_pos - tek_pos - 2
		)
	][abs(trash_row - tek_row)]

def s(lobby_dir, row1, row2):

	tekgar_row, tekgar_pos = 2, 0
	d_to_lobby, d_to_trash = {}, {}

	if "2" in row1:
		tekgar_row = 1
		tekgar_pos = row1.index("2")
	else:
		tekgar_pos = row2.index("2")

	r1_trash_dist = {
		i: distance_to_lobby(1, i, lobby_dir, len(row1))
		for i in range(len(row1)) if row1[i] == "1"
	}

	r2_trash_dist = {
		i: distance_to_lobby(2, i, lobby_dir, len(row2))
		for i in range(len(row2)) if row2[i] == "1"
	}

	r1_tek_dist = {
		i: distance_to_trash(1, i, tekgar_row, tekgar_pos, len(row1))
		for i in range(len(row1)) if row1[i] == "1"
	}

	r2_tek_dist = {
		i: distance_to_trash(2, i, tekgar_row, tekgar_pos, len(row2))
		for i in range(len(row2)) if row2[i] == "1"
	}

	for i in range(2):
		for pos in range(len(row1)):
			if [row1, row2][i][pos] == "1":
				d_to_lobby[f"{i + 1} {pos}"] = [r1_tek_dist, r2_tek_dist][i][pos] + \
					[r1_trash_dist, r2_trash_dist][i][pos]

	for i in range(2):
		for pos in range(len(row1)):
			if [row1, row2][i][pos] == "1":
				d_to_trash[f"{i + 1} {pos}"] = [r1_tek_dist, r2_tek_dist][i][pos]

	min_lobby_distance, min_trash_distance = [*d_to_lobby.values()][0], [*d_to_trash.values()][0]

	for i in d_to_lobby:
		if d_to_lobby[i] < min_lobby_distance:
			min_lobby_distance = d_to_lobby[i]
			min_trash_distance = d_to_trash[i]
		elif d_to_lobby[i] == min_lobby_distance:
			if d_to_trash[i] < min_trash_distance:
				min_trash_distance = d_to_trash[i]

	print(min_trash_distance)

for i in range(int(input())):
	_, l, r, b = [*map(input, "0000")]
	s(l, r, b)