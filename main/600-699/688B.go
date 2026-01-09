package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf688B(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	Fprintf(out, "%s", s)
	slices.Reverse(s)
	Fprintf(out, "%s", s)
}

//func main() { cf688B(bufio.NewReader(os.Stdin), os.Stdout) }
