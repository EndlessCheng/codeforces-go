package main

import (
	. "fmt"
	"io"
	"sort"
)

func cf1996F(in io.Reader, out io.Writer) {
	var T, n, k, ans int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]struct{ a, b int }, n)
		for i := range a {
			Fscan(in, &a[i].a)
		}
		for i := range a {
			Fscan(in, &a[i].b)
		}
		sort.Search(1e9+1, func(mx int) bool {
			k, s := k, 0
			for _, p := range a {
				if p.a > mx {
					t := (p.a-mx-1)/p.b + 1
					k -= t
					s += (p.a*2 - (t-1)*p.b) * t / 2
				}
			}
			if k < 0 {
				return false
			}
			// 二分结束后，剩余的 k 次操作，每次操作的得分一定都恰好等于 mx
			// 反证法：如果操作若干次后，后续的操作只能得到 < mx 的分数，那么二分结果必然 < mx，矛盾
			ans = s + mx*k
			return true
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1996F(bufio.NewReader(os.Stdin), os.Stdout) }
