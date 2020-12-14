package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF427E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := int64(0)
	for i, j := 0, n-1; i < j; i += m {
		ans += int64(a[j] - a[i])
		j -= m
	}
	Fprint(out, ans*2)
}

//func main() { CF427E(os.Stdin, os.Stdout) }
