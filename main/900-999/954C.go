package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF954C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n)
	a := make([]int, n)
	Fscan(in, &a[0])
	for i := 1; i < n; i++ {
		Fscan(in, &a[i])
		d := a[i] - a[i-1]
		if d < 0 {
			d = -d
		}
		if d == 0 {
			Fprint(out, "NO")
			return
		}
		if d > 1 {
			if m == 0 {
				m = d
			} else if d != m {
				Fprint(out, "NO")
				return
			}
		}
	}
	if m == 0 {
		Fprintln(out, "YES")
		Fprint(out, int(1e9), 1)
		return
	}
	for i := 1; i < n; i++ {
		if a[i] == a[i-1]+1 && a[i-1]%m == 0 || a[i] == a[i-1]-1 && a[i]%m == 0 {
			Fprint(out, "NO")
			return
		}
	}
	Fprintln(out, "YES")
	Fprint(out, int(1e9), m)
}

//func main() { CF954C(os.Stdin, os.Stdout) }
