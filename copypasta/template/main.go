package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)


	Fprintln(out, )
}

func main() {
	solve(os.Stdin, os.Stdout)
}
