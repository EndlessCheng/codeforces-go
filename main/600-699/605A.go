package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF605A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		p[v-1] = i
	}
	for i := 0; i < n; {
		st := i
		for i++; i < n && p[i] > p[i-1]; i++ {
		}
		if i-st > ans {
			ans = i - st
		}
	}
	Fprint(out, n-ans)
}

//func main() { CF605A(os.Stdin, os.Stdout) }
