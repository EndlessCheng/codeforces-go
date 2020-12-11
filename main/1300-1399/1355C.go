package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1355C(in io.Reader, out io.Writer) {
	var a, b, c, d, ans int64
	Fscan(in, &a, &b, &c, &d)
	const mx int = 1e6 + 5
	sum := [mx]int64{}
	for i := a; i <= b; i++ {
		sum[i+b]++
		sum[i+c+1]--
	}
	for i := 1; i < mx; i++ {
		sum[i] += sum[i-1]
	}
	for i := mx - 1; i > 0; i-- {
		sum[i-1] += sum[i] // 最终得到了 sum[i] 表示 x+y>=i 的数量
	}
	for i := c; i <= d; i++ {
		ans += sum[i+1] // x+y>i 的数量
	}
	Fprint(out, ans)
}

//func main() { CF1355C(os.Stdin, os.Stdout) }
