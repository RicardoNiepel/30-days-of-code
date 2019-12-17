package day5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func run(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReaderSize(stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	for i := int32(1); i <= 10; i++ {
		fmt.Fprintf(stdout, "%v x %v = %v\n", n, i, n*i)
	}
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
