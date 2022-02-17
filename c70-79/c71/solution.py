for i in range(int(input())):
	a, b, _, c, d = input()
	print(
		abs(int(d) - int(b)) == abs(ord(a) - ord(c))
	)