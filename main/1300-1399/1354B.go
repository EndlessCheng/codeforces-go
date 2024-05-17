package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1354B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := int(1e9)
		cnt := [3]int{}
		l := 0
		for r, v := range s {
			cnt[v-'1']++
			for cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
				ans = min(ans, r-l+1)
				cnt[s[l]-'1']--
				l++
			}
		}
		if ans == 1e9 {
			ans = 0
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1354B(bufio.NewReader(os.Stdin), os.Stdout) }
