package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf689B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	to := make([]int, n)
	for i := range to {
		Fscan(in, &to[i])
		to[i]--
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	q := []int{0}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w := to[v]
		if dis[w] < 0 {
			dis[w] = dis[v] + 1
			q = append(q, w)
		}
		if v < n-1 {
			w := v + 1
			if dis[w] < 0 {
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
		if v > 0 {
			w := v - 1
			if dis[w] < 0 {
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
	}
	for _, v := range dis {
		Fprint(out, v, " ")
	}
}

//func main() { cf689B(bufio.NewReader(os.Stdin), os.Stdout) }
