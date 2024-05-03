package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2  {
        fmt.Println("Please provide ")
        return
    }
    s := os.Args[1]
    var builder strings.Builder
	// if s == "" {
	// 	continue
	// }

    splitted := strings.Split(s, "\\n")
	// fmt.Println(len(splitted))
    for j, line := range splitted {
		if line==""&&j!=0{
			builder.WriteString("\n")

		}else if line==""&&j==0{
			j=j+1
		}else {
        for i := 0; i < 8; i++ {
            for _, char := range line {
                builder.WriteString(sis(char, i))
            }
            builder.WriteString("\n")
        } }

	

		

    }

    fmt.Print(builder.String())
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
	if int(a)-32 < 0 || int(a)-32 > 127 {
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
