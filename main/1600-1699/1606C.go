package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1606C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		k++
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := int64(0)
		for i, v := range a[:n-1] {
			w := int64(math.Pow10(a[i+1] - v))
			if w > k {
				ans += k * int64(math.Pow10(v))
				k = 0
				break
			}
			w--
			ans += w * int64(math.Pow10(v))
			k -= w
		}
		if k > 0 {
			ans += k * int64(math.Pow10(a[n-1]))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1606C(os.Stdin, os.Stdout) }
