package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf1864D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		diag1 := make([]byte, n*2)
		diag2 := make([]byte, n*2)
		op := make([]byte, n)
		for i := 0; i < n; i++ {
			for j := range op {
				op[j] ^= diag1[i-j+n] ^ diag2[i+j]
			}
			Fscan(in, &s)
			for j, v := range op {
				if v != s[j]&1 {
					ans++
					op[j] ^= 1
					diag1[i-j+n] ^= 1
					diag2[i+j] ^= 1
				}
			}
			// 进入下一轮循环，op 数组直接复用，这样只需要考虑两条斜线上的影响
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1864D(bufio.NewReader(os.Stdin), os.Stdout) }
