package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func intAbs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

type Coordinate struct {
  X int
  Y int
}

func movePos(x int) int {
  if x < 0 {
    return -1
  } else if x == 0 {
    return 0
  }

  return 1
}

func moveTail(head Coordinate, tail Coordinate) Coordinate {
  x := head.X - tail.X
  y := head.Y - tail.Y

  if intAbs(x) > 1 || intAbs(y) > 1 {
    return Coordinate{tail.X + movePos(x), tail.Y + movePos(y)}
  }
  return tail 
}

func main() {
  part1 := make(map[Coordinate]bool)
  part2 := make(map[Coordinate]bool)
  rope := make([]Coordinate, 10)

  file, _ := os.ReadFile("input.txt")
  lines := strings.Split(string(file), "\n")

  for x := 0; x < len(lines)-1; x++ {
    directions := strings.Split(lines[x], " ")
    distance, _ := strconv.Atoi(directions[1])
    for move := 0; move < distance; move++ {
      switch directions[0] {
        case "U":
          rope[0].Y++
        case "D":
          rope[0].Y--
        case "L":
          rope[0].X--
        case "R":
          rope[0].X++
        }

        for z := 1; z < 10; z++ {
          rope[z] = moveTail(rope[z-1], rope[z])
        }
        part1[rope[1]] = true
        part2[rope[9]] = true
    }
  }

  fmt.Println("Part 1:", len(part1), "Part 2:", len(part2))
}
