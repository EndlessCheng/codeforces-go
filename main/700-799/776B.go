package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf776B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e5 + 1
	np := [mx + 1]int{}
	for i := 2; i <= mx; i++ {
		if np[i] == 0 {
			for j := i * i; j <= mx; j += i {
				np[j] = 1
			}
		}
	}
	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, strings.Repeat("1 ", n+1))
		return
	}
	Fprintln(out, 2)
	for i := 2; i <= n+1; i++ {
		Fprint(out, np[i]+1, " ")
	}
}

//func main() { cf776B(os.Stdin, os.Stdout) }
