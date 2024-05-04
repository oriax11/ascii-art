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
	filecontent, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(0)
	}
	str := string(filecontent)
	str = strings.ReplaceAll(str, "\r", "")
	lines := strings.Split(str, "\n")

	splitted := strings.Split(s, "\\n")
	// fmt.Println(len(splitted))
	operation(splitted, lines)
}

func FindLine(lines []string, char rune, pos int) string {
	// Calculate the index of the line based on the ASCII value and offset
	index := (int(char)-32)*9 + pos + 1
	return lines[index]
}

func operation(splitted, lines []string) {
	for j, line := range splitted {
		if line == "" && j == 0 {
			j = j + 1
		} else if line == "" && j != 0 {
			fmt.Println() // newline
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range line {
					// Ensure that the ASCII value is within the valid range
					if int(char)-32 < 0 || int(char)-32 > 127 {
						fmt.Print("Invalid ASCII value:", char)
						j = j + 1
					} else {
						fmt.Print(FindLine(lines, char, i))
					}
				}
				fmt.Println() // newline
			}
		}
	}
}
