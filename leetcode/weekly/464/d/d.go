package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxWalls1(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	memo := make([][2]int, n+1)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}

		// 往左射，墙的范围为 [leftX, a[i].x]
		leftX := max(a[i].x-a[i].d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		left := sort.SearchInts(walls, leftX)
		cur := sort.SearchInts(walls, a[i].x+1)
		res := dfs(i-1, 0) + cur - left // [left, cur-1] 中的墙都能摧毁

		// 往右射，墙的范围为 [a[i].x, rightX]
		x2 := a[i+1].x
		if j == 0 { // 右边那个机器人往左射
			x2 -= a[i+1].d
		}
		rightX := min(a[i].x+a[i].d, x2-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		right := sort.SearchInts(walls, rightX+1)
		cur = sort.SearchInts(walls, a[i].x)
		res = max(res, dfs(i-1, 1)+right-cur) // [cur, right-1] 中的墙都能摧毁

		*p = res
		return res
	}
	return dfs(n, 1)
}

func maxWalls2(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	f := [2]int{}
	for i := 1; i <= n; i++ {
		p := a[i]
		// 往左射，墙的坐标范围为 [leftX, p.x]
		leftX := max(p.x-p.d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		left := sort.SearchInts(walls, leftX)
		cur := sort.SearchInts(walls, p.x+1)
		leftRes := f[0] + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		cur = sort.SearchInts(walls, p.x)
		for j := range 2 {
			// 往右射，墙的坐标范围为 [p.x, rightX]
			x2 := a[i+1].x
			if j == 0 { // 右边那个机器人往左射
				x2 -= a[i+1].d
			}
			rightX := min(p.x+p.d, x2-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
			right := sort.SearchInts(walls, rightX+1)
			f[j] = max(leftRes, f[1]+right-cur) // 下标在 [cur, right-1] 中的墙都能摧毁
		}
	}
	return f[1]
}

func maxWalls(robots []int, distance []int, walls []int) int {
	n, m := len(robots), len(walls)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	var f0, f1, left, cur, right0, right1 int
	for i := 1; i <= n; i++ {
		p := a[i]
		// 往左射，墙的坐标范围为 [leftX, p.x]
		leftX := max(p.x-p.d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		for left < m && walls[left] < leftX {
			left++
		}
		for cur < m && walls[cur] < p.x {
			cur++
		}
		cur1 := cur
		if cur < m && walls[cur] == p.x {
			cur++
		}
		leftRes := f0 + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		// 往右射，右边那个机器人往左射，墙的坐标范围为 [p.x, rightX]
		q := a[i+1]
		rightX := min(p.x+p.d, q.x-q.d-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		for right0 < m && walls[right0] <= rightX {
			right0++
		}
		f0 = max(leftRes, f1+right0-cur1) // 下标在 [cur1, right0-1] 中的墙都能摧毁

		// 往右射，右边那个机器人往右射，墙的坐标范围为 [p.x, rightX]
		rightX = min(p.x+p.d, q.x-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		for right1 < m && walls[right1] <= rightX {
			right1++
		}
		f1 = max(leftRes, f1+right1-cur1) // 下标在 [cur1, right0-1] 中的墙都能摧毁
	}
	return f1
}
