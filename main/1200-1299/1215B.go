package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1215B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, s, v int
	cnt := [2]int64{1}
	var ans0, ans1 int64
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v < 0 {
			s ^= 1
		}
		ans0 += cnt[s^1]
		ans1 += cnt[s]
		cnt[s]++
	}
	Fprint(out, ans0, ans1)
}

//func main() { CF1215B(os.Stdin, os.Stdout) }
