package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
)

// make slice to represent CRT
// if drawCRT == true append #, else append .
//make sure to pass divideLines(cycle) into drawCRT as cycle to represent correct position online
//make displayCRT function that prints out each 40 character line of CRT slice

func divideLines(a int) int {
  if (a / 40) >= 1 {
    return (a - ((a / 40) * 40))
  }
  return a
}

func drawCRT(cycle int, X int) string {
  for sprite := -1; sprite < 2; sprite++ {
    if cycle == (X + sprite) {
      return "#"
    }
  }
  return "."
}

func displayCRT(a [240]string) {
  for x := 0; x < 240; x = x+40 {
    fmt.Println(a[x:x+40])
  }
}

func partOne(x map[int]int) int {
  var solution int
  for a := 20; a <= 220; a=a+40 {
    solution = solution + (x[a]*a)
  }
  return solution
}

func main() {
  file, _ := os.ReadFile("input.txt")
  instructions  := strings.Split(string(file), "\n")
  stack := make(map[int]int)
  X, cycle, line, exec := 1, 1, 0, 0
  var CRT [240]string

  for line < len(instructions) - 1 {
    CRT[cycle-1] = drawCRT(divideLines(cycle-1), X)
    fmt.Println("Cycle", cycle, "value of X is", X, "line is", line)
    stack[cycle] = X

    operation := strings.Split(instructions[line], " ")
    if operation[0] == "noop" {
      //fmt.Println("noop, continue")
      line++
    } else {
      if cycle == exec {
        value, _ := strconv.Atoi(operation[1])
        X = X + value
        line++
      } else {
        //fmt.Println("increase X value by:", operation[1], "at end of next cycle")
        exec = cycle + 1
      }
    }
    cycle++
  }

  fmt.Println("Part 1:", partOne(stack))
//  fmt.Println(CRT)
  displayCRT(CRT)
  fmt.Println(stack)
}
