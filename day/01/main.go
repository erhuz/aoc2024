package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
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
	valueDelimiter := []byte("   ")
	leftList := []int{}
	rightList := []int{}

	for index := 0; index < len(rows); index++ {
		values := bytes.Split(rows[index], valueDelimiter)

		value1, err := strconv.Atoi(string(values[0]))
		value2, err := strconv.Atoi(string(values[1]))
		if err != nil {
			fmt.Println("Couldn't convert int")
		}

		leftList = append(leftList, value1)
		rightList = append(rightList, value2)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	diff := 0
	for index := 0; index < len(rows); index++ {
		if leftList[index] > rightList[index] {
			diff += leftList[index] - rightList[index]
		} else {
			diff += rightList[index] - leftList[index]
		}
	}
	fmt.Println(diff)
}

func partTwo(rows [][]byte) {
	valueDelimiter := []byte("   ")
	leftList := []int{}
	rightListValueByCount := make(map[int]int)

	for index := 0; index < len(rows); index++ {
		values := bytes.Split(rows[index], valueDelimiter)

		value1, err := strconv.Atoi(string(values[0]))
		value2, err := strconv.Atoi(string(values[1]))
		if err != nil {
			fmt.Println("Couldn't convert int")
		}

		_, ok := rightListValueByCount[value2]
		if !ok {
			rightListValueByCount[value2] = 0
		}

		rightListValueByCount[value2] += 1

		leftList = append(leftList, value1)
	}

	similarityScore := 0
	for index := 0; index < len(leftList); index++ {

		val, ok := rightListValueByCount[leftList[index]]
		if ok {
			similarityScore += leftList[index] * val
		}
	}
	fmt.Println(similarityScore)
}
