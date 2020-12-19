package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF272B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	cnt := [30]int64{}
	var n, v uint
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[bits.OnesCount(v)]++
	}
	ans := int64(0)
	for _, c := range cnt {
		ans += c * (c - 1) / 2
	}
	Fprint(out, ans)
}

//func main() { CF272B(os.Stdin, os.Stdout) }
