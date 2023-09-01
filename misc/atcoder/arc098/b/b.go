package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, left, or, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		for or&a[i] > 0 {
			or ^= a[left]
			left++
		}
		or |= a[i]
		ans += i - left + 1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
