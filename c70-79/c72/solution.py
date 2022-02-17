for i in range(int(input())):
	n = input()

	if len(n) == 1:
		print(False)
		break


	for l in range(1, len(n) // 2 + 1):
		s, e = int(n[:l]), int(n[-l:])

		if len(n) % l:
			continue

		is_ = False

		for w in range(s, s + len(n) // l):
			
			if str(w) != n[(w - s) * l:(w - s) * l + l]:
				is_ = False
				break

			is_ = True

		if is_:
			print(True)
			break

	else:
		print(False)