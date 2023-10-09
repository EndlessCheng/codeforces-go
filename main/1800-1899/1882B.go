package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF1882B(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	in := bufio.NewReader(_r)
	var T, n, k, v, lb int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		all := 0
		a := make([]int, n)
		for i := range a {
			for Fscan(in, &k); k > 0; k-- {
				Fscan(in, &v)
				a[i] |= 1 << (v - 1)
			}
			all |= a[i]
		}
		ans := 0
		for ; all > 0; all ^= lb {
			lb = all & -all
			m := 0
			for _, v := range a {
				if v&lb == 0 {
					m |= v
				}
			}
			ans = max(ans, bits.OnesCount(uint(m)))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1882B(os.Stdin, os.Stdout) }
