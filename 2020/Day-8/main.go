package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("values.txt")

	if err != nil {
		log.Fatal("Can't open the file")
	}
	defer f.Close()

	valueMap := make(map[int]string)
	countMap := make(map[int]int)

	s := bufio.NewScanner(f)
	i := 1

	for s.Scan() {
		fmt.Println(s.Text())
		valueMap[i] = s.Text()
		countMap[i] = 0
		i++
	}

	fmt.Println(valueMap)
	fmt.Println(countMap)
	order := make([]int, len(valueMap))
	i = 0
	for k := range valueMap {
		order[i] = k
		i++
	}

	sort.Ints(order)
	fmt.Println(order)
	accCount := 0

	i = 1

	for {
		fmt.Println(i, valueMap[i])
		ins := valueMap[i]
		if countMap[i] == 1 {
			break
		} else {
			countMap[i] = 1
		}

		insSep := strings.Split(ins, " ")
		fmt.Println(insSep)

		if insSep[0] == "acc" {
			i++
			accVal, _ := strconv.Atoi(insSep[1])
			accCount = accCount + accVal
		} else if insSep[0] == "jmp" {
			jump, _ := strconv.Atoi(insSep[1])
			fmt.Println("Jump ", jump)
			i = i + jump
		} else if insSep[0] == "nop" {
			i++
		}

	}

	fmt.Println(countMap)
	fmt.Println(accCount)

}
