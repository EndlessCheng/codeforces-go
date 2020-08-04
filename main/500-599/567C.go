package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol567C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k, v int
	Fscan(in, &n, &k)
	if n < 3 {
		Fprint(out, 0)
		return
	}

	ans := int64(0)
	if k == 1 {
		cnt := map[int]int64{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if cnt[v] >= 2 {
				ans += cnt[v] * (cnt[v] - 1) / 2
			}
			cnt[v]++
		}
		Fprint(out, ans)
		return
	}

	cnt0 := map[int]int64{}
	cnt1 := map[int]int64{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v%k == 0 {
			if v/k%k == 0 {
				ans += cnt1[v/k]
			}
			cnt1[v] += cnt0[v/k]
		}
		cnt0[v]++
	}
	Fprint(out, ans)
}

//func main() {
//	Sol567C(os.Stdin, os.Stdout)
//}
