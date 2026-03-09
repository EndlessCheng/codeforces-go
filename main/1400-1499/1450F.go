package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1450F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		cnt := make([]int, n+1)
		cnt1 := make([]int, n+1)
		mx := 0
		sum1 := 2
		mx1 := 0
		for i := range a {
			Fscan(in, &a[i])
			cnt[a[i]]++
			mx = max(mx, cnt[a[i]])
			if i > 0 && a[i] == a[i-1] {
				cnt1[a[i]] += 2
				mx1 = max(mx1, cnt1[a[i]])
				sum1 += 2
			}
		}
		cnt1[a[0]]++
		cnt1[a[n-1]]++
		mx1 = max(mx1, cnt1[a[0]], cnt1[a[n-1]])
		if mx > (n+1)/2 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, sum1/2-1+max(mx1-sum1/2-1, 0))
		}
	}
}

//func main() { cf1450F(bufio.NewReader(os.Stdin), os.Stdout) }
