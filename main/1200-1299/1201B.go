package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1201B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, s, v, mx int64
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v > mx {
			mx = v
		}
		s += v
	}
	if s%2 == 0 && s >= mx*2 {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

//func main() { CF1201B(os.Stdin, os.Stdout) }
