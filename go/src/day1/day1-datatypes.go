package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func run(stdin io.Reader, stdout io.Writer) {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(stdin)
	writer := bufio.NewWriter(stdout)

	scanner.Scan()
	inputLine := scanner.Text()
	iInput, _ := strconv.ParseUint(inputLine, 0, 64)
	i = i + iInput

	scanner.Scan()
	inputLine = scanner.Text()
	dInput, _ := strconv.ParseFloat(inputLine, 64)
	d = d + dInput

	scanner.Scan()
	inputLine = scanner.Text()
	s = s + inputLine

	writer.WriteString(fmt.Sprint(i))
	writer.WriteString("\n")

	writer.WriteString(fmt.Sprintf("%.1f", d))
	writer.WriteString("\n")

	writer.WriteString(s)
	writer.Flush()
}

func main() {
	run(os.Stdin, os.Stdout)
}
