package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	pos := make([]int, n+1)
	L := make([]int, n+2)
	R := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		pos[v] = i
		L[i] = i - 1
		R[i] = i + 1
	}
	for v := 1; v <= n; v++ {
		i := pos[v]
		l, r := L[i], R[i]
		if l > 0  { ans += v * (l - L[l]) * (r - i) }
		if r <= n { ans += v * (R[r] - r) * (i - l) }
		L[r] = l
		R[l] = r // 删除 v
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
