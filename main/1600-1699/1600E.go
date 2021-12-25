package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1600E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	p, q := 1, n-1
	for ; p < n && a[p-1] < a[p]; p++ {}
	for ; q > 0 && a[q] < a[q-1]; q-- {}
	if p&1 > 0 || (n-q)&1 > 0 {
		Fprint(out, "Alice")
	} else {
		Fprint(out, "Bob")
	}
}

//func main() { CF1600E(os.Stdin, os.Stdout) }
