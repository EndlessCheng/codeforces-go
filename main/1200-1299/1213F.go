package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1213F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, K, v int
	Fscan(in, &n, &K)
	p := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		p[v] = i
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	ans := make([]byte, n)
	k := -1
	for i, mi := 0, 0; i < n; {
		st := i
		mx := p[b[i]]
		for i++; i < n && mx-mi+1 != i-st; i++ {
			if p[b[i]] > mx {
				mx = p[b[i]]
			}
		}
		if k < K-1 {
			k++
		}
		for ; st < i; st++ {
			ans[b[st]-1] = 'a' + byte(k)
		}
		mi = mx + 1
	}
	if k < K-1 {
		Fprint(out, "NO")
	} else {
		Fprintf(out, "YES\n%s", ans)
	}
}

//func main() { CF1213F(os.Stdin, os.Stdout) }
