package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1409E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
		}

		sort.Ints(a)
		pre := make([]int, n+1)
		ans, left := 0, 0
		for right, v := range a {
			for v-a[left] > k {
				left++
			}
			ans = max(ans, right-left+1+pre[left])
			pre[right+1] = max(pre[right], right-left+1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1409E(os.Stdin, os.Stdout) }
