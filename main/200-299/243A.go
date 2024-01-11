package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf243A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	ans := map[int]struct{}{}
	opRes := []int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		for j := range opRes {
			opRes[j] |= v
		}
		opRes = append(opRes, v)
		opRes = slices.Compact(opRes)
		for _, w := range opRes {
			ans[w] = struct{}{}
		}
	}
	Fprint(out, len(ans))
}

//func main() { cf243A(os.Stdin, os.Stdout) }
