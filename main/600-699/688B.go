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
	t := slices.Clone(s)
	slices.Reverse(t)
	Fprintf(out, "%s%s", s, t)
}

//func main() { cf688B(bufio.NewReader(os.Stdin), os.Stdout) }
