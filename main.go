package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run . <your string> ")
		return
	}
	s := os.Args[1]

	splitted := strings.Split(s, "\\n")
	// fmt.Println(len(splitted))
	for j, line := range splitted {
		if line == "" && j != 0 {
			fmt.Println()
		} else if line == "" && j == 0 {
			j = j + 1
		} else {
			for i := 0; i < 8; i++ {
				for p, char := range line {
				// Ensure that the ASCII value is within the valid range
					if int(char)-32 < 0 || int(char)-32 > 127 {
						fmt.Print("Invalid ASCII value:", line[p])
						j=j+1

					}
					fmt.Print(ReplaceString(char, i))
				}
				fmt.Println()
			}
		}
	}
}

func ReplaceString(a rune, b int) string {
	lines := readLinesFromFile("standard.txt")
	// Calculate the index of the line based on the ASCII value and offset
	index := (int(a)-32)*9 + b + 1

	// Ensure that the index is within the valid range
	if index < 0 || index >= len(lines) {
		return ""
	}
	
	return lines[index]
}

func readLinesFromFile(filename string) []string {
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(0)
	}
	lines := strings.Split(string(filecontent), "\n")
	return lines
}
