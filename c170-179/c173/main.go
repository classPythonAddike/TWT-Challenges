package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
)

type Box struct {
    label int
    left *Box
    right *Box
    next *Box
    toppled bool
}

var input *bufio.Scanner = bufio.NewScanner(os.Stdin)
var output *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1024 * 256)

var boxes [1414][]Box;

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }
    
    construct_boxes()

    input.Split(bufio.ScanWords)
    input.Scan()
    T, _ := strconv.Atoi(input.Text())
    
    for i := 0; i < T; i++ {
        input.Scan()
        N, _ := strconv.Atoi(input.Text())
        box_row := row(N)
        box_column := N - (box_row - 1) * box_row / 2
        box := &boxes[box_row - 1][box_column - 1]
        fmt.Println(get_sum(box))
        clear_toppled(box)
    }
}

func row(n int) int {
    return int(math.Ceil((-1 + math.Sqrt(float64(8*n + 1))) / 2))
}

func construct_boxes() {
    var row, column int
    for row = 1; row <= 1414; row++ {
        boxes[row - 1] = make([]Box, row)
        for column = 1; column <= row; column++ {
            label := ((row - 1) * row) / 2 + column

            if label > 1000000 {
                return
            }

            var box Box;
            box.label = label
            box.toppled = false

            if (column == 1) {
                box.left = nil
            } else {
                box.left = &boxes[row - 2][column - 2]
            }

            if (column == row) {
                box.right = nil
            } else {
                box.right = &boxes[row - 2][column - 1]
            }

            boxes[row - 1][column - 1] = box
        }
    }
}

func get_sum(box *Box) int {
    if (box == nil) {
        return 0
    }

    if box.toppled {
        return get_sum(box.left) + get_sum(box.right)
    } else {
        box.toppled = true;
        return get_sum(box.left) + get_sum(box.right) + box.label
    }
}

func clear_toppled(box *Box) {
    box.toppled = false
    if box.left != nil { clear_toppled(box.left) }
    if box.right != nil { clear_toppled(box.right) }
}
