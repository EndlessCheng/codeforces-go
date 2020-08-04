package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol118D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n1, n2, k1, k2 int
	Fscan(in, &n1, &n2, &k1, &k2)

	const mod = 1e8
	cache := map[string]int{}
	var f func(left1, left2, c1, c2 int) int
	f = func(left1, left2, c1, c2 int) (sum int) {
		if left1 == 0 {
			if left2+c2 <= k2 {
				return 1
			}
			return 0
		}
		if left2 == 0 {
			if left1+c1 <= k1 {
				return 1
			}
			return 0
		}
		hash := Sprintf("%d;%d;%d;%d", left1, left2, c1, c2)
		if val, ok := cache[hash]; ok {
			return val
		}
		if c1 < k1 {
			sum += f(left1-1, left2, c1+1, 0)
		}
		if c2 < k2 {
			sum += f(left1, left2-1, 0, c2+1)
		}
		sum %= mod
		cache[hash] = sum
		return
	}
	Fprintln(out, f(n1, n2, 0, 0))
}

//func main() {
//	Sol118D(os.Stdin, os.Stdout)
//}
