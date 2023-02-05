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
	reader := NewIOReader()
	inputs := reader.ReadInputs()
	fmt.Println(inputs)
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

// Switch Snippet

func SwapValuesWithValue[T comparable](values []T, value T) *T {
	for _, v := range values {
		if v != value {
			return &v
		}
	}
	return nil
}

// Calc Snippet

func MakeNumberRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

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

func Max[T Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func Map[T any, S any](values []T, mapping func(T) S) []S {
	var mappedValues = make([]S, 0)
	for _, v := range values {
		mappedValues = append(mappedValues, mapping(v))
	}
	return mappedValues
}

func Filter[T any](values []T, filter func(T) bool) []T {
	var filteredValues = make([]T, 0)
	for _, v := range values {
		if filter(v) {
			filteredValues = append(filteredValues, v)
		}
	}
	return filteredValues
}

func ContainsWithValue[T comparable](content []T, value T) bool {
	for _, v := range content {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsWithValues[T comparable](content []T, values []T) bool {
	for _, c := range content {
		for _, v := range values {
			if c == v {
				return true
			}
		}
	}
	return false
}

func PowByInt(value int, pow int) int {
	return int(math.Pow(float64(value), float64(pow)))
}

// Constraints

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}
