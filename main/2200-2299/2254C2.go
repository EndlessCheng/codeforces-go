package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2254C2(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		cnt := [2]int{}
		ans := 0
		for i := range n {
			p := i & 1
			cnt[p] += int(s[i] - '0')
			cnt[p] -= int(t[i] - '0')
			ans += abs(cnt[p])
		}
		if cnt[0] != 0 || cnt[1] != 0 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2254C2(bufio.NewReader(os.Stdin), os.Stdout) }
