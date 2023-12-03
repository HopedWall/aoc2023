package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func GetSumOfFile(file *os.File) (sum int) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() { // for line in lines...
		text := fileScanner.Text()
		var numbers string

		// map used to check for substrings like "one", "two", ...
		// and to convert said substring to digit
		numbersMap := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}

		// find all numbers in line
		for i, r := range text {
			if unicode.IsDigit(r) {
				// if rune is digit, add it to numbers
				fmt.Printf("Found rune: %s\n", string(r))
				numbers += string(r)
			} else {
				// check for substring like "one", "two", "three"...
				substring := text[i:]
				for key, value := range numbersMap {
					if strings.HasPrefix(substring, key) {
						fmt.Printf("Found substring: %s\n", key)
						numbers += strconv.Itoa(value)
					}
				}
			}
		}

		fmt.Printf("Numbers are: %s\n", numbers)

		// only keep relevant numbers
		if len(numbers) >= 2 {
			// keep only first and last number
			numbers = string(numbers[0]) + string(numbers[len(numbers)-1])
		} else if len(numbers) == 1 {
			// repeat single number twice (i.e. 7 -> 77)
			numbers = string(numbers[0]) + string(numbers[0])
		}

		fmt.Printf("Numbers corrected are: %s\n", numbers)

		num := 0
		var err error
		num, err = strconv.Atoi(numbers)

		if err == nil {
			fmt.Printf("Number is %d\n", num)
			// add num from this line to overall sum
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
