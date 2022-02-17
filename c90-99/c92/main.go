package main

import (
	"bufio"
	"os"
	"strconv"
)

var inp *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 2048) // 2 MB is the stdout buffer size

func main() {
    rawT, _, _ := inp.ReadLine()
    T, _ := strconv.Atoi(string(rawT))

    for i := 0; i < T; i++ {
        n, _, _ :=  inp.ReadLine()
        num, _ := strconv.Atoi(string(n))
        out.WriteString(strconv.Itoa((num * (num + 3) / 2) & 4095))
        out.WriteByte('\n')
    }
    out.Flush()
}
