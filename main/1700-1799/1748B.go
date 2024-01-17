package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1748B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := 0
		for i := range s {
			cnt := [10]int{}
			k, mxC := 0, 0
			for _, v := range s[i:] {
				v -= '0'
				if cnt[v] == 10 {
					break
				}
				if cnt[v] == 0 {
					k++
				}
				cnt[v]++
				mxC = max(mxC, cnt[v])
				if mxC <= k {
					ans++
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1748B(os.Stdin, os.Stdout) }
