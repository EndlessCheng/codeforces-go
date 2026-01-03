package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p10719(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	ans := int(1e9)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
		sum := make([]int, m)
		for up := i; up >= 0; up-- {
			s, l := 0, 0
			for j, b := range a[up] {
				sum[j] += int(b & 1)
				s += sum[j]
				for s >= k {
					ans = min(ans, (i-up+1)*(j-l+1))
					s -= sum[l]
					l++
				}
			}
		}
	}
	if ans == 1e9 {
		ans = 0
	}
	Fprint(out, ans)
}

//func main() { p10719(bufio.NewReader(os.Stdin), os.Stdout) }
