package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1560E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &t)
		n := len(t)
		del := []byte{}
		vis := 0
		for i := n - 1; i >= 0; i-- {
			if b := t[i]; 1<<(b-'a')&vis == 0 {
				vis |= 1 << (b - 'a')
				del = append(del, b)
			}
		}
		m := len(del)
		for i := 0; i < m/2; i++ {
			del[i], del[m-1-i] = del[m-1-i], del[i]
		}
		p := bytes.LastIndexByte(t, del[0]) + 1
		ex := 0
		for i := 1; i < m; i++ {
			ex += i * bytes.Count(t[:p], del[i:i+1])
		}
		for ; p+ex < n; p++ {
			ex += bytes.IndexByte(del, t[p])
		}
		if p+ex > n {
			Fprintln(out, -1)
			continue
		}
		s := t[:p]
		t2 := make([]byte, 0, n)
		t2 = append(t2, s...)
		for i := range del {
			s = bytes.ReplaceAll(s, del[i:i+1], nil)
			t2 = append(t2, s...)
		}
		if bytes.Compare(t2, t) != 0 {
			Fprintln(out, -1)
		} else {
			Fprintf(out, "%s %s\n", t[:p], del)
		}
	}
}

//func main() { CF1560E(os.Stdin, os.Stdout) }
