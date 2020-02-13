package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 1e9~1e18 √n logn 1     二分 二进制
// 1e5~1e6  nlogn nαn n   RMQ 并查集
// 1e3~1e4  n^2 n^2logn   RMQ DP
// 300~500  n^3           DP 二分图
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

}

func main() { run(os.Stdin, os.Stdout) }
