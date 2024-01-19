package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1922D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+2)
		l := make([]int, n+2)
		r := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			l[i], r[i] = i-1, i+1
		}
		d := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &d[i])
		}

		dead := make([]bool, n+1)
		q := make([]int, n)
		for i := range q {
			q[i] = i + 1
		}
		for k := 0; k < n; k++ {
			del := []int{}
			for _, i := range q {
				if 0 < i && i <= n && !dead[i] && a[l[i]]+a[r[i]] > d[i] {
					dead[i] = true
					del = append(del, i)
				}
			}
			q = make([]int, 0, len(del)*2)
			for _, i := range del {
				q = append(q, l[i], r[i])
				r[l[i]] = r[i]
				l[r[i]] = l[i]
			}
			Fprint(out, len(del), " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1922D(os.Stdin, os.Stdout) }
