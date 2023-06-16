package main

import (
    "fmt"
    "strings"
    "os"
)

var priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func find_priority(a byte) int {
    i := strings.Index(priority, string(a)) + 1
    return i
}

func main() {
    var total int

    file, _ := os.ReadFile("input.txt")
    input := strings.Split(string(file), "\n")

    for y, _ := range input {
        rucksack := input[y]
        size := len(rucksack) / 2
        for x := 0; x < len(priority); x++ {
            if strings.Contains(rucksack[:size], string(priority[x])) && strings.Contains(rucksack[size:], string(priority[x])) {
                fmt.Println(string(priority[x]), find_priority(priority[x]))
                total += find_priority(priority[x])
            }
        }
    }

    fmt.Println("Total:", total)
}
