package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var stdin *bufio.Scanner = bufio.NewScanner(os.Stdin)
var stdout *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024*1024*2)

func readInput(N int) *[][2]int {

	var a, b int
	parties := make([][2]int, N)

	for i := 0; i < N; i++ {
		stdin.Scan()
		a, _ = strconv.Atoi(stdin.Text())
		stdin.Scan()
		b, _ = strconv.Atoi(stdin.Text())
		parties[i] = [2]int{a, b}
	}

	return &parties
}

func orderByEndTime(parties *[][2]int) {
	sort.Slice(
		*parties,
		func(i, j int) bool {
			return (*parties)[i][1] < (*parties)[j][1]
		},
	)
}

func chooseParties(parties [][2]int, N int) int {

	var chosen, pointer, lastPartyEndTime, currentEndTime int = 0, 0, 0, 0
	var changed, nonZero bool = false, false

	for pointer < N {

		nonZero = false
		changed = false

		if pointer == N-1 {
			if parties[pointer][0] >= lastPartyEndTime {
				chosen++
			}
			break
		}

		if parties[pointer][0] < lastPartyEndTime {
			pointer++
			continue
		}

		currentEndTime = parties[pointer][1]

		for parties[pointer][1] == currentEndTime {
			if parties[pointer][0] >= lastPartyEndTime {

				if parties[pointer][0] == parties[pointer][1] {
					chosen++
					changed = true
				} else {
					nonZero = true
				}
			}
			pointer++
		}

		if nonZero || changed {
			lastPartyEndTime = currentEndTime

			if nonZero {
				chosen++
			}
		}

	}

	return chosen
}

func main() {
	stdin.Split(bufio.ScanWords)
	stdin.Scan()

	var N int

	for stdin.Scan() {
		N, _ = strconv.Atoi(stdin.Text())
		parties := readInput(N)
		orderByEndTime(parties)
		stdout.WriteString(strconv.Itoa(chooseParties(*parties, N)))
		stdout.WriteByte('\n')
	}

	stdout.Flush()
}
