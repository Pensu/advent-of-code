package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var resCount int
	f, err := os.Open("values.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		charCount := 0
		line := strings.Split(s.Text(), " ")

		recurr := string(line[0])

		splitRecurr := strings.Split(recurr, "-")

		min, _ := strconv.Atoi(splitRecurr[0])
		max, _ := strconv.Atoi(splitRecurr[1])

		char := rune(line[1][0])

		passwd := string(line[2])

		for pos, c := range passwd {
			if c == char && (pos+1 == min || pos+1 == max) {
				charCount++
			}
		}

		if charCount == 1 {
			resCount++
		}
	}

	fmt.Println("Final result: ", resCount)

}
