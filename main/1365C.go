package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1365C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	cnt := make([]int, n)
	pos := make([]int, n+1)
	for i := range cnt {
		Fscan(in, &v)
		pos[v] = i
	}
	for i := range cnt {
		Fscan(in, &v)
		p := pos[v]
		if p < i {
			p += n
		}
		cnt[p-i]++
	}
	for _, c := range cnt {
		if c > ans {
			ans = c
		}
	}
	Fprint(out, ans)
}

//func main() { CF1365C(os.Stdin, os.Stdout) }
