package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1799F(in io.Reader, out io.Writer) {
	var T, n, b, div, sub int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &b, &div, &sub)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })

		s1 := make([]int, n+1)
		s2 := make([]int, n+1)
		s3 := make([]int, n+1)
		for i, v := range a {
			s1[i+1] = s1[i] + (v+1)/2
			s2[i+1] = s2[i] + max(v-b, 0)
			s3[i+1] = s3[i] + v
		}

		ans := int(1e18)
		s := 0
		for i := range min(div, sub) + 1 {
			leftDiv := div - i
			leftSub := sub - i
			if n-i >= leftDiv+leftSub {
				for j := i; j <= i+leftDiv; j++ {
					k := j + leftSub
					r := k + leftDiv - (j - i)
					ans = min(ans, s+s1[j]-s1[i]+s2[k]-s2[j]+s1[r]-s1[k]+s3[n]-s3[r])
				}
			}
			if i < n {
				s += max((a[i]+1)/2-b, 0)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1799F(bufio.NewReader(os.Stdin), os.Stdout) }
