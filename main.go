package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("TODO: HERE SOLVE CODE")
}

// SOLVE CODE

// PRESET CODE

// IOReader

type IOReader struct {
	ShouldSuspendForEmptyInput bool
	scanner                    bufio.Scanner
}

func NewIOReader() *IOReader {
	i := IOReader{}
	i.scanner = *bufio.NewScanner(os.Stdin)
	i.ShouldSuspendForEmptyInput = true
	return &i
}

func (i *IOReader) ReadInput() string {
	i.scanner.Scan()
	return i.scanner.Text()
}

func (i *IOReader) ReadInputs() []string {
	var inputs = make([]string, 0)
	for {
		input := i.ReadInput()
		if i.ShouldSuspendForEmptyInput && input == "" {
			return inputs
		}
		inputs = append(inputs, input)
	}
}

func (i *IOReader) ReadNumber() int {
	return *i.ReadNumberWithFailureSkip(false)
}

func (i *IOReader) ReadNumberWithFailureSkip(failureSkip bool) *int {
	i.scanner.Scan()
	input := i.scanner.Text()
	number, err := strconv.Atoi(input)
	if err != nil {
		if failureSkip {
			return nil
		}
		log.Fatal(err)
		notUseNumber := -1
		return &notUseNumber
	}
	return &number
}

func (i *IOReader) ReadNumbers() []int {
	var numbers = make([]int, 0)
	inputs := i.ReadInputs()
	for _, input := range inputs {
		n, err := strconv.Atoi(input)
		if err != nil {
			break
		}
		numbers = append(numbers, n)
	}
	return numbers
}

// Split Snippet

func SplitNumbers(text string) []int {
	splited := strings.Split(text, " ")
	var numbers = make([]int, 0)
	for _, s := range splited {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}
	return numbers
}

// Calc Snippet

func SumNumbers(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func PercentageRatio(percentage int) float64 {
	return float64(percentage) * 0.01
}

func MaxWithNumebers(numbers []int) int {
	max := math.MinInt
	for _, n := range numbers {
		if max < n {
			max = n
		}
	}
	return max
}

func MinWithNumebers(numbers []int) int {
	min := math.MaxInt
	for _, n := range numbers {
		if min > n {
			min = n
		}
	}
	return min
}
