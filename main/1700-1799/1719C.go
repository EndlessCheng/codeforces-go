package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1719C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, v, i, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		nxt := make([]int, n+1)
		mx, mxI := 0, 0
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if v > mx {
				nxt[mxI] = i
				mx, mxI = v, i
			}
		}
		nxt[mxI] = 2e9

		for range q {
			Fscan(in, &i, &k)
			ans := min(nxt[i], k+2) - i
			if i == 1 {
				ans--
			}
			Fprintln(out, max(ans, 0))
		}
	}
}

//func main() { cf1719C(bufio.NewReader(os.Stdin), os.Stdout) }
