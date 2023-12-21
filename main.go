package main

import (
    "bufio"
	"fmt"
	"os"
)
import "strconv"

type LineReader = func(string)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("You must pass the dayNumber argument")
		return
	}

	dayNumber, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("You need to pass a valid day number")
		return
	}

    fmt.Printf("Running day %d\n", dayNumber)
    
	file, err := os.Open(fmt.Sprintf("data/day%d.txt", dayNumber))
	if err != nil {
		fmt.Println("Couldn't open input file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch dayNumber {
	case 1:
		RunDay1(scanner)
	case 2:
		RunDay2(scanner)
	case 3:
		RunDay3(scanner)
	case 4:
		RunDay4(scanner)
	case 5:
		RunDay5(scanner)
	case 6:
		RunDay6(scanner)
	case 7:
		RunDay7(scanner)
	default:
		fmt.Println("This day is not yet implemented")
	}
}
