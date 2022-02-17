# Challenge 77: kickstart

## Task

You will be given an unsigned integer T, indicating number of tests you will receive. For each of the tests, you will be given one line containing one 64-bit unsigned integer N.
Given the integer N, find the last non-zero digit of N!, where N! means the factorial of N.
For example, given the number 7, we need to find the last non-zero digit of 7!. 7! = 7 * 6 * 5 * 4 * 3 * 2 * 1 = 5040, so the result is 4.

### Constraints:
0 <= N <= 2^64 - 1 (18,446,744,073,709,551,615)
(N is any unsigned 64-bit integer)

## Examples:

### Input:
```
3
7
19
98
```

### Output:
```
4
2
6
```

7! = 50**4**0
19! = 12164510040883**2**000

## Hint
Since the input can go as large as 2^64 - 1, it is not possible to actually calculate the result of N! and get its last digit.
