def solution(nums):
	num_dict = {i:0 for i in range(len(nums))} # Position:  number of times hit

	pos = 0

	while all([i < 2 for i in num_dict.values()]):
		pos += nums[pos]
		pos %= len(nums)
		num_dict[pos] += 1

		if all([i == 1 for i in num_dict.values()]):
			return True and nums != [0]

	return False



for i in range(int(input())):
	inp = list(map(int, input().split()))

	print(solution(inp))