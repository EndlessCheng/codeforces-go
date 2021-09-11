package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF330B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v int
	Fscan(in, &n, &m)
	vis := make([]bool, n+1)
	for m *= 2; m > 0; m-- {
		Fscan(in, &v)
		vis[v] = true
	}
	rt := 1
	for ; vis[rt]; rt++ {
	}
	Fprintln(out, n-1)
	for i := 1; i <= n; i++ {
		if i != rt {
			Fprintln(out, rt, i)
		}
	}
}

//func main() { CF330B(os.Stdin, os.Stdout) }
