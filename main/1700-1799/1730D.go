package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1730D(in io.Reader, out io.Writer) {
	var T, n int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		cnt := [26]int{}
		cnt2 := [26][26]int{}
		for i, v := range s {
			v -= 'a'
			w := t[n-1-i] - 'a'
			cnt[v] ^= 1
			cnt[w] ^= 1
			cnt2[min(v, w)][max(v, w)] ^= 1
		}
		odd := 0
		for i, row := range cnt2 {
			for _, c := range row[i:] {
				odd += c
			}
		}
		if odd <= n%2 && cnt == [26]int{} {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1730D(bufio.NewReader(os.Stdin), os.Stdout) }
