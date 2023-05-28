package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1102D(in io.Reader, out io.Writer) {
	var n int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &s)
	c0 := bytes.Count(s, []byte{'0'})
	c1 := bytes.Count(s, []byte{'1'})
	cnt := [3]int{c0 - n/3, c1 - n/3, n*2/3 - c0 - c1}
	for i, b := range s {
		b -= '0'
		if b > 0 && cnt[b] > 0 {
			if cnt[0] < 0 {
				s[i] = '0'
				cnt[0]++
				cnt[b]--
			} else if cnt[1] < 0 {
				s[i] = '1'
				cnt[1]++
				cnt[b]--
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		b := s[i] - '0'
		if cnt[b] > 0 {
			if cnt[2] < 0 {
				s[i] = '2'
				cnt[2]++
				cnt[b]--
			} else if cnt[1] < 0 {
				s[i] = '1'
				cnt[1]++
				cnt[b]--
			}
		}
	}
	Fprintf(out, "%s", s)
}

//func main() { CF1102D(os.Stdin, os.Stdout) }
