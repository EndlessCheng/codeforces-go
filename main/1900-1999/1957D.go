package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1957D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		suf := [30][2]int{}
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s ^= a[i]
			for j := range suf {
				suf[j][s>>j&1]++
			}
		}

		ans := 0
		pre := [30][2]int{}
		s = 0
		for _, v := range a {
			for j := range pre {
				pre[j][s>>j&1]++
			}
			hb := bits.Len(uint(v)) - 1
			ans += pre[hb][0]*suf[hb][0] + pre[hb][1]*suf[hb][1]
			s ^= v
			for j := range suf {
				suf[j][s>>j&1]--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1957D(bufio.NewReader(os.Stdin), os.Stdout) }
