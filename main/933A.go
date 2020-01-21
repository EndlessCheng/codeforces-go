package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF933A(_r io.Reader, _w io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	sum := make([][2]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		sum[i+1][0] = sum[i][0]
		sum[i+1][1] = sum[i][1]
		sum[i+1][a[i]-1]++
	}
	ans := 0
	for i := range a {
		maxL := 0
		for _, s := range sum[:i+1] {
			maxL = max(maxL, s[0]+sum[i][1]-s[1])
		}
		maxR := 0
		for _, s := range sum[i:] {
			maxR = max(maxR, s[0]-sum[i][0]+sum[n][1]-s[1])
		}
		ans = max(ans, maxL+maxR)
	}
	Fprint(out, ans)
}

//func main() {
//	CF933A(os.Stdin, os.Stdout)
//}
