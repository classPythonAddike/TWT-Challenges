package main

import (
	"bufio"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

var out *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 10240) // 10 MB for buffer size
var ibytes []byte
var l, max int

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

func main() {
    ibytes, _ = ioutil.ReadAll(os.Stdin)
    max = len(ibytes) - 1

    T := scan()
    for i := 0; i < T; i++ {
        s := float64(scan())
        n := float64(scan())

        if n < 0 || n > (s * (s + 1)) / 2 {
            out.WriteString("BAD INPUT\n")
            continue
        } else if n == 0 {
            if s == 2 {
                out.WriteString("2 1\n")
                continue
            } else {
                out.WriteString("BAD INPUT\n")
                continue
            }
        }

        D := ((2 * s - 1) * (2 * s - 1)) - 8 * n + 8 * s
        i := ((2 * s - 1) - math.Sqrt(D)) / 2

        if math.Mod(i, 1.0) == 0.0 {
            i -= 1
        }

        row := int(math.Floor(i)) + 1
        c := ((row + 1) * (2 * int(s) - row)) / 2
        column := int(s) - (c - int(n))
        
        out.WriteString(strconv.Itoa(row + 1))
        out.WriteByte(' ')
        out.WriteString(strconv.Itoa(column))
        out.WriteByte(10)
    }
    out.Flush()
}
