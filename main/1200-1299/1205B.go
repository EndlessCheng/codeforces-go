package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1205B(reader io.Reader, writer io.Writer) {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	const inf int = 1e8

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	arr := make([]int64, n)
	bitIndexes := make([][]int, 60)
	for i := 0; i < n; i++ {
		var a int64
		Fscan(in, &a)
		arr[i] = a
		for j := 0; a > 0; j++ {
			if a&1 == 1 {
				bitIndexes[j] = append(bitIndexes[j], i)
			}
			a >>= 1
		}
	}

	uniqueSet := map[int64]int{}
	for _, indexes := range bitIndexes {
		if len(indexes) >= 3 {
			Fprintln(out, 3)
			return
		}
		for _, idx := range indexes {
			uniqueSet[arr[idx]] = 1
		}
	}
	uniqueArr := []int64{}
	for k := range uniqueSet {
		uniqueArr = append(uniqueArr, k)
	}

	// 剩下不超过120个数，暴力求最小环即可
	sz := len(uniqueArr)
	weights := make([][]int, sz)
	for i := range weights {
		weights[i] = make([]int, sz)
		for j := range weights[i] {
			weights[i][j] = inf
		}
	}
	for i := range weights {
		weights[i][i] = 0
	}
	for i, ai := range uniqueArr {
		for j := i + 1; j < sz; j++ {
			if ai&uniqueArr[j] > 0 {
				weights[i][j] = 1
				weights[j][i] = 1
			}
		}
	}
	dist := make([][]int, sz)
	for i := range dist {
		dist[i] = make([]int, sz)
		copy(dist[i], weights[i])
	}
	ans := inf
	for k := 0; k < sz; k++ {
		for i := 0; i < k; i++ {
			for j := 0; j < i; j++ {
				ans = min(ans, dist[i][j]+weights[i][k]+weights[k][j])
			}
		}
		for i := 0; i < sz; i++ {
			for j := 0; j < sz; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1205B(os.Stdin, os.Stdout)
//}
