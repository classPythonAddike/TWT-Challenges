package main

import (
	"bufio"
	"strconv"
	"os"
)

var inp *bufio.Scanner = bufio.NewScanner(os.Stdin)
var out *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 2048) // 2 MB Buffer Size

func main() {
	inp.Split(bufio.ScanWords)
	inp.Scan()

	for inp.Scan() {
		a, _ := strconv.Atoi(inp.Text())
		inp.Scan()
		b, _ := strconv.Atoi(inp.Text())

		switch a > b {
		case true:
			out.WriteString(strconv.Itoa(2 * b - 1))
		case false:
			out.WriteString(strconv.Itoa(2 * a - 1))
		}

		out.WriteByte('\n')
	}

	out.Flush()
}