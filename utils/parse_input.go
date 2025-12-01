package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

const defaultInputPath = "input.txt"

func ReadInput(inputPath string, isAbsolute bool) []string {
	var input []string

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if len(inputPath) == 0 {
		inputPath = defaultInputPath
	}

	path := filepath.Join(wd, inputPath)

	if isAbsolute {
		path = inputPath
	}

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input

}
