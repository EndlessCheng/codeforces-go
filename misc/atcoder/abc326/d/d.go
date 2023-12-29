package main

import (
	"bytes"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n int
	var r0, c0 string
	Fscan(in, &n, &c0, &r0)
	mr := make([]int, n)
	mc := make([]int, n)
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = bytes.Repeat([]byte{'.'}, n)
	}
	var f func(int, int) bool
	f = func(i, j int) bool {
		if i == n-1 && j > 0 && mc[j-1] != 7 {
			return false
		}
		if j == n {
			return mr[i] == 7 && f(i+1, 0)
		}
		if i == n {
			return true
		}
		if f(i, j+1) {
			return true
		}
		for k := byte(0); k < 3; k++ {
			if mr[i]>>k&1 > 0 || mc[j]>>k&1 > 0 || mc[j] == 0 && r0[j]-'A' != k || mr[i] == 0 && c0[i]-'A' != k {
				continue
			}
			ans[i][j] = 'A' + k
			mr[i] ^= 1 << k
			mc[j] ^= 1 << k
			if f(i, j+1) {
				return true
			}
			mr[i] ^= 1 << k
			mc[j] ^= 1 << k
		}
		ans[i][j] = '.'
		return false
	}
	if f(0, 0) {
		Fprintln(out, "Yes")
		for _, v := range ans {
			Fprintf(out, "%s\n", v)
		}
	} else {
		Fprint(out, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }
