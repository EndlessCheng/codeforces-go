package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2039C2(in io.Reader, out io.Writer) {
	var T, x, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &m)
		ans := 0
		l := bits.Len(uint(x))
		mask := 1<<l - 1
		for y := 1; y <= min(mask, m); y++ {
			if (x^y)%x == 0 || (x^y)%y == 0 {
				ans++
			}
		}
		if m > mask {
			// 最多循环两次
			for k := (m | mask) / x; ; k-- {
				y := k*x ^ x
				if y>>l < m>>l {
					ans += k - 1
					break
				}
				if y <= m {
					ans++
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2039C2(bufio.NewReader(os.Stdin), os.Stdout) }
