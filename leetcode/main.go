package main

import (
	"fmt"
)

// n, m := len(mat), len(mat[0])

func _permute(arr []int, i int, do func([]int)) {
	if i == len(arr) {
		do(arr)
		return
	}
	_permute(arr, i+1, do)
	for j := i + 1; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		_permute(arr, i+1, do)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func permute(arr []int, do func([]int)) { _permute(arr, 0, do) }

func largestTimeFromDigits(arr []int) (ans string) {
	permute(arr, func(a []int) {
		h, m := a[0]*10+a[1], a[2]*10+a[3]
		t := fmt.Sprintf("%02d:%02d", h, m)
		if h < 24 && m < 60 && t > ans {
			ans = t
		}
	})
	return
}

func main() {
	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}
	_ = toBytes

	fmt.Println(largestTimeFromDigits([]int{1, 1, 1, 1}))
}

func collections() {
	const mod int = 1e9 + 7
	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

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
	ifElseI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ifElseS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	_ = []interface{}{fmt.Print, ifElseI, ifElseS, dirOffset4, min, max}
}
