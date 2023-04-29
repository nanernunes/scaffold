package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FileContains(fileName, content string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return false
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == content {
			return true
		}
	}

	return false
}

func AppendInFile(fileName, upperBound, lowerBound, newLine string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	var lines []string
	var linesReturning []int

	scanner := bufio.NewScanner(file)

	isInFunc := false

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if strings.HasPrefix(line, upperBound) {
			isInFunc = true
		} else if isInFunc && line == lowerBound {
			isInFunc = false
			linesReturning = append(linesReturning, i)
		}

		lines = append(lines, line)
	}

	if len(linesReturning) > 0 {
		last := linesReturning[len(linesReturning)-1]

		if lines[last-1] == "" {
			last -= 1
		}

		lines[last] = fmt.Sprintf("%s", newLine) + lines[last]
	}

	file.Close()

	file, err = os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}
