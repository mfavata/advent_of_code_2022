package main

import (
    "fmt"
    "strings"
    "os"
)

func duplicateCheck(input string) bool {
    alphabet := "abcdefghijklmnopqrstuvwxyz"

    for x := 0; x < len(alphabet); x++ {
        if strings.Count(input, string(alphabet[x])) >= 2 {
            return false
            break
        }
    }

    return true
}

func main() {
    datastream, _ := os.ReadFile("input.txt")

    for x := 0; x < len(datastream) - 3; x++ {
        if duplicateCheck(string(datastream[x:x+4])) {
            fmt.Println("Part one:", x+4, string(datastream[x:x+4]))
            break
        }
    }

    for x := 0; x < len(datastream) - 13; x++ {
        if duplicateCheck(string(datastream[x:x+14])) {
            fmt.Println("Part two:", x+14, string(datastream[x:x+14]))
            break
        }
    }
}
