package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func sliceContains(sl []string, name string) bool {
    for _, value := range sl {
        if value == name {
            return true
        }
    }
    return false
}

func main() {
    file, _ := os.ReadFile("input.txt")
    input := strings.Split(string(file), "\n")
    
    var files = make(map[string]int)
    var directorySizes = make(map[string]int)
    var currentDir, directories []string
    var part1solution, multipleDir, fileSystemSize int

    for x := 0; x < len(input)-1; x++ {
        working := strings.Split(input[x], " ")

        if working[0] == "$" {
            if working[1] == "cd" {
                switch working[2] {
                    case "..":
                        currentDir = currentDir[0:len(currentDir)-1]
                    default:
                        if sliceContains(directories,working[2]) {
                            multipleDir += 1
                            currentDir = append(currentDir, working[2]+strconv.Itoa(multipleDir))
                            directories = append(directories, working[2]+strconv.Itoa(multipleDir))
                        } else {
                            currentDir = append(currentDir, working[2])
                            directories = append(directories, working[2])
                        }
                }
            }
        } else {
            switch working[0] {
                case "dir":
                    continue
                default:
                    fileName := strings.Join(currentDir, "/") + "/" + working[1]
                    files[fileName], _ = strconv.Atoi(working[0])
            }
        }
    }

    for y := 0; y < len(directories); y++ {
        directorySize := 0
        for file, size := range files {
            splitFile := strings.Split(file, "/")
            for x := 0; x < len(splitFile); x++ {
                if splitFile[x] == directories[y] {
                    directorySize += size
                }
            }
        }

        directorySizes[directories[y]] = directorySize
        if directorySize <= 100000 {
            part1solution += directorySize
        }
    }

    for _, value := range files {
        fileSystemSize += value
    }

    currentFree := 70000000-fileSystemSize
    minimumToDelete := 30000000-currentFree
    currentValue := fileSystemSize

    for _, value := range directorySizes {
        if value >= minimumToDelete && value < currentValue {
            currentValue = value
        }
    }
    fmt.Println("Solution Part 1:", part1solution)
    fmt.Println("Solution Part 2:", currentValue)
}
