package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1753B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, v int
	Fscan(in, &n, &x)
	cnt := make([]int, x+1)
	for ; n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}
	for i := 1; i < x; i++ {
		if cnt[i]%(i+1) > 0 {
			Fprint(out, "No")
			return
		}
		cnt[i+1] += cnt[i] / (i + 1)
	}
	Fprint(out, "Yes")
}

//func main() { cf1753B(os.Stdin, os.Stdout) }
