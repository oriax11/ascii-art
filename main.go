package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "hello\n\nhey"
	str := ""
	for i := 0; i < 8; i++ {
		for j,_:=  range s {
		
			str = str + sis(rune(s[j]), i)
		}
		str=str+"\n"
	}
	fmt.Print(str)
}

func sis(a rune, b int) string {
	filename := "standard.txt"
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	lines := strings.Split(string(filecontent), "\n")

	// Ensure that the ASCII value is within the valid range
	if int(a)-32 < 0 || int(a)-32 >= len(lines) {
		fmt.Println("Invalid ASCII value:", a)
	}

	// Calculate the index of the line based on the ASCII value and offset
	index := (int(a)-32)*9 + b + 1

	// Ensure that the index is within the valid range
	if index < 0 || index >= len(lines) {
		fmt.Println("Invalid index:", index)
		return ""
	}

	return lines[index]
}
