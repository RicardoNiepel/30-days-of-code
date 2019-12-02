package day0

import (
	"bufio"
	"io"
	"os"
)

func run(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReader(stdin)
	writer := bufio.NewWriter(stdout)

	inputLine, _ := reader.ReadString('\n')

	writer.WriteString("Hello, World.\n")
	writer.WriteString(inputLine)
	writer.Flush()
}

func main() {
	run(os.Stdin, os.Stdout)
}
