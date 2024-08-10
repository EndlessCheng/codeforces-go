package main

import (
	"bytes"
	. "fmt"
	"io"
)

func cf716B(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	cnt := [32]int{}
	ok := func() bool {
		for _, c := range cnt[1:27] {
			if c > 1 {
				return false
			}
		}
		return true
	}
	for r, c := range s {
		cnt[c&31]++
		l := r - 25
		if l < 0 {
			continue
		}
		if !ok() {
			cnt[s[l]&31]--
			continue
		}
		for i, j := l, 1; i <= r; i++ {
			if s[i] != '?' {
				continue
			}
			for cnt[j] > 0 {
				j++
			}
			s[i] = 'A' - 1 + byte(j)
			j++
		}
		Fprintf(out, "%s", bytes.ReplaceAll(s, []byte{'?'}, []byte{'A'}))
		return
	}
	Fprint(out, -1)
}

//func main() { cf716B(bufio.NewReader(os.Stdin), os.Stdout) }
