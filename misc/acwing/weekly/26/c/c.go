package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7
	var q, k, l, r int
	Fscan(in, &q, &k)
	f := [1e5 + 1]int{}
	for i := range f {
		if i < k {
			f[i] = 1
		} else {
			f[i] = (f[i-1] + f[i-k]) % mod
		}
	}
	for i := 1; i <= 1e5; i++ {
		f[i] += f[i-1] // 计算 f[] 的前缀和
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, (f[r]-f[l-1])%mod)
	}
}

func main() { run(os.Stdin, os.Stdout) }
