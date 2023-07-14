package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF1238E(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var n, m int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &m, &s)
	cnt := make([][]int, m)
	for i := range cnt {
		cnt[i] = make([]int, m)
	}
	cntAll := make([]int, m)
	for i := 1; i < n; i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		if x != y {
			cnt[x][y]++
			cnt[y][x]++
			cntAll[x]++
			cntAll[y]++
		}
	}

	f := make([]int, 1<<m)
	for i := 1; i < len(f); i++ {
		f[i] = 1e9
	}
	for s, dv := range f {
		i := bits.OnesCount(uint(s))
		for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			ns := s | lb
			p := bits.TrailingZeros(uint(lb))
			sum := 0
			for _s := uint(s); _s > 0; _s &= _s - 1 {
				sum += cnt[p][bits.TrailingZeros(_s)]
			}
			f[ns] = min(f[ns], dv+(sum*2-cntAll[p])*i)
		}
	}
	Fprint(out, f[len(f)-1])
}

//func main() { CF1238E(os.Stdin, os.Stdout) }
