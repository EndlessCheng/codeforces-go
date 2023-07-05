package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1157C2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := []byte{}
	l, r, pre := 0, n-1, 0
	for l <= r && (a[l] > pre || a[r] > pre) {
		isR := a[l] <= pre || a[r] > pre && a[r] < a[l]
		if a[l] == a[r] { // 至多发生一次
			ll := l + 1
			for ll < r && a[ll] > a[ll-1] {
				ll++
			}
			rr := r - 1
			for rr > l && a[rr] > a[rr+1] {
				rr--
			}
			isR = r-rr > ll-l
		}
		if isR {
			ans = append(ans, 'R')
			pre = a[r]
			r--
		} else {
			ans = append(ans, 'L')
			pre = a[l]
			l++
		}
	}
	Fprintf(out, "%d\n%s", len(ans), ans)
}

//func main() { CF1157C2(os.Stdin, os.Stdout) }
