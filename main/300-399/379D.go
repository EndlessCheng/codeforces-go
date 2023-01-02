package main

import (
	. "fmt"
	"io"
	. "strings"
)

// https://space.bilibili.com/206214
func CF379D(in io.Reader, out io.Writer) {
	var k, x, n, m int
	Fscan(in, &k, &x, &n, &m)
	for c1 := 0; c1 < 2; c1++ {
		for a1 := 0; a1 < 2 && c1+a1 <= n; a1++ {
			for c2 := 0; c2 < 2; c2++ {
				for a2 := 0; a2 < 2 && c2+a2 <= m; a2++ {
					for ac1 := 0; ac1 <= (n-c1-a1)/2; ac1++ {
						for ac2 := 0; ac2 <= (m-c2-a2)/2; ac2++ {
							C1, C2, A1, AC1, AC2 := c1, c2, a1, ac1, ac2 // A2 在末尾，用不上
							for i := 2; i < k && AC2 <= x; i++ { // 防止溢出
								C1, C2, A1, AC1, AC2 = C2, C1, a2, AC2, AC1+A1&C2+AC2
							}
							if AC2 == x {
								Fprintln(out, []string{"", "C"}[c1]+Repeat("AC", ac1)+Repeat("B", n-c1-ac1*2-a1)+[]string{"", "A"}[a1])
								Fprintln(out, []string{"", "C"}[c2]+Repeat("AC", ac2)+Repeat("B", m-c2-ac2*2-a2)+[]string{"", "A"}[a2])
								return
							}
						}
					}
				}
			}
		}
	}
	Fprint(out, "Happy new year!")
}

//func main() { CF379D(os.Stdin, os.Stdout) }
