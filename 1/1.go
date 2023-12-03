package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	fileName := "example.txt"
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		text := fileScanner.Text()
		var numbers string

		for _, r := range text {
			if unicode.IsDigit(r) {
				fmt.Printf("Found rune: %s\n", string(r))
				numbers += string(r)
			}
		}

		fmt.Printf("Numbers are: %s\n", numbers)

		if len(numbers) >= 2 {
			numbers = string(numbers[0]) + string(numbers[len(numbers)-1])
		}

		fmt.Printf("Numbers corrected are: %s\n", numbers)

		var num int
		num, err = strconv.Atoi(numbers)

		if err == nil {
			fmt.Printf("Number is %d\n", num)
			sum += num
		} else {
			fmt.Println(err)
		}
	}

	fmt.Printf("Sum for file %s is %d\n", fileName, sum)

	readFile.Close()
}
