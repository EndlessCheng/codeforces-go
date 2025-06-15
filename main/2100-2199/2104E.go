package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2104E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, q, mask int
	var s, t string
	Fscan(in, &n, &k, &s)

	nxt := make([][26]int, n+1)
	for j := range nxt[n] {
		nxt[n][j] = n
	}
	f := make([]int, n+2)
	f[n] = 1
	for i := n - 1; i >= 0; i-- {
		nxt[i] = nxt[i+1]
		nxt[i][s[i]-'a'] = i
		mask |= 1 << (s[i] - 'a')
		f[i] = f[i+1]
		if mask == 1<<k-1 {
			f[i]++
			mask = 0
		}
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &t)
		i := -1
		for _, b := range t {
			i = nxt[i+1][b-'a']
			if i == n {
				break
			}
		}
		Fprintln(out, f[i+1])
	}
}

//func main() { cf2104E(bufio.NewReader(os.Stdin), os.Stdout) }
