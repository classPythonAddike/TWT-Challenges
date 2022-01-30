package main

import (
	"bufio"
	"os"
	"strconv"
)

// Read input and parse into two binary numbers (array of int8)
func input(reader *bufio.Reader) ([2050]int8, [2050]int8) {
    var num1, num2 [2050]int8
    line, _ := reader.ReadSlice('\n')
    
    stringPointer := len(line) - 2

    // Read backwards, upto first space
    for numPointer := 2049;; stringPointer-- {
        char := line[stringPointer]
        if char == ' ' {
            break
        }
        num1[numPointer] = int8(char - 48)
        numPointer--
    }

    stringPointer--

    // Read backwards, from the space to the beginning of the line
    for numPointer := 2049; stringPointer >= 0; stringPointer--{
        num2[numPointer] = int8(line[stringPointer] - 48)
        numPointer--
    }

    // Numbers look like - [0, 0, 0, 0, ....... 1, 1, 0, 1, 0, 0, 1, ....]
    return num1, num2
}

func main() {
    out := bufio.NewWriter(os.Stdout)
    inp := bufio.NewReaderSize(os.Stdin, 4100)

    rawT, _ := inp.ReadString('\n')
    T, _ := strconv.Atoi(rawT[:len(rawT) - 1])

    for i := 0; i < T; i++ {
        pointer := 2049
        result := [2050]byte{}
        // Don't need to print trailing zeroes, so we keep track of first 1 in the number
        firstOne := 2049
        // Carry over
        var carry int8 = 0

        for num1, num2 := input(inp); pointer >= 0; pointer-- {
            a := num1[pointer]
            b := num2[pointer]
        
            sum := a ^ b ^ carry // Use XOR instead of `+` for speed
            result[pointer] = byte(sum + 48)
            carry = (a & b) | (a & carry) | (b & carry) // Use binary OR and AND instead of if/switch for speed
            
            if sum == 1 { firstOne = pointer }
        }

        if carry == 1 {
            result[pointer - 1] = 49
            firstOne = 0
        }

        out.Write(result[firstOne:])
        out.WriteByte('\n')
    }
    out.Flush()
}
