package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultTimer = 30

func main() {
	lines, err := ReadCsv("problems.csv")

	if err != nil {
		panic(err)
	}

	defaultTimer := time.NewTimer(time.Second * defaultTimer)
	var numberOfRightAnswers = 0

	defer defaultTimer.Stop()

	go func() {
		// Block until timer finishes. When it is done, it sends a message
		// on the channel timer.C. No other code in
		// this goroutine is executed until that happens.
		<-defaultTimer.C
		// If main() finishes before the second timer, we won't get here
		fmt.Println("Right Answers: " + strconv.Itoa(numberOfRightAnswers))
		fmt.Println("Bad Answers: " + strconv.Itoa(len(lines)-numberOfRightAnswers))
	}()

	for _, line := range lines {
		fmt.Println(line[0] + "=")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		textWithoutDelimiter := strings.TrimSuffix(text, "\n")

		if textWithoutDelimiter == line[1] {
			numberOfRightAnswers = numberOfRightAnswers + 1
		}
	}

	fmt.Println("Right Answers: " + strconv.Itoa(numberOfRightAnswers))
	fmt.Println("Bad Answers: " + strconv.Itoa(len(lines)-numberOfRightAnswers))
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
