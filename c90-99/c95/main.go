package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

var out *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 10240)
var primes [75000]int = [75000]int{}
var ibytes []byte
var l, max int

const n int = 951162 // 951161 is 75000th prime
const n_lim int = 31701 // Just eyeballed this value, was too lazy to properly solve inequalities

func scan() int {
    b := ibytes[l]

    // Check if the byte is not a number
    for b < 48 || b > 57 {
        l++
        b = ibytes[l]
    }

    result := 0
    for 47 < b && b < 58 {
        result = (result << 3) + (result << 1) + int(b-48)

        l++
        if l > max {
            return result
        }
        b = ibytes[l]
    }
    return result
}

func generatePrimes() {
    numbers := [951162]bool{false, false, true, true}

    for i := 1; i <= n / 6; i++ {
        j := 6 * i + 1
        if j < n { numbers[j] = true }
        numbers[6 * i - 1] = true
    }

    for d := 5; d <= n_lim; d += 6 {
        if d*d < n {
            numbers[d*d] = false
        }

        for x := 1; x <= n_lim; x++ {
            y := 6 * d * x + d

            if y < n {
                numbers[y] = false
            } else {
                break
            }

            y = 6 * d * x + d * d
            if y < n { numbers[y] = false }
        }
        
        g := d + 2

        if g*g < n {
            numbers[g*g] = false
        }

        for x := 1; x <= n_lim; x++ {
            y := 6 * g * x + g * g - 2 * g
            if y < n {
                numbers[y] = false
            } else {
                break
            }

            y = 6 * g * x + g * g
            if y < n { numbers[y] = false }
        }
    }

    pointer := 0
    for i := 0; i < n; i++ {
        if numbers[i] {
            primes[pointer] = i
            pointer++
        }
    }
}

func main() {
    generatePrimes()
    
    ibytes, _ = ioutil.ReadAll(os.Stdin)
    max = len(ibytes) - 1

    T := scan()
    for i := 0; i < T; i++ {
        out.WriteString(strconv.Itoa(primes[scan() - 1]))
        out.WriteByte(10)
    }
    out.Flush()
}
