package main

import (
    "bufio"
    "os"
    "strconv"
)

var n int = 100000
var palindromes [1099]int

var input *bufio.Scanner = bufio.NewScanner(os.Stdin)
var output *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 100)

func isPalindrome(num int) bool {
    var rem, rev int

    for temp := num; temp > 0; temp = temp / 10 {
        rem = temp % 10
        rev = rev*10 + rem
    }

    return rev == num
}

func solve(a, b, m int) int {
    count := 0

    for _, i := range palindromes {
        if i < a {
            continue
        } else if i > b {
            return count
        } else if isPalindrome(i / m) {
            count++
        }
    }

    return count
}

func main() {
    pointer := 0
    for i := 0; i < n; i++ {
        if isPalindrome(i) {
            palindromes[pointer] = i
            pointer++
        }
    }

    input.Split(bufio.ScanWords)
    input.Scan()

    var a, b, m int

    for input.Scan() {
        a, _ = strconv.Atoi(input.Text())
        input.Scan()
        b, _ = strconv.Atoi(input.Text())
        input.Scan()
        m, _ = strconv.Atoi(input.Text())
        
        output.WriteString(strconv.Itoa(solve(a, b, m)))
        output.WriteByte('\n')
    }

    output.Flush()
}
