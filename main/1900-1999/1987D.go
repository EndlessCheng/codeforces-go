package main

import (
	. "fmt"
	"io"
)

func cf1987D(in io.Reader, out io.Writer) {
	var T, n, v int16
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int16, n)
		for i := int16(0); i < n; i++ {
			Fscan(in, &v)
			cnt[v-1]++
		}
		f := make([]int16, n+1)
		for i := n - 1; i >= 0; i-- {
			c := cnt[i]
			if c == 0 {
				continue
			}
			pre := int16(0)
			for j := n - 1; j >= 0; j-- {
				res := pre + 1
				pre = f[j]
				if j >= c {
					res = min(res, f[j-c])
				}
				f[j] = res
			}
		}
		Fprintln(out, f[0])
	}
}

//func main() { cf1987D(bufio.NewReader(os.Stdin), os.Stdout) }
