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
			has := map[int]bool{}
			mex := 0
			for _, v := range a {
				has[v] = true
				for has[mex] {
					mex++
				}
				if mex >= low {
					k--
					has = map[int]bool{}
					mex = 0
				}
			}
			return k > 0
		})
		Fprintln(out, ans)
	}
}

//func main() { cf2093E(bufio.NewReader(os.Stdin), os.Stdout) }
