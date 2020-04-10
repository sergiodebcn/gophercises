package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := ReadCsv("problems.csv")
	if err != nil {
		panic(err)
	}

	var numberOfRightAnswers = 0

	for _, line := range lines {


		reader := bufio.NewReader(os.Stdin)
		fmt.Println(line[0] + "=")
		text, _ := reader.ReadString('\n')
		textWithoutDelimiter := strings.TrimSuffix(text, "\n")

		if textWithoutDelimiter == line[1] {
			numberOfRightAnswers = numberOfRightAnswers +1
		}
	}
	fmt.Println("Right Answers: " + strconv.Itoa(numberOfRightAnswers))
	fmt.Println("Bad Answers: " + strconv.Itoa(len(lines) - numberOfRightAnswers))
}


// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

