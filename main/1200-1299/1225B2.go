package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1225B2(in io.Reader, out io.Writer) {
	var T, n, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d, &d)
		a := make([]int, n)
		cnt := map[int]int{}
		ans := n
		for i := range a {
			Fscan(in, &a[i])
			cnt[a[i]]++
			l := i - d + 1
			if l < 0 {
				continue
			}
			ans = min(ans, len(cnt))
			v := a[l]
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1225B2(bufio.NewReader(os.Stdin), os.Stdout) }
