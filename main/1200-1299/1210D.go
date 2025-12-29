package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1210D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, v, w, ans int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	inD := make([]int, n+1)
	outD := make([]int, n+1)
	for range m {
		Fscan(in, &v, &w)
		if v > w {
			v, w = w, v
		}
		g[v] = append(g[v], w)
		inD[w]++
		outD[v]++
	}

	for i, d := range inD {
		ans += d * outD[i]
	}
	Fprintln(out, ans)

	Fscan(in, &q)
	for range q {
		Fscan(in, &v)
		for _, w := range g[v] {
			inD[w]--
			ans += inD[w] - outD[w]
			outD[w]++
			g[w] = append(g[w], v)
		}

		ans -= inD[v] * outD[v]
		inD[v] += outD[v]
		outD[v] = 0
		g[v] = g[v][:0]
		Fprintln(out, ans)
	}
}

//func main() { cf1210D(bufio.NewReader(os.Stdin), os.Stdout) }
