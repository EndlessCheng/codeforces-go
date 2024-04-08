package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1875D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v < n {
				cnt[v]++
			}
		}
		mex := 0
		for cnt[mex] > 0 {
			mex++
		}
		ans := mex * (cnt[0] - 1)
		f := make([]int, mex)
		for i := 1; i < mex; i++ {
			f[i] = 1e9
			for j, c := range cnt[:i] {
				f[i] = min(f[i], f[j]+i*c)
			}
			ans = min(ans, f[i]+mex*(cnt[i]-1))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1875D(os.Stdin, os.Stdout) }
