package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1500A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	p := [5e6 + 1][2]int{}
	for i, v := range a {
		for j := i + 1; j < n; j++ {
			s := v + a[j]
			if p := p[s]; p[1] > 0 && i != p[0] && i != p[1] && j != p[0] && j != p[1] {
				Fprint(out, "YES\n", p[0]+1, p[1]+1, i+1, j+1)
				return
			}
			p[s] = [2]int{i, j}
		}
	}
	Fprint(out, "NO")
}

//func main() { CF1500A(os.Stdin, os.Stdout) }
