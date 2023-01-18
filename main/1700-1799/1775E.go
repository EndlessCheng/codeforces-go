package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1775E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		var v, s, mi, mx int64
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			s += v
			if s > mx {
				mx = s
			} else if s < mi {
				mi = s
			}
		}
		Fprintln(out, mx-mi)
	}
}

//func main() { CF1775E(os.Stdin, os.Stdout) }
