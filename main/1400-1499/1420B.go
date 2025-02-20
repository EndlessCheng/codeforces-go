package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1420B(in io.Reader, out io.Writer) {
	var T, n, v uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		cnt := [30]int{}
		for range n {
			Fscan(in, &v)
			b := bits.Len(v) - 1
			ans += cnt[b]
			cnt[b]++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1420B(bufio.NewReader(os.Stdin), os.Stdout) }
