package main

import (
	. "fmt"
	"io"
	"math/rand/v2"
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

// 确定性做法
func cf2275F_2(in io.Reader, out io.Writer) {
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

	pos := [mx]int{}
	for i := range pos {
		pos[i] = -1
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][]int, n)
		cnt := map[int]int{}
		for i := range a {
			Fscan(in, &v)
			core := 1
			for v > 1 {
				p := spf[v]
				for v%(p*p) == 0 {
					v /= p * p
				}
				if v%p == 0 {
					v /= p
					a[i] = append(a[i], p)
					core *= p
				}
			}
			cnt[core]++
		}

		coreP := []int{}
		ans := 0
	o:
		for _, ps := range a {
			for _, p := range ps {
				i := pos[p]
				if i >= 0 {
					j := len(coreP) - 1
					pos[coreP[j]] = i
					coreP[i] = coreP[j]
					coreP = coreP[:j]
					pos[p] = -1
				} else {
					pos[p] = len(coreP)
					coreP = append(coreP, p)
				}
			}
			if len(coreP) < 8 {
				mul := 1
				for _, p := range coreP {
					mul *= p
					if mul >= mx {
						continue o
					}
				}
				ans += cnt[mul]
			}
		}
		Fprintln(out, ans)

		for _, ps := range a {
			for _, p := range ps {
				pos[p] = -1
			}
		}
	}
}

//func main() { cf2275F(bufio.NewReader(os.Stdin), os.Stdout) }
