package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1681D(in io.Reader, out io.Writer) {
	var n, ans int
	q := []int64{0}
	Fscan(in, &n, &q[0])
	for ; len(q) > 0; ans++ {
		tmp := q
		q = nil
		inq := map[int64]bool{}
		for _, v := range tmp {
			cnt := 0
			for x := v; x > 0; x /= 10 {
				cnt++
				d := x % 10
				if d > 1 && !inq[v*d] {
					inq[v*d] = true
					q = append(q, v*d)
				}
			}
			if cnt == n {
				Fprint(out, ans)
				return
			}
		}
	}
	Fprint(out, -1)
}

//func main() { CF1681D(os.Stdin, os.Stdout) }
