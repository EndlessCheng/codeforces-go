package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1620G(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, ans int
	var s string
	Fscan(in, &n)
	cnt := make([][26]int, n)
	for i := range cnt {
		Fscan(in, &s)
		for _, b := range s {
			cnt[i][b-'a']++
		}
	}

	f := make([]int32, 1<<n)
	mn := [26]int{}
	for mask := 1; mask < 1<<n; mask++ {
		for i := range 26 {
			mn[i] = 1e9
		}
		for m := uint(mask); m > 0; m &= m - 1 {
			for i, c := range cnt[bits.TrailingZeros(m)][:] {
				mn[i] = min(mn[i], c)
			}
		}

		res := bits.OnesCount(uint(mask))%2*2 - 1 + mod
		for i := range 26 {
			res = res * (mn[i] + 1) % mod
		}
		f[mask] = int32(res)
	}

	for i := range n {
		for j := 0; j < 1<<n; j++ {
			j |= 1 << i
			f[j] = (f[j] + f[j^1<<i]) % mod
		}
	}

	for mask, v := range f {
		sum := 0
		for m := uint(mask); m > 0; m &= m - 1 {
			sum += bits.TrailingZeros(m)
		}
		k := bits.OnesCount(uint(mask))
		ans ^= int(v) * k * (sum + k)
	}
	Fprint(out, ans)
}

//func main() { cf1620G(bufio.NewReader(os.Stdin), os.Stdout) }
