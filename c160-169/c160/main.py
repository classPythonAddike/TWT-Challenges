for i in range(int(input())):
	w = input()

	if w == "HelloWorld":
		print("!!!")
	elif w in "HelloWorld" and w[0] == "H":
		print("HelloWorld"[len(w):])
	else:
		print("???")
