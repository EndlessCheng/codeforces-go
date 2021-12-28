package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1283C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	hasFrom := make([]bool, n)
	to := make([]int, n)
	for i := range to {
		Fscan(in, &to[i])
		if to[i] > 0 {
			hasFrom[to[i]-1] = true
		}
	}
	head := []int{}
	for i, f := range hasFrom {
		if !f {
			head = append(head, i)
		}
	}
	for i, v := range head {
		for to[v] > 0 {
			v = to[v] - 1
		}
		to[v] = head[(i+1)%len(head)] + 1
	}
	for _, v := range to {
		Fprint(out, v, " ")
	}
}

//func main() { CF1283C(os.Stdin, os.Stdout) }
