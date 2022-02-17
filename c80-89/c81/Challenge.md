# Challenge 81: Permutations

## Task
You are given a number T and T testcases follow. For each testcase, you are given a number n. Print the number of permutations (not compulsory to use all digits of n in a permutation) which are even.

## Examples

### Input
```
2
223
458
```
### Output
```
4
6
```

### Explanation
1. The different [non repeating] permutations of N, where the permutation could even be only one digit (this digit d has to be a part of n) are:
2
22
223
232
322

Only 4 of these values are even ie: 2, 22, 232, 322

2. The different [non repeating] permutations of N, where the permutation could even be only one digit (this digit d has to be a part of n) are:
4
45
54
458
485
548
584
845
854

Out of these only 6 values are even ie: 4, 54, 458, 548, 584, 854

Formally, a permutation will be all combinations of
```
arr[0:1]
arr[0:2]
arr[0:3]
......
arr[0:len(arr)-1]
```

## Submissions and grading
Code can be written in either of these languages:
```
- Python 3.9
- C 10.3
- C++ (G++ 10.3)
- Ruby 3.0
- Golang 1.16
```
