package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF962B(in io.Reader, out io.Writer) {
	var a, b int
	var s string
	Fscan(bufio.NewReader(in), &a, &a, &b, &s)
	ans := a + b
	for _, s := range strings.FieldsFunc(s, func(r rune) bool { return r == '*' }) {
		n := len(s)
		if a < b {
			a, b = b, a
		}
		a -= (n + 1) / 2
		if a < 0 {
			a = 0
		}
		b -= n / 2
		if b < 0 {
			b = 0
		}
	}
	Fprint(out, ans-a-b)
}

//func main() { CF962B(os.Stdin, os.Stdout) }
