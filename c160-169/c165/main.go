package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var stdin *bufio.Scanner = bufio.NewScanner(os.Stdin)
var stdout *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024*512)

func in(a string, l []string) bool {
	for _, i := range l {
		if a == i {
			return true
		}
	}
	return false
}

func solve(snowballs [][]string) {
	for index, traits := range snowballs {
		for i := 0; i < len(traits); i++ {
			for j := i + 1; j < len(traits); j++ {
				var onlyA, onlyB bool = false, false
				for num, snowball := range snowballs {
					if num == index {
						continue
					}

					aIn, bIn := in(traits[i], snowball), in(traits[j], snowball)

					onlyA = onlyA || (aIn && !bIn)
					onlyB = onlyB || (bIn && !aIn)

					if onlyA && onlyB {
						stdout.WriteString("NO\n")
						return
					}
				}
			}
		}
	}
	stdout.WriteString("YES\n")
}

func main() {
	stdin.Split(bufio.ScanLines)
	stdin.Scan()

	var N int
	var snowballs [][]string

	for stdin.Scan() {
		N, _ = strconv.Atoi(stdin.Text())
		snowballs = make([][]string, N)
		for j := 0; j < N; j++ {
			stdin.Scan()
			snowballs[j] = strings.Split(stdin.Text(), " ")
		}
		solve(snowballs)
	}

	stdout.Flush()
}
