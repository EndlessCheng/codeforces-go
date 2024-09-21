package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1016C(in io.Reader, out io.Writer) {
	var n, s int
	Fscan(in, &n)
	a := make([]int, n*2)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := n*2 - 1; i >= n; i-- {
		Fscan(in, &a[i])
	}
	sum := make([][3]int, n*2+1)
	for i, v := range a {
		sum[i+1][0] = sum[i][0] + v*i
		sum[i+1][1] = sum[i][1] + v*(n*2-1-i)
		sum[i+1][2] = sum[i][2] + v
	}

	ans := sum[n*2][0]
	for i := 0; i < n; i++ {
		s += a[i]*(i*2+i%2) + a[n*2-1-i]*(i*2+(i%2^1))
		l, r := sum[i+1], sum[n*2-1-i]
		ans = max(ans, s+r[i%2^1]-l[i%2^1]+(r[2]-l[2])*(i+1))
	}
	Fprint(out, ans)
}

//func main() { cf1016C(bufio.NewReader(os.Stdin), os.Stdout) }
