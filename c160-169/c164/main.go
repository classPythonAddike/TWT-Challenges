package main

import (
	"bufio"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var stdout *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024*1024)
var stdin *bufio.Scanner = bufio.NewScanner(os.Stdin)

func readInput(K int) *[]int {
	var numbers []int = make([]int, K)
	for i := 0; i < K; i++ {
		stdin.Scan()
		numbers[i], _ = strconv.Atoi(stdin.Text())
	}

	sort.Slice(numbers, func(i, j int) bool { return numbers[i] > numbers[j] })
	return &numbers
}

/*
Generate combinations choosing r = 1 to K values from the input
Use principle of inclusion and exclusion
*/
func generateCombinations(numbers *[]int, K int, N int, caseNum int, outputChannel chan [2]int) {
	var numDivisible, numMultiples int
	var oddNumChosen int

	for r := 1; r < (1 << K); r++ {
		oddNumChosen = bits.OnesCount(uint(r))
		numMultiples = N

		for i := 0; i < K; i++ {
			if (r>>i)&1 == 1 {
				if numMultiples < (*numbers)[i] {
					numMultiples = 0
					break
				}
				numMultiples /= (*numbers)[i]
			}
		}

		if oddNumChosen&1 == 1 {
			numDivisible += numMultiples
		} else {
			numDivisible -= numMultiples
		}
	}

	outputChannel <- [2]int{caseNum, numDivisible}

}

func main() {
	stdin.Split(bufio.ScanWords)
	stdin.Scan()

	T, _ := strconv.Atoi(stdin.Text())

	var N, K int
	var output []int = make([]int, T)
	var outputChannel chan [2]int = make(chan [2]int, T)
	var testCase [2]int

	for i := 0; i < T; i++ {
		stdin.Scan()
		K, _ = strconv.Atoi(stdin.Text())
		stdin.Scan()
		N, _ = strconv.Atoi(stdin.Text())
		go generateCombinations(readInput(K), K, N, i, outputChannel)

	}

	for i := 0; i < T; i++ {
		testCase = <-outputChannel
		output[testCase[0]] = testCase[1]
	}

	for _, val := range output {
		stdout.WriteString(strconv.Itoa(val))
		stdout.WriteByte('\n')
	}

	stdout.Flush()
}
