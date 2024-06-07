package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1362C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		for i := 0; n>>i > 0; i++ {
			ans += n >> i
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1362C(bufio.NewReader(os.Stdin), os.Stdout) }
