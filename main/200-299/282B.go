package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF282B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, a, g int
	Fscan(in, &n)
	ans := bytes.Repeat([]byte{'A'}, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v, &w)
		if a+v-g <= 500 {
			a += v
		} else {
			g += w
			ans[i] = 'G'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF282B(os.Stdin, os.Stdout) }
