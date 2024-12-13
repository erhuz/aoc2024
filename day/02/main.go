package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Specify the path to your .txt file
	filePath := "data.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire contents of the file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rowDelimiter := []byte("\n")

	rows := bytes.Split(content, rowDelimiter)
	rows = rows[:len(rows)-1]

	partOne(rows)
	partTwo(rows)
}

func partOne(rows [][]byte) {
	valueDelimiter := []byte(" ")

	safeReportsCounter := 0
	for i := 0; i < len(rows); i++ {
		values := bytes.Split(rows[i], valueDelimiter)

		if determineSafePartOne(values) {
			safeReportsCounter += 1
		}
	}
	fmt.Println("Safe reports: ", safeReportsCounter)
}

func determineSafePartOne(values [][]byte) bool {
	value1, err := strconv.Atoi(string(values[0]))
	value2, err2 := strconv.Atoi(string(values[1]))

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}

	var increasing bool = value1 < value2
	for i := 0; i < len(values)-1; i++ {
		value1, err := strconv.Atoi(string(values[i]))
		value2, err2 := strconv.Atoi(string(values[i+1]))

		if err != nil {
			panic(err)
		}
		if err2 != nil {
			panic(err2)
		}

		diff := 0
		if increasing {
			diff = value2 - value1
		} else {
			diff = value1 - value2
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func partTwo(rows [][]byte) {
	valueDelimiter := []byte(" ")

	safeReportsCounter := 0
	for i := 0; i < len(rows); i++ {
		values := bytes.Split(rows[i], valueDelimiter)
		if determineSafePartTwo(values) {
			fmt.Println(i, "Safe")
			safeReportsCounter += 1
		} else {
			fmt.Println(i, "Unsafe")
		}
	}
	fmt.Println("Safe reports: ", safeReportsCounter)
}

func determineSafePartTwo(values [][]byte) bool {
	// Test every combination each with one value removed
	for arrIndex := 0; arrIndex < len(values); arrIndex++ {
		var newValues [][]byte
		for i := 0; i < len(values); i++ {
			if arrIndex == i {
				continue
			}
			newValues = append(newValues, values[i])
		}

		if determineSafePartOne(newValues) {
			return true
		}
	}

	return false
}
