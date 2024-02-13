package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1029B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, cnt, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if i > 0 && a[i] > a[i-1]*2 {
			cnt = 0
		}
		cnt++
		ans = max(ans, cnt)
	}
	Fprint(out, ans)
}

//func main() { cf1029B(os.Stdin, os.Stdout) }
