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
		ans := 0
		for Fscan(in, &n); n > 0; n /= 2 {
			ans += n
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1362C(bufio.NewReader(os.Stdin), os.Stdout) }
