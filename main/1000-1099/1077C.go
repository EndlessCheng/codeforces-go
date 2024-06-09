package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1077C(in io.Reader, out io.Writer) {
	var n, sum, mx, mx2, mxI int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		sum += v
		if v > mx {
			mx2 = mx
			mx = v
			mxI = i
		} else if v > mx2 {
			mx2 = v
		}
	}

	ans := []any{}
	for i, v := range a {
		if i == mxI && sum-v == mx2*2 || i != mxI && sum-v == mx*2 {
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { cf1077C(bufio.NewReader(os.Stdin), os.Stdout) }
