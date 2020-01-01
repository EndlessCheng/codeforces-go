package main

import (
	. "fmt"
)

// n, m := len(mat), len(mat[0])

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func leastOpsExpressTarget(x, target int) int {
	type pair struct{ x, y int }
	dp := map[pair]int{}
	var f func(int, int) int
	f = func(i, tar int) (ans int) {
		p := pair{i, tar}
		if v, ok := dp[p]; ok {
			return v
		}
		defer func() { dp[p] = ans }()
		c := i
		if i == 0 {
			c = 2
		}
		switch {
		case tar == 0:
			return 0
		case tar == 1:
			return c
		case i >= 30:
			return 1e9
		default:
			tar, r := tar/x, tar%x
			return min(c*r+f(i+1, tar), c*(x-r)+f(i+1, tar+1))
		}
	}
	return f(0, target) - 1
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

	Println(leastOpsExpressTarget(2, 125046))
	//Println(leastOpsExpressTarget(3, 19))
	//Println(leastOpsExpressTarget(5, 501))
	//Println(leastOpsExpressTarget(100, 100000000))
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

	_ = []interface{}{Print, ifElseI, ifElseS, dirOffset4, min, max}
}
