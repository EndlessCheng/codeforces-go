package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf809C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	const inv2 = (mod + 1) / 2
	var T, xl, xr, yl, yr, k int
	var f func(int, int, int) int
	f = func(x, y, add int) int {
		if x < y {
			x, y = y, x
		}
		if y == 0 || k <= add {
			return 0
		}
		hb := 1 << (bits.Len(uint(x)) - 1)
		sz := min(hb+add, k) - add
		res := (add*2 + 1 + sz) % mod
		res = res * sz % mod
		res = res * inv2 % mod
		res = res * min(y, hb) % mod
		if y > hb {
			res += f(hb, y-hb, add+hb)
			res += f(x-hb, y-hb, add)
		}
		res += f(x-hb, min(y, hb), add+hb)
		return res % mod
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &xl, &yl, &xr, &yr, &k)
		ans := f(xr, yr, 0) - f(xr, yl-1, 0) - f(xl-1, yr, 0) + f(xl-1, yl-1, 0)
		Fprintln(out, (ans%mod+mod)%mod)
	}
}

//func main() { cf809C(bufio.NewReader(os.Stdin), os.Stdout) }
