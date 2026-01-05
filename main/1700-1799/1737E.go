package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1737E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	const mod2 = (mod + 1) / 2
	const mx int = 1e6 + 1
	powInv := [mx]int{1}
	for i := 1; i < mx; i++ {
		powInv[i] = powInv[i-1] * mod2 % mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := make([]int, n)
		f[n-1] = powInv[(n-1)/2]
		s := 0
		for i := n - 2; i >= 1; i-- {
			if i*2+1 < n {
				s += f[i*2+1]
			}
			if i*2+2 < n {
				s += f[i*2+2]
			}
			f[i] = powInv[i/2+1] * (1 - s%mod + mod) % mod
		}
		for _, v := range f {
			Fprintln(out, v)
		}
	}
}

//func main() { cf1737E(bufio.NewReader(os.Stdin), os.Stdout) }
