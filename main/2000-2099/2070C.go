package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2070C(in io.Reader, out io.Writer) {
	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := sort.Search(1e9, func(mx int) bool {
			left := k
			re := true
			for i, v := range a {
				if v <= mx {
					continue
				}
				if s[i] == 'R' {
					re = true
				} else if re {
					re = false
					left--
				}
			}
			return left >= 0
		})
		Fprintln(out, ans)
	}
}

//func main() { cf2070C(bufio.NewReader(os.Stdin), os.Stdout) }
