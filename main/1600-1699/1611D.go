package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1611D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, root, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pa := make([]int, n)
		for i := range pa {
			Fscan(in, &pa[i])
			if pa[i]--; pa[i] == i {
				root = i
			}
		}
		Fscan(in, &v)
		ok := v-1 == root
		dep := make([]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			if !ok {
				continue
			}
			v--
			if pa[v] != root && dep[pa[v]] == 0 {
				ok = false
			}
			dep[v] = i
		}
		if ok {
			for v, d := range dep {
				Fprint(out, d-dep[pa[v]], " ")
			}
			Fprintln(out)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1611D(os.Stdin, os.Stdout) }
