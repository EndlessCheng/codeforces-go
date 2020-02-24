package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1313B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func() (ans1, ans2 int) {
		var n, x, y int
		Fscan(in, &n, &x, &y)
		if x > y {
			x, y = y, x
		}
		if y == 1 {
			return 1, 1
		}

		if x >= n-1 {
			ans1 = y // 比赛时这里我写成了 n
		} else {
			if x+y < n+1 {
				ans1 = 1
			} else {
				remove := n - x - 1
				ans1 = y - remove
			}
		}
		if x+y < n+1 {
			ans2 = x + y - 1
		} else {
			ans2 = n
		}
		return
	}
	_ = solve

	solve2 := func() (ans1, ans2 int) {
		var n, x, y int
		Fscan(in, &n, &x, &y)
		if x+y <= n {
			return 1, x + y - 1
		}
		if x+y < 2*n {
			return x + y - n + 1, n
		}
		return n, n
	}
	_ = solve2

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		a, b := solve2()
		Fprintln(out, a, b)
	}
}

// 打表就行了
// 2 1 1 1 1
// 3 1 2 1 2
// 3 2 1 1 2
// 4 1 3 1 3
// 4 2 2 1 3
// 4 3 1 1 3
// 5 1 4 1 4
// 5 2 3 1 4
// 5 3 2 1 4
// 5 4 1 1 4
// 6 1 5 2 5
// 6 2 4 2 5
// 6 3 3 2 5
// 6 4 2 2 5
// 6 5 1 2 5
// 7 2 5 3 5
// 7 3 4 3 5
// 7 4 3 3 5
// 7 5 2 3 5
// 8 3 5 4 5
// 8 4 4 4 5
// 8 5 3 4 5
// 9 4 5 5 5
// 9 5 4 5 5
//10 5 5 5 5
func table(n int) {
	permutations := func(n, r int, do func(indexes []int)) {
		indexes := make([]int, n)
		for i := range indexes {
			indexes[i] = i + 1
		}
		do(indexes[:r])
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := indexes[i]
					copy(indexes[i:], indexes[i+1:])
					indexes[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indexes[i], indexes[n-j] = indexes[n-j], indexes[i]
					do(indexes[:r])
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	getPlace := func(indexes1, indexes2 []int) (ans int) {
		score := indexes1[0] + indexes2[0]
		for i := 1; i < 5; i++ {
			if indexes1[i]+indexes2[i] <= score {
				ans++
			}
		}
		return ans + 1
	}

	for sum := 0; sum < 2*n-1; sum++ {
		for x := 0; x <= sum; x++ {
			y := sum - x
			xx, yy := x+1, y+1
			if xx > n || yy > n {
				continue
			}
			minPlace, maxPlace := int(1e9), 0
			permutations(n, n, func(indexes1 []int) {
				if indexes1[0] == xx {
					permutations(n, n, func(indexes2 []int) {
						if indexes2[0] == yy {
							p := getPlace(indexes1, indexes2)
							minPlace = min(minPlace, p)
							maxPlace = max(maxPlace, p)
						}
					})
				}
			})
			Println(xx+yy, xx, yy, minPlace, maxPlace)
		}
	}
}

func main() { CF1313B(os.Stdin, os.Stdout) }
