package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1348B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, 0, k)
		vis := make([]bool, n+1)
		for i := 0; i < n; i++ {
			if Fscan(in, &v); !vis[v] {
				vis[v] = true
				a = append(a, v)
			}
		}
		if len(a) > k {
			Fprintln(out, -1)
			continue
		}
		for i := len(a); i < k; i++ {
			a = append(a, 1)
		}
		Fprintln(out, len(a)*n)
		for ; n > 0; n-- {
			for _, v := range a {
				Fprint(out, v, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1348B(os.Stdin, os.Stdout) }
