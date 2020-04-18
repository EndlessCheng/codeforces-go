package copypasta

import "sort"

// NOTE: 对于搜索格子的题，可以不用创建 vis 而是通过修改格子的值为范围外的值（如零、负数、'#' 等）来做到这一点
func searchCollection() {
	// DFS 格点找有多少个连通分量
	// 下列代码来自 LC162C https://leetcode-cn.com/problems/number-of-closed-islands/
	dfsGrids := func(grid [][]int) (comps int) {
		dir4 := [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
		n, m := len(grid), len(grid[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		var f func(i, j int) bool
		f = func(i, j int) bool {
			// 出边界的不算
			if i < 0 || i >= n || j < 0 || j >= m {
				return false
			}
			if vis[i][j] || grid[i][j] == 1 {
				return true
			}
			vis[i][j] = true
			validComp := true
			for _, dir := range dir4 {
				if !f(i+dir[0], j+dir[1]) {
					// 遍历完该连通分量再 return，保证不重不漏
					validComp = false
				}
			}
			return validComp
		}
		for i, gi := range grid {
			for j, gij := range gi {
				if gij == 0 && !vis[i][j] && f(i, j) {
					comps++
				}
			}
		}
		return
	}

	// 生成字符串 s 的所有长度至多为 r 的非空子串
	// https://codeforces.ml/problemset/problem/120/H
	genSubStrings := func(s string, r int) []string {
		a := []string{}
		var f func(s, sub string)
		f = func(s, sub string) {
			a = append(a, sub)
			if len(sub) < r {
				for i, b := range s {
					f(s[i+1:], sub+string(b))
				}
			}
		}
		f(s, "")
		a = a[1:] // 去掉空字符串
		sort.Strings(a)
		j := 0
		for i := 1; i < len(a); i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		return a[:j+1]
	}

	// 从左往右枚举排列（可以剪枝）
	// 即有 n 个位置，从左往右地枚举每个位置上可能出现的值（值必须在 sets 中），且每个位置上的元素不能重复
	// 例题见 LC169D https://leetcode-cn.com/problems/verbal-arithmetic-puzzle/
	dfsPermutations := func(n int, sets []int) bool {
		used := make([]bool, len(sets))
		//used := [10]bool{}
		var f func(cur, x, y int) bool
		f = func(pos, x, y int) bool {
			if pos == n {
				return true // custom
			}
			// 对每个位置，枚举可能出现的值，跳过已经枚举的值
			for i, v := range sets {
				_ = v
				// custom pruning
				//if  {
				//	continue
				//}
				if used[i] {
					continue
				}
				used[i] = true
				// custom calc x y
				if f(pos+1, x, y) {
					return true
				}
				used[i] = false
			}
			return false
		}
		return f(0, 0, 0)
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有组合，每个组合用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 indexes。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations
	// https://stackoverflow.com/questions/41694722/algorithm-for-itertools-combinations-in-python
	combinations := func(n, r int, do func(indexes []int)) {
		indexes := make([]int, r)
		for i := range indexes {
			indexes[i] = i
		}
		do(indexes)
		for {
			i := r - 1
			for ; i >= 0; i-- {
				if indexes[i] != i+n-r {
					break
				}
			}
			if i == -1 {
				return
			}
			indexes[i]++
			for j := i + 1; j < r; j++ {
				indexes[j] = indexes[j-1] + 1
			}
			do(indexes)
		}
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有排列，每个排列用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 indexes。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.permutations
	permutations := func(n, r int, do func(indexes []int)) {
		indexes := make([]int, n)
		for i := range indexes {
			indexes[i] = i
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

	// 生成全排列（非字典序）
	// 会修改 arr
	// Permute the values at index i to len(arr)-1.
	// https://codeforces.ml/problemset/problem/910/C
	var _permute func([]int, int, func())
	_permute = func(arr []int, i int, do func()) {
		if i == len(arr) {
			do()
			return
		}
		_permute(arr, i+1, do)
		for j := i + 1; j < len(arr); j++ {
			arr[i], arr[j] = arr[j], arr[i]
			_permute(arr, i+1, do)
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	permuteAll := func(arr []int, do func()) { _permute(arr, 0, do) }

	// 剪枝:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104613920

	// A*:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104935011

	// 舞蹈链
	// TODO: https://oi-wiki.org/search/dlx/

	_ = []interface{}{
		dfsPermutations, dfsGrids,
		genSubStrings, combinations, permutations, permuteAll,
	}
}
