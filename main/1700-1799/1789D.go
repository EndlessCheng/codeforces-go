package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1789D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	lsh := func(s string, k int) string { return s[k:] + strings.Repeat("0", k) }
	rsh := func(s string, k int) string { return strings.Repeat("0", k) + s[:len(s)-k] }
	xor := func(s, t string) string {
		x := []byte(s)
		for i := range x {
			x[i] ^= t[i] ^ '0'
		}
		return string(x)
	}

	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		ans := []any{}
		i := strings.IndexByte(s, '1')
		j := strings.IndexByte(t, '1')
		if i < 0 && j < 0 {
			Fprintln(out, 0)
			continue
		}
		if i < 0 || j < 0 {
			Fprintln(out, -1)
			continue
		}

		if i > j {
			ans = append(ans, i-j)
			s = xor(s, lsh(s, i-j))
			i = j
		}

		k := i + 1
		for s[i+1:] != t[i+1:] {
			for s[k] == t[k] {
				k++
			}
			ans = append(ans, i-k)
			s = xor(s, rsh(s, k-i))
		}

		r := strings.LastIndexByte(s, '1')
		k = i
		for s != t {
			for s[k] == t[k] {
				k--
			}
			ans = append(ans, r-k)
			s = xor(s, lsh(s, r-k))
		}

		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1789D(os.Stdin, os.Stdout) }
