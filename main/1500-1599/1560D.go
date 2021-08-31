package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF1560D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := 99
		for k := 0; k < 61; k++ {
			t := strconv.FormatInt(1<<k, 10)
			j := 0
			for _, b := range s {
				if j < len(t) && t[j] == b {
					j++
				}
			}
			ans = min(ans, len(s)+len(t)-j*2)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1560D(os.Stdin, os.Stdout) }
