package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1208B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := sort.Search(n, func(m int) bool {
		cnt := map[int]int{}
		for _, v := range a[m:] {
			cnt[v]++
		}
		if len(cnt) == n-m {
			return true
		}
		for i := m; i < n; i++ {
			v := a[i]
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
			cnt[a[i-m]]++
			if len(cnt) == n-m {
				return true
			}
		}
		return false
	})
	Fprint(out, ans)
}

//func main() { cf1208B(bufio.NewReader(os.Stdin), os.Stdout) }
