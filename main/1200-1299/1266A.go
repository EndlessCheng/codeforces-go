package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1266A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		cnt := [10]int{}
		sum := 0
		for _, c := range s {
			c -= '0'
			cnt[c]++
			sum += int(c)
		}
		if sum != 0 && (cnt[0] == 0 || cnt[0] == 1 && cnt[0]+cnt[2]+cnt[4]+cnt[6]+cnt[8] == 1 || sum%3 != 0) {
			Fprintln(out, "cyan")
		} else {
			Fprintln(out, "red")
		}
	}
}

//func main() { CF1266A(os.Stdin, os.Stdout) }
