package main

import (
    "os"
	"fmt"
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

	max := 0
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			if sum > max {
                max = sum
			}
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

    fmt.Println(max)

}
