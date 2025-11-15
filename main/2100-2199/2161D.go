package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2161D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([][]int, n+1)
		for i := range n {
			Fscan(in, &v)
			pos[v] = append(pos[v], i)
		}

		f := make([]int, n+1)
		mx := 0
		for x := 1; x <= n; x++ {
			d, u := pos[x-1], pos[x]
			j := len(d) - 1
			sufMax := 0
			for i := len(u) - 1; i >= 0; i-- {
				for j >= 0 && d[j] > u[i] {
					sufMax = max(sufMax, f[d[j]])
					j--
				}
				sufMax++
				// max(选 x-1, 不选 x-1)
				f[u[i]] = max(sufMax, mx+len(u)-i)
			}
			for _, i := range d {
				mx = max(mx, f[i])
			}
		}
		Fprintln(out, n-slices.Max(f))
	}
}

//func main() { cf2161D(bufio.NewReader(os.Stdin), os.Stdout) }
