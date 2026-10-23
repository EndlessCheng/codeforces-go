package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2276C(in io.Reader, out io.Writer) {
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		s := make([]int, n+1)
		cnt := make([]int, k)
		ans := 0
		for r := 1; r <= n; r++ {
			Fscan(in, &v)
			s[r] = (s[r-1] + v) % k
			if r > k {
				cnt[s[r-k-1]]++
				ans += cnt[s[r]]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2276C(bufio.NewReader(os.Stdin), os.Stdout) }
