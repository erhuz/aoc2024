package main

import (
	"fmt"
	"io"
	"os"
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

	// Print the content of the file
	fmt.Println(string(content))
}
