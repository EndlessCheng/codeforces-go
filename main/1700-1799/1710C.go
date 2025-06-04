package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1710C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var s string
	Fscan(in, &s)
	dp := make([][2][2][2][2][2][2]int, len(s))
	var f func(p int, x, y, z byte, limX, limY, limZ bool) int
	f = func(p int, x, y, z byte, limX, limY, limZ bool) (res int) {
		if p == len(s) {
			return int(x & y & z)
		}
		t := &dp[p][x][y][z][b2i10(limX)][b2i10(limY)][b2i10(limZ)]
		if *t > 0 {
			return *t - 1
		}
		var upX, upY, upZ byte = 1, 1, 1
		if limX { upX = s[p] - '0' }
		if limY { upY = s[p] - '0' }
		if limZ { upZ = s[p] - '0' }
		for i := range upX + 1 {
			for j := range upY + 1 {
				for k := range upZ + 1 {
					res += f(p+1, x|(i^j)&(j^k), y|(i^k)&(j^k), z|(i^k)&(i^j), 
						limX && i == upX, limY && j == upY, limZ && k == upZ)
				}
			}
		}
		res %= mod
		*t = res + 1
		return
	}
	Fprint(out, f(0, 0, 0, 0, true, true, true))
}

//func main() { cf1710C(bufio.NewReader(os.Stdin), os.Stdout) }
func b2i10(b bool)uint8{if b{return 1};return 0}
