package day3

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func solve(number int32) string {
	if number%2 != 0 {
		return "Weird"
	}
	if (number >= 2 && number <= 5) || number > 20 {
		return "Not Weird"
	}
	return "Weird"
}

func run(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReaderSize(stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(NTemp)

	result := solve(n)

	writer := bufio.NewWriter(stdout)
	writer.WriteString(fmt.Sprint(result))
	writer.Flush()
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
