package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1285D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	q := [][]int{a}
	for i := 30; i >= 0; i-- {
		two := true
		tmp := q
		q = nil
		for _, a := range tmp {
			if a == nil {
				continue
			}
			b := [2][]int{}
			for _, v := range a {
				b[v>>i&1] = append(b[v>>i&1], v)
			}
			if b[0] != nil && b[1] != nil {
				if two {
					q = append(q, b[0], b[1])
				}
			} else {
				if two {
					q = nil
				}
				two = false
				q = append(q, b[0], b[1])
			}
		}
		if two {
			ans |= 1 << i
		}
	}
	Fprint(out, ans)
}

//func main() { CF1285D(os.Stdin, os.Stdout) }
