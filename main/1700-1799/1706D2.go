package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1706D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	ge := [1e5 + 1]int{}
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		up := -1
		for ; n > 0; n-- {
			Fscan(in, &v)
			for i := up + 1; i <= v; i++ {
				ge[i] = v
			}
			up = v
		}
		if up <= k {
			Fprintln(out, 0)
			continue
		}
		ans := int(1e9)
	o:
		for mx := up / k; mx <= up; mx++ {
			mi := mx
			for j := 1; (j-1)*(mx+1) <= up; j++ {
				v := ge[(j-1)*(mx+1)]
				if v >= j*(mx+1) {
					j = v/(mx+1) + 1 // 快速 jump 到下一个 j
				}
				if v/j < mi {
					mi = v / j
					if mx-mi >= ans { // 剪枝
						continue o
					}
				}
			}
			ans = mx - mi // 有了上面的剪枝，这里直接赋值
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1706D2(os.Stdin, os.Stdout) }
