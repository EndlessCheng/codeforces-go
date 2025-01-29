package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2040B(in io.Reader, out io.Writer) {
	var T, n uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, 1+bits.Len((n+1)/3))
	}
}

//func main() { cf2040B(bufio.NewReader(os.Stdin), os.Stdout) }

// 初步写法
func cf2040B2(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 1
		for 3<<(ans-1)-2 < n {
			ans++
		}
		Fprintln(out, ans)
	}
}
