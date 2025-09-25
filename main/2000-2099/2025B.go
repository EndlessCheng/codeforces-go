package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2025B(in io.Reader, _w io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, v int
	Fscan(in, &T)
	for range T {
		Fscan(in, &v)
	}
	for range T {
		Fscan(in, &v)
		Fprintln(out, pow(2, v))
	}
}

//func main() { cf2025B(bufio.NewReader(os.Stdin), os.Stdout) }
