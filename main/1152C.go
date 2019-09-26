package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1152C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var a, b int64
	Fscan(in, &a, &b)
	if a > b {
		a, b = b, a
	}

	delta := b - a
	minLcm := int64(1e18)
	var ans int64
	calcK := func(x int64) {
		ak := (a + x - 1) / x * x
		if lcm := ak / x * (ak + delta); lcm <= minLcm {
			minLcm = lcm
			ans = ak - a
		}
	}
	for i := int64(1); i*i <= delta; i++ {
		if delta%i == 0 {
			calcK(i)
			calcK(delta / i)
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1152C(os.Stdin, os.Stdout)
//}
