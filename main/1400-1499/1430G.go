package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf1430G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	w := make([][]int, n)
	for i := range w {
		w[i] = make([]int, n)
	}
	inMask := make([]int, 1<<n)
	g := make([]int, 1<<n)
	for range m {
		var u, v, k int
		Fscan(in, &u, &v, &k)
		u--
		v--
		w[u][v] += k
		s := (1<<n - 1) ^ (1 << u) ^ (1 << v)
		for j := s; ; j = (j - 1) & s {
			g[j|1<<u] += k
			inMask[j|1<<v] |= 1 << u
			inMask[j|1<<v|1<<u] |= 1 << u
			if j == 0 {
				break
			}
		}
	}

	f := make([]int, 1<<n)
	pre := make([]int, 1<<n)
	for i := range f {
		f[i] = math.MaxInt
	}
	f[0] = 0
	for i := 0; i < 1<<n; i++ {
		if f[i] == math.MaxInt {
			continue
		}
		s := (1<<n - 1) ^ i
		for j := s; j > 0; j = (j - 1) & s {
			if inMask[j]|i == i && f[i|j] > f[i]+g[i] {
				pre[i|j] = i
				f[i|j] = f[i] + g[i]
			}
		}
	}

	ans := make([]int, n)
	for now := (1 << n) - 1; now > 0; now = pre[now] {
		for i := range n {
			if now>>i&1 > 0 {
				ans[i]++
			}
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { cf1430G(bufio.NewReader(os.Stdin), os.Stdout) }
