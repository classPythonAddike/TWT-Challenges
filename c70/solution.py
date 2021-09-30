m = {
	chr(65 + e): e for e in range(26)
}
m.update(
	{
		j:e + 52 for e, j in enumerate("0123456789+/")
	}
)
m.update(
	{
		chr(97 + e): e + 26 for e in range(26)
	}
)

for i in range(int(input())):
	k = "".join([f"{m[j]:06b}" for j in input().strip("=")])
	l = len(k) // 8
	k = k[:l * 8]
	k = [chr(int(k[j:j+8],2)) for j in range(0,len(k),8)]
	print(*k, sep="")