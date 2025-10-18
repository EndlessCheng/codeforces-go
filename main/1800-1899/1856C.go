package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1856C(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := sort.Search(1e9, func(mx int) bool {
			mx++
			for i := range a {
				tar := mx
				k := k
				for _, v := range a[i:] {
					if v >= tar {
						return false
					}
					k -= tar - v
					if k < 0 {
						break
					}
					tar--
				}
			}
			return true
		})
		ans = max(ans, slices.Max(a))
		Fprintln(out, ans)
	}
}

//func main() { cf1856C(bufio.NewReader(os.Stdin), os.Stdout) }
