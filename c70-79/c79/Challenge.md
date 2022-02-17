# Challenge 79: Spontaneous interview | A programming journey #1

## Task
You are given a number T the number of tests cases. T cases follow. For each case you are given a string and a character on the next line. Encode the string using the single character XOR cipher with the given character as your cipher key.

## Example
### Input
```
2
aaa
5
bye
7
```
### Output
```
TTT
UNR
```

### Explanation
Let's say you have a random string of any length, and a single character we'll call a key. Split the string into individual characters then take each character separately and perform a XOR operation on it with the key you have, then add the resulting character into another string. That resulting string is what you need to print, and the whole process is called single character XOR encoding.

## Submissions and grading
Code can be written in either of these languages:
```
- Python 3.9
- C++ 10.3
- Ruby 3.0
- Golang 1.16
```
