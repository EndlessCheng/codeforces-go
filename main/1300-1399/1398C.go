package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1398C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		c := map[int]int{0: 1}
		ans := int64(0)
		sum := 0
		for i, b := range s {
			sum += int(b & 15)
			ans += int64(c[sum-i-1])
			c[sum-i-1]++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1398C(os.Stdin, os.Stdout) }
