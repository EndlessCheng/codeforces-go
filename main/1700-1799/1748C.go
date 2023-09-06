package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1748C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		i := 0
		s := int64(0)
		for ; i < n && a[i] != 0; i++ {
			s += a[i]
			if s == 0 {
				ans++
			}
		}
		a = a[i:]
		for i, n := 0, len(a); i < n; {
			cnt := map[int64]int{}
			s := int64(0)
			mxC := 0
			for i++; i < n && a[i] != 0; i++ {
				s += a[i]
				cnt[s]++
				mxC = max(mxC, cnt[s])
			}
			ans += max(mxC, cnt[0]+1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1748C(os.Stdin, os.Stdout) }
