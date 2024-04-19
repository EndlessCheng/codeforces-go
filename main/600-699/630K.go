package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf630K(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprint(out, n-n/2-n/3-n/5-n/7+n/6+n/10+n/14+n/15+n/21+n/35-n/30-n/42-n/70-n/105+n/210)
}

//func main() { cf630K(os.Stdin, os.Stdout) }
