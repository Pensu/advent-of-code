package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("values.txt")

	if err != nil {
		log.Fatal("Can't open the file")
	}

	var lineparts []string
	var res int
	params := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		matched, _ := regexp.MatchString("^\\s*$", line)
		if matched {
			var val []string
			for _, item := range lineparts {
				if strings.Split(item, ":")[0] != "cid" {
					val = append(val, strings.Split(item, ":")[0])
				}
			}
			sort.Strings(val)
			fmt.Println(params, val)
			if reflect.DeepEqual(params, val) {
				res++
			}
			lineparts = nil
		} else {
			lineparts = append(lineparts, (strings.Split(line, " "))...)
		}
	}

	if lineparts != nil {
		var val []string
		for _, item := range lineparts {
			if strings.Split(item, ":")[0] != "cid" {
				val = append(val, strings.Split(item, ":")[0])
			}
		}
		sort.Strings(val)
		fmt.Println(params, val)
		if reflect.DeepEqual(params, val) {
			res++
		}
	}

	fmt.Println("The result is: ", res)

}
