package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1203E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx = 150000
	cnt := [mx + 2]int{}
	var n, v, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}
	for i := 1; i <= mx; i++ {
		if cnt[i] == 0 {
			continue
		}
		if i > 1 && cnt[i-1] == 0 {
			cnt[i-1]++
			cnt[i]--
		}
		if cnt[i] > 1 {
			cnt[i+1]++
		}
	}
	for _, c := range cnt {
		if c > 0 {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1203E(os.Stdin, os.Stdout) }
