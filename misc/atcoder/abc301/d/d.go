package main

import (
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var s string
	var n uint64
	Fscan(in, &s, &n)
	ans, _ := strconv.ParseUint(strings.ReplaceAll(s, "?", "0"), 2, 64)
	if ans > n {
		Fprint(out, -1)
		return
	}
	for i, b := range s {
		if b == '?' && ans+1<<(len(s)-1-i) <= n {
			ans += 1 << (len(s) - 1 - i)
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
