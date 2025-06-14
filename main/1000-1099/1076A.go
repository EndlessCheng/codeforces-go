package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1076A(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(in, &n, &s)
	for i := range n - 1 {
		if s[i] > s[i+1] {
			Fprint(out, s[:i], s[i+1:])
			return
		}
	}
	Fprint(out, s[:n-1])
}

//func main() { cf1076A(bufio.NewReader(os.Stdin), os.Stdout) }
