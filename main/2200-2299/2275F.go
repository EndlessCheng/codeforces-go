package main

import (
	. "fmt"
	"io"
	"math/rand"
)

// https://github.com/EndlessCheng
func cf2275F(in io.Reader, out io.Writer) {
	const mx = 1_000_001
	spf := [mx]int{}
	for i := 2; i < mx; i++ {
		if spf[i] == 0 {
			for j := i; j < mx; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mp := map[int]uint{}
		cnt := map[uint]int{}
		a := make([]uint, n)
		for i := range a {
			Fscan(in, &v)
			for v > 1 {
				p := spf[v]
				e := 1
				for v /= p; v%p == 0; v /= p {
					e ^= 1
				}
				if e == 0 {
					continue
				}
				if mp[p] == 0 {
					mp[p] = uint(rand.Uint64())
				}
				a[i] ^= mp[p]
			}
			cnt[a[i]]++
		}

		ans := 0
		s := uint(0)
		for _, x := range a {
			s ^= x
			ans += cnt[s]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2275F(bufio.NewReader(os.Stdin), os.Stdout) }
