package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1796D(in io.Reader, out io.Writer) {
	var T, n, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x)
		if x < 0 {
			x = -x
			k = n - k
		}

		var ans, pre, pre2, minS int
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			pre += a[i] - x
			if i >= k {
				ans = max(ans, pre-minS+k*x*2)
				pre2 += a[i-k] - x
				minS = min(minS, pre2)
			}
		}

		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v + x
		}
		q := []int{0}
		for i := 1; i <= n; i++ {
			if q[0] < i-k {
				q = q[1:]
			}
			for len(q) > 0 && sum[i] <= sum[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
			ans = max(ans, sum[i]-sum[q[0]])
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1796D(bufio.NewReader(os.Stdin), os.Stdout) }
