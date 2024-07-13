package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	q, s := 0, ""
	Fscan(in, &s)
	sa := suffixarray.New([]byte(s))
	ans := map[string]int{}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &s)
		if res, ok := ans[s]; ok {
			Fprintln(out, res)
		} else {
			ans[s] = len(sa.Lookup([]byte(s), -1))
			Fprintln(out, ans[s])
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
