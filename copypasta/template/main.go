package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)

}

func main() { run(os.Stdin, os.Stdout) }
