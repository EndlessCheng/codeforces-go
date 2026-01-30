package main

import (
	"bytes"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1673E(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	c2 := func(n, k int) byte {
		if k <= 0 {
			if n == 0 {
				return 1
			}
			return 0
		}
		if n >= k && (n-1)&(k-1) == k-1 {
			return 1
		}
		return 0
	}

	const w = 20
	ans := bytes.Repeat([]byte{'0'}, 1<<w)
	for i, v := range a {
		s := -v
		for j := i; j < n; j++ {
			s += a[j]
			if s >= w || v<<s >= len(ans) {
				break
			}
			xor := 0
			if i > 0 {
				xor = 1
			}
			if j < n-1 {
				xor++
			}
			ans[v<<s] ^= c2(n-1-(j-i)-xor, k-xor)
		}
	}

	i := len(ans) - 1
	for i > 0 && ans[i] == '0' {
		i--
	}
	ans = ans[:i+1]
	slices.Reverse(ans)
	Fprintf(out, "%s", ans)
}

//func main() { cf1673E(bufio.NewReader(os.Stdin), os.Stdout) }
