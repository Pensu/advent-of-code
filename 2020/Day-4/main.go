package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("values.txt")

	if err != nil {
		log.Fatal("Can't open the file")
	}

	var lineparts []string
	var res int
	var count int
	params := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}
	// params := map[string]string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		matched, _ := regexp.MatchString("^\\s*$", line)
		if matched {
			var val []string
			fullVal := make(map[string]string)
			for _, item := range lineparts {
				if strings.Split(item, ":")[0] != "cid" {
					val = append(val, strings.Split(item, ":")[0])
					fullVal[strings.Split(item, ":")[0]] = strings.Split(item, ":")[1]
				}
			}
			sort.Strings(val)
			// sort.Strings(fullVal)
			fmt.Println(params, val, fullVal)
			if reflect.DeepEqual(params, val) {
				for key := range fullVal {
					if key == "byr" {
						byrVal, _ := strconv.Atoi(fullVal[key])
						if byrVal >= 1920 && byrVal <= 2002 {
							res++
						}
					} else if key == "iyr" {
						iyrVal, _ := strconv.Atoi(fullVal[key])
						if iyrVal >= 2010 && iyrVal <= 2020 {
							res++
						}
					} else if key == "eyr" {
						eyrVal, _ := strconv.Atoi(fullVal[key])
						if eyrVal >= 2020 && eyrVal <= 2030 {
							res++
						}
					} else if key == "hgt" {
						hgtVal := fullVal[key]
						matched, _ := regexp.MatchString("^(1[5-8][0-9]cm)|(19[0-3]cm)|(59in)|([6-7][0-9]in)$", hgtVal)
						if matched {
							res++
						}
					} else if key == "hcl" {
						hclVal := fullVal[key]
						matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", hclVal)
						if matched {
							res++
						}
					} else if key == "ecl" {
						eclVal := fullVal[key]
						if eclVal == "amb" || eclVal == "blu" || eclVal == "brn" || eclVal == "gry" || eclVal == "grn" || eclVal == "hzl" || eclVal == "oth" {
							res++
						}
					} else if key == "pid" {
						pidVal := fullVal[key]
						matched, _ := regexp.MatchString("^[0-9]{9}$", pidVal)
						if matched {
							res++
						}
					}
				}
			}
			fmt.Println("Value of res: ", res)
			if res == 7 {
				count++
			}
			res = 0
			lineparts = nil
		} else {
			lineparts = append(lineparts, (strings.Split(line, " "))...)
		}
	}

	if lineparts != nil {
		var val []string
		fullVal := make(map[string]string)
		for _, item := range lineparts {
			if strings.Split(item, ":")[0] != "cid" {
				val = append(val, strings.Split(item, ":")[0])
				fullVal[strings.Split(item, ":")[0]] = strings.Split(item, ":")[1]
			}
		}
		sort.Strings(val)
		fmt.Println(params, val)
		if reflect.DeepEqual(params, val) {
			for key := range fullVal {
				if key == "byr" {
					byrVal, _ := strconv.Atoi(fullVal[key])
					if byrVal >= 1920 && byrVal <= 2002 {
						res++
					}
				} else if key == "iyr" {
					iyrVal, _ := strconv.Atoi(fullVal[key])
					if iyrVal >= 2010 && iyrVal <= 2020 {
						res++
					}
				} else if key == "eyr" {
					eyrVal, _ := strconv.Atoi(fullVal[key])
					if eyrVal >= 2020 && eyrVal <= 2030 {
						res++
					}
				} else if key == "hgt" {
					hgtVal := fullVal[key]
					matched, _ := regexp.MatchString("^(1[5-8][0-9]cm)|(19[0-3]cm)|(59in)|([6-7][0-9]in)$", hgtVal)
					if matched {
						res++
					}
				} else if key == "hcl" {
					hclVal := fullVal[key]
					matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", hclVal)
					if matched {
						res++
					}
				} else if key == "ecl" {
					eclVal := fullVal[key]
					if eclVal == "amb" || eclVal == "blu" || eclVal == "brn" || eclVal == "gry" || eclVal == "grn" || eclVal == "hzl" || eclVal == "oth" {
						res++
					}
				} else if key == "pid" {
					pidVal := fullVal[key]
					matched, _ := regexp.MatchString("^[0-9]{9}$", pidVal)
					if matched {
						res++
					}
				}
			}
		}
		fmt.Println("Value of res: ", res)
		if res == 7 {
			count++
		}
	}

	fmt.Println("The result is: ", count)

}
