package main

import "fmt"

type Monkey struct {
  Id    int
  Op    int
  Test  int
  True  int
  False int
}

func operationMul(list []int, op int) []int {
  for x := 0; x < len(list); x++ {
    if op > 0 {
      list[x] = ((list[x] * op) / 3)
    } else {
      list[x] = ((list[x] * list[x]) / 3)
    }
  }
  return list
}

func operationAdd(list []int, op int) []int {
  for x := 0; x < len(list); x++ {
    list[x] = ((list[x] + op) / 3)
  }
  return list
}

func monkeyRound(count map[int]int, mnky Monkey, list map[int][]int) {
  var roundList []int
  switch mnky.Id {
  case 0, 4, 7:
    roundList = operationMul(list[mnky.Id], mnky.Op)
  case 1, 2, 3, 5, 6:
    roundList = operationAdd(list[mnky.Id], mnky.Op)
  }

  for x := 0; x < len(roundList); x++ {
    if roundList[x]%mnky.Test == 0 {
      list[mnky.True] = append(list[mnky.True], roundList[x])
    } else {
      list[mnky.False] = append(list[mnky.False], roundList[x])
    }
    count[mnky.Id]++
  }
  list[mnky.Id] = []int{}
}

func main() {
  var Items = make(map[int][]int)
  var countList = make(map[int]int)
  zero := Monkey{0, 5, 3, 7, 4}
  one := Monkey{1, 6, 17, 3, 0}
  two := Monkey{2, 5, 2, 3, 1}
  three := Monkey{3, 2, 19, 7, 0}
  four := Monkey{4, 7, 11, 5, 6}
  five := Monkey{5, 7, 5, 2, 1}
  six := Monkey{6, 1, 13, 5, 2}
  seven := Monkey{7, 0, 7, 4, 6}
  Items[0] = []int{66, 71, 94}
  Items[1] = []int{70}
  Items[2] = []int{62, 68, 56, 65, 94, 78}
  Items[3] = []int{89, 94, 94, 67}
  Items[4] = []int{71, 61, 73, 65, 98, 98, 63}
  Items[5] = []int{55, 62, 68, 61, 60}
  Items[6] = []int{93, 91, 69, 64, 72, 89, 50, 71}
  Items[7] = []int{76, 50}
  monkeyList := []Monkey{zero, one, two, three, four, five, six, seven}

  for i := 0; i < 20; i++ {
    for x := 0; x < len(monkeyList); x++ {
      monkeyRound(countList, monkeyList[x], Items)
    }
  }
  fmt.Println(Items)
  fmt.Println(countList)
}
