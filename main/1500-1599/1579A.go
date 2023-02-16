package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1579A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		if strings.Count(s, "B")*2 == len(s) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1579A(os.Stdin, os.Stdout) }
