package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1338A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, mx, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		for Fscan(in, &n, &mx); n > 1; n-- {
			if Fscan(in, &v); v < mx {
				ans = max(ans, bits.Len(uint(mx-v)))
			} else {
				mx = v
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1338A(os.Stdin, os.Stdout) }
