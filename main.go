package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Необходимо 2 аргумента: имя входного файла и имя файла для вывода результатов.")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	content, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	lines := regexp.MustCompile("\r?\n").Split(string(content), -1)

	var results []string
	exprPattern := regexp.MustCompile(`^(\d+)\s*([\+\-\*\/])\s*(\d+)\s*=\?$`)

	for _, line := range lines {
		matches := exprPattern.FindStringSubmatch(line)
		if len(matches) == 4 {
			operand1, _ := strconv.Atoi(matches[1])
			operator := matches[2]
			operand2, _ := strconv.Atoi(matches[3])

			var result int
			switch operator {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 != 0 {
					result = operand1 / operand2
				} else {
					continue
				}
			default:
				continue
			}
			results = append(results, fmt.Sprintf("%s%s%s=%d", matches[1], operator, matches[3], result))
		}
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Ошибка создания файла вывода:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, result := range results {
		writer.WriteString(result + "\n")
	}
	writer.Flush()
}
