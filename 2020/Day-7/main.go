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

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
		line := s.Text()

		fmt.Println(line[0])
	}
}
