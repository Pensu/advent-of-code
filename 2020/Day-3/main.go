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
		log.Fatal("File open failed")
	}

	s := bufio.NewScanner(f)
	s.Scan()
	rep := len(s.Text())
	case1 := 1
	case2 := 3
	case3 := 5
	case4 := 7
	var res1, res2, res3, res4 int
	for s.Scan() {
		if case1 >= rep {
			case1 = case1 - rep
		}
		if case2 >= rep {
			case2 = case2 - rep
		}
		if case3 >= rep {
			case3 = case3 - rep
		}
		if case4 >= rep {
			case4 = case4 - rep
		}

		line := s.Text()
		if line[case1] == '#' {
			res1++
		}
		if line[case2] == '#' {
			res2++
		}
		if line[case3] == '#' {
			res3++
		}
		if line[case4] == '#' {
			res4++
		}

		case1 = case1 + 1
		case2 = case2 + 3
		case3 = case3 + 5
		case4 = case4 + 7
	}
	f.Close()

	fNew, err := os.Open("values.txt")

	sNew := bufio.NewScanner(fNew)
	sNew.Scan()
	sNew.Scan()
	case5 := 1
	var res5 int
	for sNew.Scan() {
		if case5 >= rep {
			case5 = case5 - rep
		}

		line := sNew.Text()

		if line[case5] == '#' {
			res5++
		}
		sNew.Scan()
		case5 = case5 + 1
	}

	fmt.Println("The result is ", res1*res2*res3*res4*res5)
}
