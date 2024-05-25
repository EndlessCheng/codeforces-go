package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1301C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		sz := (n - m) / (m + 1)
		ex := (n - m) % (m + 1)
		Fprintln(out, n*(n+1)/2-sz*(sz+1)/2*(m+1-ex)-(sz+1)*(sz+2)/2*ex)
	}
}

//func main() { cf1301C(bufio.NewReader(os.Stdin), os.Stdout) }
