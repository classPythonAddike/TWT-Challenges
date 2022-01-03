# Challenge 86: Tic-tac-toe... with numbers

Let's solve a tic-tac-toe problem

## Task
You are given a number `n` and `n` testcases follow. For each testcase, you are given a starting player `X` or `O` and then the moves made between `X` and `O` alternatively.

So if the moves are `12569` and `X` starts,
```
1 - x
2 - o
5 - x
6 - o
9 - x
```

you have to output the winner(`X` or `O`) if there is any, otherwise output `Tie`.

The positions on the board are as follows
```
1 | 2 | 3
---------
4 | 5 | 6
---------
7 | 8 | 9
```

## Example
### Input:
```
2
X
1, 2, 5, 6, 9
X
1, 3, 2, 4, 6, 9, 7, 5, 8
```
### Output:
```
X
Tie
```

## Submissions
Code can be written in either of these languages:
```
- Python 3.9
- C++ (G++ 10.3)
- Ruby 3.0
- Golang 1.16
- Java 18
- Rust 1.52
```
