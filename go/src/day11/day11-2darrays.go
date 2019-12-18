package day11

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func run(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReaderSize(stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	maxHourglassSum, _ := getHourglassSum(arr, 0, 0)

	for r := 0; r < 6; r++ {
		for c := 0; c < 6; c++ {
			sum, err := getHourglassSum(arr, r, c)
			if err == nil {
				maxHourglassSum = int32(math.Max(float64(sum), float64(maxHourglassSum)))
			}
		}
	}

	fmt.Fprint(stdout, maxHourglassSum)
}

func getHourglassSum(array [][]int32, startPosRow int, startPosCol int) (int32, error) {
	var sum int32 = 0
	rowCount := len(array)
	colCount := len(array[0])

	if startPosRow+2 > rowCount-1 || startPosCol+2 > colCount-1 {
		return -1, errors.New("Does not exist")
	}

	sum += array[startPosRow][startPosCol]
	sum += array[startPosRow][startPosCol+1]
	sum += array[startPosRow][startPosCol+2]
	sum += array[startPosRow+1][startPosCol+1]
	sum += array[startPosRow+2][startPosCol]
	sum += array[startPosRow+2][startPosCol+1]
	sum += array[startPosRow+2][startPosCol+2]

	return sum, nil
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
