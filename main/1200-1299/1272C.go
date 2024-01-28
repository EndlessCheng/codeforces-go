package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1272C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, mask, ans int
	var s, t string
	Fscan(in, &n, &k, &s)
	for ; k > 0; k-- {
		Fscan(in, &t)
		mask |= 1 << (t[0] - 'a')
	}
	for i := 0; i < n; {
		if mask>>(s[i]-'a')&1 == 0 {
			i++
			continue
		}
		st := i
		for i++; i < n && mask>>(s[i]-'a')&1 > 0; i++ {
		}
		sz := i - st
		ans += sz * (sz + 1) / 2
	}
	Fprint(out, ans)
}

//func main() { cf1272C(os.Stdin, os.Stdout) }
