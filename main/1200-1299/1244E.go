package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1244E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k int64
	Fscan(in, &n, &k)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	sort.Ints(arr)
	l, r := int64(0), n-1
	d := arr[r] - arr[l]
	for k > 0 && l < r {
		if l+1 <= n-r {
			if now, next := arr[l], arr[l+1]; now != next {
				if cost := int64(next-now) * (l + 1); cost <= k {
					k -= cost
					d -= next - now
				} else {
					d -= int(k / (l + 1))
					break
				}
			}
			l++
		} else {
			if now, next := arr[r], arr[r-1]; now != next {
				if cost := int64(now-next) * (n - r); cost <= k {
					k -= cost
					d -= now - next
				} else {
					d -= int(k / (n - r))
					break
				}
			}
			r--
		}
	}
	Fprint(out, d)
}

//func main() {
//	Sol1244E(os.Stdin, os.Stdout)
//}
