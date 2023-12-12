package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf702B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	cnt := map[int]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		for i := 0; i < 31; i++ {
			ans += cnt[1<<i-v]
		}
		cnt[v]++
	}
	Fprint(out, ans)
}

//func main() { cf702B(os.Stdin, os.Stdout) }
