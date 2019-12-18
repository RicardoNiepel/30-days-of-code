package day10

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func run(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReaderSize(stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	binary := make([]int, 0)

	for n > 0 {
		binary = append(binary, n%2)
		n = n / 2
	}

	var onesInARowCount = 0
	var maxOnesInARowCount = 0
	var currentReminder int

	for i := len(binary) - 1; i >= 0; i-- {
		currentReminder = binary[i]

		if currentReminder == 1 {
			onesInARowCount++
		} else {
			onesInARowCount = 0
		}

		maxOnesInARowCount = int(math.Max(float64(maxOnesInARowCount), float64(onesInARowCount)))
	}

	fmt.Fprint(stdout, maxOnesInARowCount)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return string(str)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
