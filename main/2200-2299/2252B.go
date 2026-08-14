package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2252B(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [2]int{}
		for i := 1; i < n; i++ {
			if s[i-1] == s[i] {
				cnt[s[i]-'0']++
			}
		}

		ans := cnt[0] + cnt[1]
		d := cnt[0] - cnt[1]
		tar := byte('1')
		if d < 0 {
			d = -d
			tar = '0'
		}
		if d > 1 {
			if s[0] == tar {
				ans++
				d--
			}
			if d > 1 && s[n-1] == tar {
				ans++
				d--
			}
			if d > 1 {
				ans = -1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2252B(bufio.NewReader(os.Stdin), os.Stdout) }
