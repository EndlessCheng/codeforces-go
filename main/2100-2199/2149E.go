package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2149E(in io.Reader, out io.Writer) {
	var T, n, k, L, R int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &L, &R)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f := func(mx int) (res int) {
			cnt := map[int]int{}
			l := 0
			for i, v := range a {
				cnt[v]++
				for len(cnt) > mx {
					w := a[l]
					cnt[w]--
					if cnt[w] == 0 {
						delete(cnt, w)
					}
					l++
				}
				res += max(min(i-l+1, R)-L+1, 0)
			}
			return
		}
		Fprintln(out, f(k)-f(k-1))
	}
}

//func main() { cf2149E(bufio.NewReader(os.Stdin), os.Stdout) }
