package _00_399

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol364A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var a int
	var s string
	Fscan(in, &a, &s)
	n := len(s)
	cnt := map[int]int{0: 1}
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + int(s[i-1]-'0')
		cnt[sum[i]]++
	}

	ans := int64(0)
	if a == 0 {
		cnt0 := int64(0)
		for _, c := range cnt {
			if c > 1 {
				cnt0 += int64(c * (c - 1) / 2)
			}
		}
		for _, c := range cnt {
			if c > 1 {
				cc := int64(c * (c - 1) / 2)
				ans += 2 * cc * int64(n*(n+1)/2)
				ans -= cc * cnt0
			}
		}
		Fprint(out, ans)
		return
	}
	for d := 1; d*d <= a; d++ {
		if a%d != 0 {
			continue
		}
		cnt1 := int64(0)
		for k, c := range cnt {
			if c > 0 && k >= d {
				cnt1 += int64(c) * int64(cnt[k-d])
			}
		}
		d2 := a / d
		cnt2 := int64(0)
		for k, c := range cnt {
			if c > 0 && k >= d2 {
				cnt2 += int64(c) * int64(cnt[k-d2])
			}
		}
		if d*d == a {
			ans += cnt1 * cnt2
		} else {
			ans += 2 * cnt1 * cnt2
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol364A(os.Stdin, os.Stdout)
//}
