package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF895C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}
	var n, v int
	Fscan(in, &n)
	cnt := [71]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	f := make([]int, 1<<len(primes))
	f[0] = 1
	for x, c := range cnt {
		if c == 0 {
			continue
		}
		mask := 0
		for i, p := range primes {
			for ; x%p == 0; x /= p {
				mask ^= 1 << i
			}
		}
		nf := make([]int, len(f))
		for s, fs := range f {
			nf[s] = (nf[s] + fs*pow2[c-1]) % mod           // 选偶数个 x
			nf[s^mask] = (nf[s^mask] + fs*pow2[c-1]) % mod // 选奇数个 x
		}
		f = nf
	}
	Fprint(out, (f[0]-1+mod)%mod)
}

//func main() { CF895C(os.Stdin, os.Stdout) }
