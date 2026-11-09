package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1924E(in io.Reader, _w io.Writer) {
	const mod = 1_000_000_007
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e6
	inv := [mx + 1]int{}
	inv[1] = 1
	for i := 2; i <= mx; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}

	solve := func() int {
		var n, m, k int
		Fscan(in, &n, &m, &k)
		if n*m < k {
			return 0
		}
		k--
		ans := 1
		for i := 1; i < n; i++ {
			if k/i < m {
				ans += inv[i+k/i]
			}
		}
		for i := 1; i < m; i++ {
			if k/i < n {
				ans += inv[i+k/i]
			}
		}
		return ans % mod
	}

	var T int
	for Fscan(in, &T); T > 0; T-- {
		Fprintln(out, solve())
	}
}

//func main() { cf1924E(bufio.NewReader(os.Stdin), os.Stdout) }
