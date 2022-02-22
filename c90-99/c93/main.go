package main

import (
    "bufio"
    "io/ioutil"
    "os"
    "strconv"
)

var out *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 10240) // Store 10 MB of output in buffer before flushing
var bytes []byte
var l, max int

// Fast way to read input
// https://gist.github.com/AkashBabu/938909700e5604483e2b44815a3e2c11
func scan() uint32 {
    b := bytes[l]

    // Check if the bytes is not a number
    for b < 48 || b > 57 {
        l++
        b = bytes[l]
    }

    result := uint32(0)
    for 47 < b && b < 58 {
        result = (result << 3) + (result << 1) + uint32(b-48)

        l++
        if l > max {
            return result
        }
        b = bytes[l]
    }
    return result
}

func solve(N, K, B uint32, lights []int8) string {
    var previous_win, current_min, num, i uint32 = 0, 0, 0, 1
    for ; num < K; num++ {
        previous_win += uint32(lights[num])
    }
    current_min = previous_win

    // K = 6, N = 14
    //                      1 1 1 1 1
    //  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4
    // [0 1 0 1 0 1 0 1 1 1 0 1 0 1 1]
    //  ^         ^ - Window 0
    //                    ^         ^ - Window 9 = N - K + 1
    // Ans: 3 (0 to 5)
    for ; i <= N - K; i++ {
        previous_win -= uint32(lights[i - 1] - lights[i + K - 1])
        if previous_win < current_min {
            current_min = previous_win
        }
    }

    return strconv.Itoa(int(current_min))
}

func main() {
    bytes, _ = ioutil.ReadAll(os.Stdin)
    max = len(bytes) - 1

    T := scan()

    var i, j uint32 = 0, 0

    for ; i < T; i++ {
        n := scan()
        k := scan()
        b := scan()

        lights := make([]int8, n)
        for j = 0; j < b; j++ {
            lights[scan() - 1] = 1
        }
        out.WriteString(solve(n, k, b, lights))
        out.WriteByte('\n')
    }
    out.Flush()
}
