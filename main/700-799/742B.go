package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf742B(in io.Reader, out io.Writer) {
	var n, x, v, ans int
	Fscan(in, &n, &x)
	cnt := [1 << 17]int{}
	for range n {
		Fscan(in, &v)
		ans += cnt[v^x]
		cnt[v]++
	}
	Fprint(out, ans)
}

//func main() { cf742B(bufio.NewReader(os.Stdin), os.Stdout) }
