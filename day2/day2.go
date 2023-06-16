package main

import (
    "fmt"
    "os"
    "strings"
    "sort"
)

func rps(elf1 int, elf2 int) int {

    var score int

    if elf1-elf2 == 1 || elf2-elf1 == 2 {
        score = elf2 
    } else if elf2-elf1 == 1 || elf1-elf2 == 2 {
        score = elf2 + 6 
    } else {
        score = elf2 + 3
    }

    return score

}

func determine(elf1 int, elf2 int) int {
    
    var choice int 

    if elf2 == 1 { 
        if elf1 > 1 {
            choice = elf1 - 1
        } else {
            choice = elf1 + 2
        }
    } else if elf2 == 3 {
        if elf1 == 3 {
            choice = elf1 - 2
        } else {
            choice = elf1 + 1
        } 
    } else {
        choice = elf1
    }

    return choice
}

func main() {

    var totalscore1, totalscore2 int
    var rpsgame []string
    opponent := []string{"","A","B","C"}
    player := []string{"","X","Y","Z"}

    input, err := os.ReadFile("day2input.txt")
    if err != nil {
        fmt.Println(err)
    }

    rpsgame = strings.Split(string(input), "\n")

    for x := 0; x < len(rpsgame) - 1; x++ {

        a := strings.Split(rpsgame[x], " ")
        elf1 := sort.SearchStrings(opponent, a[0])
        elf2 := sort.SearchStrings(player, a[1])

        totalscore1 += rps(elf1, elf2)
        totalscore2 += rps(elf1, determine(elf1, elf2))
    }

    fmt.Println("Your total score for challenge 1 is:", totalscore1)
    fmt.Println("Your total score for challenge 2 is:", totalscore2)

}
