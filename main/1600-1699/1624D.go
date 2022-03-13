package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1624D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		pair, odd := 0, 0
		for _, c := range cnt {
			pair += c / 2
			odd += c & 1
		}
		ans := pair / k * 2
		if odd+pair%k*2 >= k {
			ans++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1624D(os.Stdin, os.Stdout) }
