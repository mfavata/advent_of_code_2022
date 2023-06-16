package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
)

func main() {

    var stack = make(map[int]string)
    inputfile, _ := os.ReadFile("input.txt")

    instruction := strings.Split(string(inputfile), "\n")

//    [C]         [Q]         [V]    
//    [D]         [D] [S]     [M] [Z]
//    [G]     [P] [W] [M]     [C] [G]
//    [F]     [Z] [C] [D] [P] [S] [W]
//[P] [L]     [C] [V] [W] [W] [H] [L]
//[G] [B] [V] [R] [L] [N] [G] [P] [F]
//[R] [T] [S] [S] [S] [T] [D] [L] [P]
//[N] [J] [M] [L] [P] [C] [H] [Z] [R]
// 1   2   3   4   5   6   7   8   9 

    stack[1] = "NRGP"
    stack[2] = "JTBLFGDC"
    stack[3] = "MSV"
    stack[4] = "LSRCZP"
    stack[5] = "PSLVCWDQ"
    stack[6] = "CTNWDMS"
    stack[7] = "HDGWP"
    stack[8] = "ZLPHSCMV"
    stack[9] = "RPFLWGZ"

    for crane := 0; crane < len(instruction)-1; crane++ {

        current := strings.Split(instruction[crane], " ")
        numOfCrates, _ := strconv.Atoi(string(current[1]))
        fromCrate, _ := strconv.Atoi(string(current[3]))
        toCrate, _ := strconv.Atoi(string(current[5]))
        fmt.Println("Move", numOfCrates, "from", fromCrate, "to", toCrate)
        fmt.Println("From:", stack[fromCrate], "To:", stack[toCrate])

        movingCrates := stack[fromCrate]

//        for moved := 1; moved <= numOfCrates; moved++ {
//            movingCrates := stack[fromCrate]
//            stack[toCrate] = stack[toCrate] + string(movingCrates[len(movingCrates)-moved])
//        }

        stack[toCrate] = stack[toCrate] + string(movingCrates[len(movingCrates)-numOfCrates:])
        stack[fromCrate] = movingCrates[:len(movingCrates)-numOfCrates]

        fmt.Println("Moved", stack)
    }
}
