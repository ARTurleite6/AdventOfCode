package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

    data := string(bytes)

    var heap []int

	sum := 0
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
            heap = append(heap, sum)
            sum = 0
		} else {
            value, err := strconv.Atoi(line)
            if err != nil {
                fmt.Println("Error parsing line " + line)
                return
            }
            sum += value
        }
	}

    sort.Ints(heap)

    fmt.Println(max)

}
