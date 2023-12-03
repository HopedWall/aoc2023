package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func GetSumOfFile(file *os.File) (sum int) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
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
		} else if len(numbers) == 1 {
			numbers = string(numbers[0]) + string(numbers[0])
		}

		fmt.Printf("Numbers corrected are: %s\n", numbers)

		num := 0
		var err error
		num, err = strconv.Atoi(numbers)

		if err == nil {
			fmt.Printf("Number is %d\n", num)
			sum += num
		} else {
			fmt.Println(err)
		}
	}

	return
}

func main() {

	// recover file from env variables
	present := false
	fileName, present := os.LookupEnv("FILE")
	if !present {
		fmt.Print("FILE env variable not found, setting to example.txt")
		fileName = "files/example.txt"
	}

	// open the file
	readFile, err := os.Open(fileName)

	if err == nil {
		// compute the sum
		sum := GetSumOfFile(readFile)
		fmt.Printf("Sum for file %s is %d\n", fileName, sum)
		readFile.Close()
	} else {
		// error, do nothing
		fmt.Println(err)
	}

}
