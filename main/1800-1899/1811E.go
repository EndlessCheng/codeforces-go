package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
)

// https://space.bilibili.com/206214
func CF1811E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, k uint64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k)
		s := []byte(strconv.FormatUint(k, 9))
		for i, c := range s {
			if c >= '4' {
				s[i]++
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { CF1811E(os.Stdin, os.Stdout) }
