package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2025D(in io.Reader, out io.Writer) {
	var n, m, r, c0 int
	Fscan(in, &n, &m)
	f := make([]int, m+1)
	cnt := make([][2]int, m+1)
	for i := range n {
		Fscan(in, &r)
		if r < 0 {
			cnt[-r][0]++
		} else if r > 0 {
			cnt[r][1]++
		}
		if r != 0 && i < n-1 {
			continue
		}
		for j := 2; j <= c0; j++ {
			cnt[j][0] += cnt[j-1][0]
			cnt[j][1] += cnt[j-1][1]
		}
		for j := c0; j > 0; j-- {
			f[j] = max(f[j], f[j-1]) + cnt[j][0] + cnt[c0-j][1]
		}
		f[0] += cnt[c0][1]
		c0++
		clear(cnt)
	}
	Fprint(out, slices.Max(f))
}

//func main() { cf2025D(bufio.NewReader(os.Stdin), os.Stdout) }
