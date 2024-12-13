package fixedlab4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Calculate(x float64) float64 {
	return math.Pow(math.Abs(x*x-2.5), 0.25) + math.Pow(math.Log10(x*x), 0.33333333)
}

func Task_A(begin_x, end_x, delta_x float64) []float64 {
	var result []float64

	for x := begin_x; x < end_x; x += delta_x {
		result = append(result, Calculate(x))
	}
	return result
}

func Task_B(arguments []float64) []float64 {
	var result []float64

	for _, x := range arguments {
		result = append(result, Calculate(x))
	}
	return result
}

func ReadInput(filename string) (float64, float64, float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, nil, err
	}

	if len(lines) < 3 {
		return 0, 0, 0, nil, fmt.Errorf("input file must contain at least three lines")
	}

	begin_x, err := strconv.ParseFloat(lines[0], 64)
	if err != nil {
		return 0, 0, 0, nil, fmt.Errorf("invalid begin_x: %v", err)
	}

	end_x, err := strconv.ParseFloat(lines[1], 64)
	if err != nil {
		return 0, 0, 0, nil, fmt.Errorf("invalid end_x: %v", err)
	}

	delta_x, err := strconv.ParseFloat(lines[2], 64)
	if err != nil {
		return 0, 0, 0, nil, fmt.Errorf("invalid delta_x: %v", err)
	}

	var arguments []float64
	for _, line := range lines[3:] {
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return 0, 0, 0, nil, fmt.Errorf("invalid argument in line: %v", err)
		}
		arguments = append(arguments, value)
	}

	return begin_x, end_x, delta_x, arguments, nil
}

func FixedLab4() {
	begin_x, end_x, delta_x, arguments, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println("Задача A", Task_A(begin_x, end_x, delta_x))
	fmt.Println("Задача B", Task_B(arguments))
}
