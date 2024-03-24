package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1137B(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(bufio.NewReader(in), &s, &t)
	cnt := [2]int{}
	for _, b := range s {
		cnt[b&1]++
	}

	pi := make([]int, len(t))
	j := 0
	for i := 1; i < len(t); i++ {
		v := t[i]
		for j > 0 && t[j] != v {
			j = pi[j-1]
		}
		if t[j] == v {
			j++
		}
		pi[i] = j
	}

	ans := s[:0]
	j = 0
	for cnt[t[j]&1] > 0 {
		ans = append(ans, t[j])
		cnt[t[j]&1]--
		j++
		if j == len(t) {
			j = pi[j-1]
		}
	}
	Fprintf(out, "%s%s%s", ans, strings.Repeat("0", cnt[0]), strings.Repeat("1", cnt[1]))
}

//func main() { cf1137B(os.Stdin, os.Stdout) }
