package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1187B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	var s string
	Fscan(in, &n, &s)
	pos := [26][]int{}
	for i, b := range s {
		pos[b-'a'] = append(pos[b-'a'], i)
	}
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &s)
		cnt := [26]int{}
		mx := 0
		for _, b := range s {
			b -= 'a'
			mx = max(mx, pos[b][cnt[b]])
			cnt[b]++
		}
		Fprintln(out, mx+1)
	}
}

//func main() { cf1187B(bufio.NewReader(os.Stdin), os.Stdout) }
