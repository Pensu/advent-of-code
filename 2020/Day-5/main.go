package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("values.txt")

	if err != nil {
		log.Fatal("Can't open the file")
	}

	var seatID int

	seatMap := make([][]int, 128)
	for i := range seatMap {
		seatMap[i] = make([]int, 8)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		rowMin := 0
		rowMax := 127
		seatMin := 0
		seatMax := 7
		var mid int

		for i := 0; i <= 9; i++ {
			if i <= 6 {
				mid = (rowMax - rowMin) / 2
				if line[i] == 'F' {
					rowMax = rowMax - mid - 1
				} else if line[i] == 'B' {
					rowMin = rowMin + mid + 1
				}
			} else {
				mid = (seatMax - seatMin) / 2
				if line[i] == 'L' {
					seatMax = seatMax - mid - 1
				} else if line[i] == 'R' {
					seatMin = seatMin + mid + 1
				}
			}
		}

		seatMap[rowMin][seatMin] = 1
		if seatID < (rowMin*8 + seatMin) {
			seatID = rowMin*8 + seatMin
		}

	}

	for i := range seatMap {
		fmt.Println(i, seatMap[i])
	}
	fmt.Println("The highest seatID", seatID)
}
