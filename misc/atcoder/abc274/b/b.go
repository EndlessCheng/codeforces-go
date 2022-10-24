package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n ,m int
	Fscan(in, &n,&m)
	cnt := make([]int, m)
	for i := 0; i < n; i++ {
		s := ""
		Fscan(in, &s)
		for j, c := range s {
			if c == '#' {
				cnt[j]++
			}
		}
	}
	for _, v := range cnt {
		Fprint(out, v, " ")
	}
	Fprintln(out)
}

func main() { run(os.Stdin, os.Stdout) }
