package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2149E(in io.Reader, out io.Writer) {
	var T, n, k, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &l, &r)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f := func(mxL, mxK int) (res int) {
			cnt := map[int]int{}
			l := 0
			for i, v := range a {
				cnt[v]++
				for len(cnt) > mxK || i-l+1 > mxL {
					w := a[l]
					cnt[w]--
					if cnt[w] == 0 {
						delete(cnt, w)
					}
					l++
				}
				res += i - l + 1
			}
			return
		}
		Fprintln(out, f(r, k)-f(r, k-1)-f(l-1, k)+f(l-1, k-1))
	}
}

//func main() { cf2149E(bufio.NewReader(os.Stdin), os.Stdout) }
