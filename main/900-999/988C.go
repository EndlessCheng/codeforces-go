package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF988C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ i, j int }

	var k, n int
	Fscan(in, &k)
	mp := map[int]pair{}
	for i := 1; i <= k; i++ {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for j := range a {
			Fscan(in, &a[j])
			s += a[j]
		}
		for j, v := range a {
			if p := mp[s-v]; p.i > 0 && p.i < i {
				Fprintln(out, "YES")
				Fprintln(out, p.i, p.j+1)
				Fprint(out, i, j+1)
				return
			}
			mp[s-v] = pair{i, j}
		}
	}
	Fprint(out, "NO")
}

//func main() { CF988C(os.Stdin, os.Stdout) }
