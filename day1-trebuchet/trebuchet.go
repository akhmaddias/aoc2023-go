package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func trebuchet(lines []string) {
	sum := 0
	for i:= 0; i < len(lines); i++ {
		// Find indices of first and last digits
		firstDigitIndex := strings.IndexFunc(lines[i], unicode.IsDigit)
		lastDigitIndex := strings.LastIndexFunc(lines[i], unicode.IsDigit)

		// Find first and last digits
		firstDigit := lines[i][firstDigitIndex]
		lastDigit := lines[i][lastDigitIndex]

		// Get calibration value by concatinating digits and then converting to int
		calibrationValue, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))

		// Add to sum
		sum += calibrationValue
	}
	fmt.Printf("Total sum is %d\n", sum)
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
			log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close() // Close the file when the function returns

	var lines []string // Create a slice to hold the lines

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Read each line
			lines = append(lines, scanner.Text()) // Append the line to the slice
	}

	if err := scanner.Err(); err != nil {
			log.Fatalf("error while reading file: %s", err)
	}

	trebuchet(lines)
}