package main

import (
	"bufio"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var out *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 20480) // Store 20 MB of output in buffer before flushing
var ibytes []byte
var l, max int

func scan() uint32 {
    b := ibytes[l]

    // Check if the ibytes is not a number
    for b < 48 || b > 57 {
        l++
        b = ibytes[l]
    }

    result := uint32(0)
    for 47 < b && b < 58 {
        result = (result << 3) + (result << 1) + uint32(b-48)

        l++
        if l > max {
            return result
        }
        b = ibytes[l]
    }
    return result
}

func encodeInt(n uint32) string {
    str := strconv.Itoa(int(n))
    var builder strings.Builder
    var dig_map map[rune]int8 = map[rune]int8{'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0}

    for _, dig := range str {
        dig_map[dig] += 1
    }

    for _, dig := range "9876543210" {
        count := dig_map[dig]
        for i := int8(0); i < count; i++ {
            builder.WriteRune(dig)
        }
    }

    return builder.String()
}

// This is very slow (takes 1s) but I ran out of time :sadcat:
func generatePrimes(n uint32) map[string][7]byte {
    encoded_primes := map[string][7]byte{"2": {50}, "3": {51}}
    primes := []uint32{2, 3}

    for i := uint32(5); i <= n; i += 2 {
        lim := uint32(math.Sqrt(float64(i)))
        isPrime := true
        
        for _, j := range primes {
            if j > lim {
                break
            } else if i % j == 0 {
                isPrime = false
                break
            }
        }
        
        if isPrime {
            primes = append(primes, i)

            dig_len := 0
            for j := i; j > 0; j /= 10 {
                dig_len += 1
            }
            num_bytes := [7]byte{}
            j := i
            for l := dig_len - 1; l > -1; l-- {
                num_bytes[l] = byte(48 + j % 10)
                j /= 10
            }
            
            encoded_primes[encodeInt(i)] = num_bytes
        }
    }
    return encoded_primes
}

func main() {
    encoded_primes := generatePrimes(10000000)

    ibytes, _ = ioutil.ReadAll(os.Stdin)
    max = len(ibytes) - 1
    T := scan()

    var testcases []uint32 = make([]uint32, T)
    for i := uint32(0); i < T; i++ {
        testcases[i] = scan()
    }

   for _, num := range testcases {
        for _, b := range encoded_primes[encodeInt(num)] {
            if b == 0 {
                break
            }
            out.WriteByte(b)
        }
       out.WriteByte(10)
   }
   out.Flush()
}
