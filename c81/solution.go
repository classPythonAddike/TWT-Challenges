package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

/*
    Only works for numbers upto ~20-25 digits
    See `solution.py` for a python implementation, which works for large numbers (Python does not set limits on max int size)
*/

func factorial(n uint64) uint64 {
    var r uint64 = 1
    for n > 1 {
        r *= n
        n -= 1
    }
    return r
}

func num_even_permutations(n string) uint64 {
    
    if len(n) == 1 {
        x, _ := strconv.Atoi(n)
        return uint64((x + 1) % 2)
    }

    digits := map[int]uint64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0}

    var odd_num_count, even_num_count, zero_count, denominator uint64 = 0, 0, 0, 1

    for _, dig := range n {
        int_dig, _ := strconv.Atoi(string(dig))
        digits[int_dig] += 1

        if int_dig % 2 == 0 {
            even_num_count += 1
        }
        if int_dig == 0 {
            zero_count += 1
        }
        if int_dig % 2 == 1 {
            odd_num_count += 1
        }
    }

    for _, i := range digits { denominator *= factorial(i) }

    odd_number_starting := odd_num_count * even_num_count * factorial(uint64(len(n)) - 2)
    even_number_starting := (even_num_count - zero_count) * (even_num_count - 1) * factorial(uint64(len(n)) - 2)

    // odd + even as starting digit
    numerator :=  odd_number_starting + even_number_starting

    ans := numerator / denominator

    return ans
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanner.Text()

    for scanner.Scan() {
        n := scanner.Text()

        current_num := ""
        var sum uint64 = 0

        for i := 0; i < len(n); i++ {
            current_num += string(n[i])
            sum += num_even_permutations(current_num)
        }

        fmt.Println(sum)
    }
}
