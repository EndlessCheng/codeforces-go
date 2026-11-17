package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func cf731E(in io.Reader, out io.Writer) {
	var n, s int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}

	f := s
	for i := n - 1; i > 1; i-- {
		s -= a[i]
		f = max(f, s-f)
	}
	Fprint(out, f)
}

//func main() { cf731E(bufio.NewReader(os.Stdin), os.Stdout) }
