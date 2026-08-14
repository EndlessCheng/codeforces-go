package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2252B(in io.Reader, out io.Writer) {
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [2]int{}
		st := 0
		for i, b := range s {
			if i == n-1 || b != s[i+1] {
				cnt[b-'0'] += i - st
				st = i + 1
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
