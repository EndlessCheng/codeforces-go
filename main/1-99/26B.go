package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF26B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s string
	Fscan(in, &s)
	ans, cnt := len(s), 0
	for _, v := range s {
		if v == '(' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			ans--
		}
	}
	Fprint(out, ans-cnt)
}

//func main() { CF26B(os.Stdin, os.Stdout) }
