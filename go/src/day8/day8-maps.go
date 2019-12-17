package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func run(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReaderSize(stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	entryCount := int(nTemp)

	entries := make(map[string]int64)

	for i := 0; i < entryCount; i++ {
		arrTemp := strings.Split(readLine(reader), " ")

		nTemp, err = strconv.ParseInt(arrTemp[1], 10, 64)
		checkError(err)

		entries[arrTemp[0]] = nTemp
	}

	for query := readLine(reader); query != ""; query = readLine(reader) {
		val, prs := entries[query]
		if !prs {
			fmt.Fprintln(stdout, "Not found")
		} else {
			fmt.Fprintf(stdout, "%v=%v\n", query, val)
		}
	}
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
