package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1423J(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		m := (n/4 + 1) % mod
		Fprintln(out, m*(m+n%4/2)%mod)
	}
}

//func main() { cf1423J(bufio.NewReader(os.Stdin), os.Stdout) }
