package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const inf int = 1e9
	var n int
	Fscan(in, &n)
	a := make([]struct{ x, y, z int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].z)
	}
	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[1][0] = 0
	for s, dr := range f {
		for _s := uint(s); _s > 0; _s &= _s - 1 {
			i := bits.TrailingZeros(_s)
			for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
				lb = cus & -cus
				ns := s | lb
				j := bits.TrailingZeros(uint(lb))
				f[ns][j] = min(f[ns][j], dr[i]+abs(a[j].x-a[i].x)+abs(a[j].y-a[i].y)+max(a[j].z-a[i].z, 0))
			}
		}
	}
	ans := inf
	for i, v := range f[len(f)-1] {
		ans = min(ans, v + abs(a[0].x-a[i].x)+abs(a[0].y-a[i].y)+max(a[0].z-a[i].z, 0))
	}
	Fprint(out, ans) 
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
func abs(x int) int { if x < 0 { return -x }; return x }
