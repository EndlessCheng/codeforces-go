package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1051C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, c1, v0 int
	Fscan(in, &n)
	pos := map[int][]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		pos[v] = append(pos[v], i)
	}

	for v, p := range pos {
		if len(p) == 1 {
			c1++
		} else if len(p) > 2 {
			v0 = v
		}
	}

	ans := bytes.Repeat([]byte{'A'}, n)
	if c1%2 > 0 {
		if v0 == 0 {
			Fprint(out, "NO")
			return
		}
		ans[pos[v0][0]] = 'B'
		c1++
	}
	c1 /= 2
	for _, p := range pos {
		if len(p) == 1 {
			if c1 > 0 {
				c1--
			} else {
				ans[p[0]] = 'B'
			}
		}
	}
	Fprintf(out, "YES\n%s", ans)
}

//func main() { CF1051C(os.Stdin, os.Stdout) }
