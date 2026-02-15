package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	// FS package file system for txt reading
	var myFs fs.FS = os.DirFS(".")

	args := os.Args
	args = args[1:]

	if len(args) < 1 {
		return
	}

	// Open Standart ascii pattern by default
	data, err := fs.ReadFile(myFs, "ArtStyles/sample2.txt")
	if err != nil {
		fmt.Println("ERROR: No sample of askii symbols")
		return
	}
	// If last entered argument is style, change data to specified ascii patern and delete last argument
	style := len(args) - 1
	switch args[style] {
	case "shadow":
		data, err = fs.ReadFile(myFs, "ArtStyles/sample1.txt")
		if err != nil {
			fmt.Println("ERROR: No sample of askii shadow pattern")
			return
		}
		args = args[:style]
	case "standard":
		data, err = fs.ReadFile(myFs, "ArtStyles/sample2.txt")
		if err != nil {
			fmt.Println("ERROR: No sample of askii standart pattern")
			return
		}
		args = args[:style]
	case "thinkertoy":
		data, err = fs.ReadFile(myFs, "ArtStyles/sample3.txt")
		if err != nil {
			fmt.Println("ERROR: No sample of askii thinkertoy pattern")
			return
		}
		args = args[:style]
	}

	// len of arguments
	words := len(args)

	// Divide txt sample on lines
	pattern := strings.Split(string(data), "\n")

	// Check if size of pattern files is incorrect
	if len(pattern) < 855 {
		fmt.Println("Please, Do not touch pattern files")
		return
	}

	if words > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	color := "white"
	substr := ""

	// Checing for correct format
	if len(args[0]) > 7 && args[0][:8] == "--color=" {
		color = args[0][8:]

		if words == 2 {
			AsciiArt(args[1], pattern, color, substr, true)
		} else if words == 3 {
			substr = args[1]
			word := args[2]
			if len(substr) > 7 && substr[:8] == "--color=" || len(word) > 7 && word[:8] == "--color=" {
				fmt.Println("Usage: go run . [OPTION] [STRING]")
				fmt.Println()
				fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
				return
			}
			AsciiArt(word, pattern, color, substr, false)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
			return
		}
	} else {
		if words != 1 {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
			return
		} else {
			AsciiArt(args[0], pattern, color, substr, false)
		}
	}

}
