package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

    if len(os.Args) == 1 {
        fmt.Println("You need to pass what part of the exercise you want [part-1, part-2]")
        return
    }

    bytes, err := os.ReadFile("./content.txt")
    if err != nil {
        fmt.Println("Error reading from file")
        return
    }
    data := string(bytes)
    lines := strings.Split(data, "\n")
    values := make([]int, len(lines))

    sum := 0
    index := 0
    for _, line := range lines {
        if len(line) == 0 {
            values[index] = sum
            index += 1
            sum = 0
        } else {
            value, err := strconv.Atoi(line)
            if err != nil {
                fmt.Println("Error converting line to integer")
                return
            }
            sum += value
        }
    }

    part := os.Args[1]
    sort.Ints(values)
    result := 0
    if part == "part-1" {
        result = values[len(values)-1]
    } else {
        values_wanted := values[len(values) - 3:]
        for _, value := range values_wanted {
            result += value
        }
    }
    fmt.Println(result)
}
