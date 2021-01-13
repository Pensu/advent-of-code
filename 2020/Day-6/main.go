package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("values.txt")

	if err != nil {
		log.Fatal("Can't open the file")
	}

	s := bufio.NewScanner(f)
	var lineparts []string
	var count int
	var linecount int
	valMap := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0,
		"f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0,
		"n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0,
		"v": 0, "w": 0, "x": 0, "y": 0, "z": 0}

	for s.Scan() {
		line := s.Text()
		matched, _ := regexp.MatchString("^\\s*$", line)
		if matched {
			fmt.Println(valMap)
			fmt.Println(lineparts, linecount)
			for _, item := range lineparts {
				for _, ans := range item {
					valMap[string(ans)]++
				}
			}

			for key := range valMap {
				if valMap[key] == linecount {
					count++
				}
			}
			fmt.Println(valMap, count)
			for key := range valMap {
				delete(valMap, key)
			}
			lineparts = nil
			linecount = 0
		} else {
			linecount++
			lineparts = append(lineparts, line)
		}
	}

	if lineparts != nil {
		fmt.Println(lineparts, linecount)
		for _, item := range lineparts {
			for _, ans := range item {
				valMap[string(ans)]++
			}
		}

		for key := range valMap {
			if valMap[key] == linecount {
				count++
			}
		}
	}

	fmt.Println("Final answer: ", count)
}
