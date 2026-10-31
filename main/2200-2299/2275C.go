package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2275C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		cnt := [2]map[int]int{{}, {}}
		ans := 0
		for i := range a {
			Fscan(in, &a[i])
			if i < 4 {
				continue
			}
			s := a[i-4] + a[i-2] - a[i]
			ans += cnt[0][s] + cnt[1][s]
			cnt[i&1][s]++
			if i >= 6 && a[i-6]+a[i-4]-a[i-2] == s {
				ans--
			}
			if i >= 8 && a[i-8]+a[i-6]-a[i-4] == s {
				ans--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2275C(bufio.NewReader(os.Stdin), os.Stdout) }
