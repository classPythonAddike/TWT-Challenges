package main

import (
	"bufio"
	"math/bits"
	"os"
	"strconv"
)

var stdin *bufio.Scanner = bufio.NewScanner(os.Stdin)
var stdout *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 128)

func main() {
    var T int
    var n int64

    stdin.Split(bufio.ScanWords)

    stdin.Scan()
    T, _ = strconv.Atoi(stdin.Text())

    for i := 0; i < T; i++ {
        stdin.Scan()
        n, _ = strconv.ParseInt(stdin.Text(), 10, 64)

        if (n & 1 != 0 && n > 1) {
            stdout.WriteString("-1\n")
        } else {
            n += 2;
            if (n % 3 != 0) {
                stdout.WriteString("-1\n")
            } else {
                n /= 3;
                if (n & (n - 1) != 0) {
                    stdout.WriteString("-1\n")
                } else {
                    stdout.WriteString(strconv.Itoa(bits.Len(uint(n))))
                    stdout.WriteByte('\n')
                }
            }
        }
    }

    stdout.Flush()
}
