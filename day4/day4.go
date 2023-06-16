package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func sliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}    

func main() {

    var partOne, partTwo int

    readFile, _ := os.ReadFile("input.txt")
    inputLines := strings.Split(string(readFile), "\n")    

    for newLine := 0; newLine < len(inputLines) - 1; newLine++ {
        sections := strings.Split(inputLines[newLine], ",")
        elf1, _ := sliceAtoi(strings.Split(sections[0], "-"))
        elf2, _ := sliceAtoi(strings.Split(sections[1], "-"))

        if (elf1[0] <= elf2[0] && elf1[1] >= elf2[1]) || (elf1[0] >= elf2[0] && elf1[1] <= elf2[1]) {
            partOne++
            partTwo++
        } else if (elf2[0] <= elf1[1] && elf2[0] >= elf1[0]) || (elf1[0] <= elf2[1] && elf1[0] >= elf2[0]) {
            partTwo++
        }
    }

    fmt.Println("Complete Overlaps:", partOne)
    fmt.Println("Any Overlaps:", partTwo)
}
