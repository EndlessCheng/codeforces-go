package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func cf1550E(in io.Reader, out io.Writer) {
	n, k, s := 0, 0, ""
	Fscan(bufio.NewReader(in), &n, &k, &s)
	start := make([][17]int32, n+1)
	ans := sort.Search(n/k, func(low int) bool {
		low++
		cnt := make([]int32, k)
		for i, b := range s {
			if b == '?' {
				for j := range cnt {
					cnt[j]++
				}
			} else {
				t := cnt[b-'a']
				for j := range cnt {
					cnt[j] = 0
				}
				cnt[b-'a'] = t + 1
			}
			i++
			for j, c := range cnt {
				if c >= int32(low) {
					start[i][j] = int32(i - low + 1)
				} else {
					start[i][j] = start[i-1][j]
				}
			}
		}

		u := 1 << k
		f := make([]int32, u)
		f[0] = int32(n + 1)
		for i, fv := range f {
			fv--
			if fv < 0 {
				return true
			}
			for s := uint(u - 1 ^ i); s > 0; s &= s - 1 {
				p := bits.TrailingZeros(s)
				f[i|1<<p] = max(f[i|1<<p], start[fv][p])
			}
		}
		return false
	})
	Fprint(out, ans)
}

//func main() { cf1550E(os.Stdin, os.Stdout) }
