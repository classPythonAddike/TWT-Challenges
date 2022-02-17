package main

import (
    "bufio"
    "os"
    "strconv"
    "sync"
)

type testcases struct {
    lock sync.Mutex
    cases map[int]string
}

var inp *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

var t map[int]string = map[int]string{}

type solution struct {
    case_num int
    soln string
}

// Find minimum of two integers, as `math.Min` requires floats
func min(val1, val2 int) int {
    if val1 < val2 {
        return val1
    }
    return val2
}

// Find the largest possible area (brute force)
func solve(heights []int) string {
    max_area := 0

    for i := 0; i < len(heights); i++ {
        for j := i + 1; j < len(heights); j++ {
            min_height := min(heights[i], heights[j])
            area := (j - i) * min_height
            if area > max_area {
                max_area = area
            }
        }
    }
    
    return strconv.Itoa(max_area)
}

func main() {
    rawT, _, _ := inp.ReadLine()
    T, _ := strconv.Atoi(string(rawT))

    var wg sync.WaitGroup
    wg.Add(T)

    // Channel to store results
    solutions_channel := make(chan solution, T)
    
    for i := 0; i < T; i++ {
        n, _ :=  inp.ReadString('\n')
        num, _ := strconv.Atoi(n[:len(n) - 1])
        heights := make([]int, num)

        // Read and parse input
        for j := 0; j < num - 1; j++ {
            rawNum, _ := inp.ReadString(' ')
            heights[j], _ = strconv.Atoi(rawNum[:len(rawNum) - 1])
        }
        rawNum, _ := inp.ReadString('\n')
        heights[num - 1], _ = strconv.Atoi(rawNum[:len(rawNum) - 1])
    
        // Run solve function in a goroutine and store answer in channel
        go func(input []int, case_number int, wg *sync.WaitGroup) {
            defer wg.Done()
            solutions_channel <- solution{case_num: case_number, soln: solve(input)}
        }(heights, i, &wg)
    }

    wg.Wait()
    close(solutions_channel)

    // Move the results from channel to dictionary
    for sol := range solutions_channel {
        t[sol.case_num] = sol.soln
    }

    // Write results to buffer
    for i := 0; i < T; i++ {
        out.Write([]byte(t[i]))
        out.WriteByte('\n')
    }
    out.Flush()
}
