package main

import (
  "fmt"
  "strconv"
  "os"
  "strings"
)

func scenicScore(trees [99][99]int, x int, y int) int {
  var northScore, westScore, eastScore, southScore int

  for north := x-1; north >= 0; north-- {
    if trees[north][y] >= trees[x][y] {
      northScore = (x - north)
      break
    } else if north <= 0 {
      northScore = x - north
    }
  }

  for west := y-1; west >= 0; west-- {
    if trees[x][west] >= trees[x][y] {
      westScore = (y - west)
      break
    } else if west <= 0 {
      westScore = y - west
    }
  }

  for east := y+1; east <= 98; east++ {
    if trees[x][east] >= trees[x][y] {
      eastScore = (east - y)
      break
    } else if east >= 98 {
      eastScore = y - east
    }
  }

  for south := x+1; south <= 98; south++ {
    if trees[south][y] >= trees[x][y] {
      southScore = (south - x)
      break
    } else if south >= 98 {
      southScore = x - south
    }
  }

  return northScore * westScore * eastScore * southScore
}

func isVisible(trees [99][99]int, x int, y int) bool {
  north, south, east, west := true, true, true, true
  
  if x == 0 || y == 0 || x == 98 || y == 98 {
    return true
  }

  for col := 0; col < 99; col++ {
    for row := 0; row < 99; row++ {
      if col < x && trees[col][y] >= trees[x][y] {
        north = false
      } else if col > x && trees[col][y] >= trees[x][y] {
        south = false
      } else if row < y && trees[x][row] >= trees[x][y] {
        west = false
      } else if row > y && trees[x][row] >= trees[x][y] {
        east = false
      }
    }
  }
  if north || south || east || west {
    return true
  }

  return false
}

func createMultiArray(s []string) [99][99]int {
  var newArray [99][99]int

  for x := 0; x < len(s); x++ {
    for y := 0; y < len(s[x]); y++ {
      newArray[x][y], _ = strconv.Atoi(string(s[x][y]))
    }
  }

  return newArray
}

func main() {
  var visibleCount, scenicCounter, highestScore int
  input, _ := os.ReadFile("input.txt")
  lines := strings.Split(string(input), "\n")
  treeGrid := createMultiArray(lines)
  
  for x := 0; x < 99; x++ {
    for y := 0; y < 99; y++ {
      if isVisible(treeGrid, x, y) {
        visibleCount++
      }
      scenicCounter = scenicScore(treeGrid, x, y)
      if scenicCounter > highestScore {
        highestScore = scenicCounter
      }
    }
  }
  fmt.Println(visibleCount)
  fmt.Println(highestScore)
}
