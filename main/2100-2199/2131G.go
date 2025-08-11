package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2131G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)

		cnt := map[int]int{}
		var del func(int)
		del = func(v int) {
			k--
			cnt[v]++
			if v < 31 && k >= 1<<(v-1)-1 {
				k -= 1<<(v-1) - 1
				for i := 1; i < v; i++ {
					cnt[i] += 1 << (v - 1 - i)
				}
				return
			}
			for i := 1; i < v && k > 0; i++ {
				del(i)
			}
		}
		for _, v := range a {
			if v > 30 || 1<<(v-1) >= k {
				del(v)
				break
			}
			k -= 1 << (v - 1)
			cnt[v]++
			for i := 1; i < v; i++ {
				cnt[i] += 1 << (v - 1 - i)
			}
		}
		ans := 1
		for v, c := range cnt {
			ans = ans * pow(v, c) % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2131G(bufio.NewReader(os.Stdin), os.Stdout) }
