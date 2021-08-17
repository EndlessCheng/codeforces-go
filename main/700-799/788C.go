package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF788C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e9
	const offset = 1000
	var n, k, v int
	Fscan(in, &n, &k)
	dis := make([]int, offset*2+1)
	for i := range dis {
		dis[i] = inf
	}
	for ; k > 0; k-- {
		Fscan(in, &v)
		dis[v-n+offset] = 1
	}

	q := []int{}
	for i, d := range dis {
		if d == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for i, d := range dis {
			// 是否有这种可乐
			if d > 1 {
				continue
			}
			w := v + i - offset
			if 0 <= w && w < len(dis) && dis[w] == inf {
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
	}
	if dis[offset] < inf {
		Fprint(out, dis[offset])
	} else {
		Fprint(out, -1)
	}
}

//func main() { CF788C(os.Stdin, os.Stdout) }
