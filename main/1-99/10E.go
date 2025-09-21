package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf10E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := int(1e18)
	for i, v := range a {
		for j := i + 1; j < n; j++ {
			v := v - 1
			res := a[j]
			cnt := 1
			for _, x := range a[i+1 : j+1] {
				cnt += v / x
				res += v - v%x
				v %= x
			}

			// 正常贪心
			t := res
			for _, x := range a {
				cnt -= t / x
				t %= x
			}
			if cnt < 0 {
				ans = min(ans, res)
			}
		}
	}
	if ans == 1e18 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf10E(os.Stdin, os.Stdout) }
