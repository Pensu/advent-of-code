package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("value.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var numbers []int

	s := bufio.NewScanner(f)
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		numbers = append(numbers, n)
	}

	// fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i], numbers[j], numbers[k])
					fmt.Println(numbers[i] * numbers[j] * numbers[k])
				}
			}
		}
	}

}
