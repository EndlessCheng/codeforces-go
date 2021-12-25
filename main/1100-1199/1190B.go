package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1190B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	sum := int(int64(n) * int64(n-1) / 2 & 1)
	cnt := make(map[int]int8, n)
	for dup := false; n > 0; n-- {
		Fscan(in, &v)
		if cnt[v]++; cnt[v] > 1 {
			if dup || v == 0 || cnt[v-1] > 0 {
				Fprint(out, "cslnb")
				return
			}
			dup = true
		}
		if cnt[v+1] > 1 {
			Fprint(out, "cslnb")
			return
		}
		sum ^= v & 1
	}
	if sum == 0 {
		Fprint(out, "cslnb")
	} else {
		Fprint(out, "sjfnb")
	}
}

//func main() { CF1190B(os.Stdin, os.Stdout) }
