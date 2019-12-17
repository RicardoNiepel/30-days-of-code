package day7

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
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	for k := len(arr) - 1; k >= 0; k-- {
		fmt.Fprintf(stdout, "%v", arr[k])
		if k > 0 {
			fmt.Fprint(stdout, " ")
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
