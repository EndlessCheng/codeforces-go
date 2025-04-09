package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2093E(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := sort.Search(n, func(low int) bool {
			low++
			k := k
			for i := 0; i < n && k > 0; i++ {
				has := map[int]bool{}
				for mex := 0; i < n; i++ {
					has[a[i]] = true
					for has[mex] {
						mex++
					}
					if mex >= low {
						k--
						break
					}
				}
			}
			return k > 0
		})
		Fprintln(out, ans)
	}
}

//func main() { cf2093E(bufio.NewReader(os.Stdin), os.Stdout) }
