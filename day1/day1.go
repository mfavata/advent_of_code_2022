package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "sort"
)

func main() {

    var current, topthree int
    totals := []int{}
    var input []string

    readFile, err := os.ReadFile("day1.txt")
    if err != nil {
        fmt.Println(err)
    }

    input = strings.Split(string(readFile), "\n")

    for x := 0; x < len(input) - 1; x++ {
        i, _ := strconv.Atoi(input[x])
        if i != 0 {
            current += i 
        } else {
            totals = append(totals, current)
            current = 0
        }
    }

    sort.Ints(totals)

    for i := len(totals)-3; i < len(totals); i++ {
        topthree += totals[i]
    }

    fmt.Println("The highest calorie count is:", totals[len(totals)-1])
    fmt.Println("The three highest totals are:", totals[len(totals)-3:], "and their sum is:", topthree)
}
