package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2247C(in io.Reader, out io.Writer) {
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		cnt := [2]int{}
		c1 := 0
		for _, v := range a {
			c1 += v
			Fscan(in, &w)
			if v != w {
				cnt[v]++
			}
		}
		if cnt[0] == 0 && cnt[1] == 0 {
			Fprintln(out, 0)
		} else if c1 == 0 && cnt[0] > 0 || cnt[1] == 0 && n-c1 == cnt[0] {
			Fprintln(out, -1)
		} else {
			Fprintln(out, 2-cnt[1]%2)
		}
	}
}

//func main() { cf2247C(bufio.NewReader(os.Stdin), os.Stdout) }
