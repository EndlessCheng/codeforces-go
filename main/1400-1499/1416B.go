package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1416B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		s := 0
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s += a[i]
		}
		if s%n > 0 {
			Fprintln(out, -1)
			continue
		}
		Fprintln(out, 3*n-3)
		for i := 2; i <= n; i++ {
			Fprintln(out, 1, i, (i-a[i]%i)%i)
			Fprintln(out, i, 1, (a[i]-1)/i+1)
		}
		for i := 2; i <= n; i++ {
			Fprintln(out, 1, i, s/n)
		}
	}
}

//func main() { CF1416B(os.Stdin, os.Stdout) }
