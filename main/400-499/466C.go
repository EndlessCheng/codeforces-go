package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol466C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprintln(out, 0)
		return
	}
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		var val int64
		Fscan(in, &val)
		sum[i] = sum[i-1] + val
	}
	allSum := sum[n]

	if allSum%3 != 0 {
		Fprintln(out, 0)
		return
	}
	if allSum == 0 {
		cnt0 := int64(0)
		for _, s := range sum[1:n] {
			if s == 0 {
				cnt0++
			}
		}
		Fprintln(out, cnt0*(cnt0-1)/2)
		return
	}
	partSum := allSum / 3
	var ans, cnt1 int64
	for i := 1; i < n; i++ {
		if sum[i] == partSum {
			cnt1++
		} else if sum[i] == 2*partSum {
			ans += cnt1
		}
	}
	Fprintln(out, ans)
}

//func main() {
//	Sol466C(os.Stdin, os.Stdout)
//}
