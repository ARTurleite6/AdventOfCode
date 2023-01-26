package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Play int

const (
	Paper    Play = 0
	Rock     Play = 1
	Scissors Play = 2
)

type Round struct {
	first  Play
	second Play
}

func roundFromString(line string) Round, err {
    values := strings.Split(line, " ")
    first, err1 := strconv.Atoi(values[0])
    second, err2 := strconv.Atoi(values[1])
    if err != nil {
    }
    return Round {
        first: values[0],
        second: values[1],
    }
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading from file" + os.Args[1])
		return
	}

	buffer := string(bytes)

	for _, line := range strings.Split(buffer, "\n") {

	}
}
