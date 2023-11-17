package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func run1484(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans, mx int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] > mx {
			mx = a[i]
		}
	}
	sort.Search(mx+1, func(fee int) bool {
		var f0, f1, cnt0, cnt1 int
		for _, v := range a {
			v -= fee
			if f1 > f0+v { // 不选
				f0 = f1
				cnt0 = cnt1
			} else { // 选（相等时也选）
				f0, f1 = f1, f0+v
				cnt0, cnt1 = cnt1, cnt0+1
			}
		}
		if cnt1 >= k {
			ans = f1 + k*fee
			return false
		}
		if fee == 0 {
			ans = f1
		}
		return true
	})
	Fprint(out, ans)
}

//func main() { run1484(os.Stdin, os.Stdout) }
