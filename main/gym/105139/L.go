package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func cfL(in io.Reader, out io.Writer) {
	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		if a == b {
			Fprintln(out, 0)
			continue
		}
		x := []int{a, b, lpf(a), lpf(b), 2}
		if g := gcd(a, b); g > 1 {
			x = append(x, g)
		}
		g := make([][]int, len(x))
		for i := range g {
			g[i] = make([]int, len(x))
		}
		for i, v := range x {
			for j, w := range x[:i] {
				g[i][j] = v / gcd(v, w) * w
				g[j][i] = g[i][j]
			}
		}
		for k := range g {
			for i := range g {
				for j := range g {
					g[i][j] = min(g[i][j], g[i][k]+g[k][j])
				}
			}
		}
		Fprintln(out, g[0][1])
	}
}

func main() { cfL(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func lpf(n int) int {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}
