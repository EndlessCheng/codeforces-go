package main

import (
	"sort"
)

// https://space.bilibili.com/206214
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// Graham 扫描法，计算 points 的上凸包
// 由于横坐标是严格递增的，无需排序
func convexHull(points []vec) []vec {
	q := points[:0]
	for _, p := range points {
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	return q
}

func minTime(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	vs := make([]vec, n)
	for i, x := range skill {
		s[i+1] = s[i] + x
		vs[i] = vec{s[i], x}
	}
	vs = convexHull(vs) // 去掉无用数据

	start := 0
	for j := 1; j < m; j++ {
		p := vec{mana[j-1] - mana[j], mana[j-1]}
		// p.dot(vs[i]) 是个单峰函数，二分找最大值
		i := sort.Search(len(vs)-1, func(i int) bool { return p.dot(vs[i]) > p.dot(vs[i+1]) })
		start += p.dot(vs[i])
	}
	return int64(start + mana[m-1]*s[n])
}

func minTime3(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	for i, x := range skill {
		s[i+1] = s[i] + x
	}

	suf := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		if skill[i] > skill[suf[len(suf)-1]] {
			suf = append(suf, i)
		}
	}

	pre := []int{0}
	for i := 1; i < n; i++ {
		if skill[i] > skill[pre[len(pre)-1]] {
			pre = append(pre, i)
		}
	}

	start := 0
	for j := 1; j < m; j++ {
		record := suf
		if mana[j-1] < mana[j] {
			record = pre
		}
		mx := 0
		for _, i := range record {
			mx = max(mx, mana[j-1]*s[i+1]-mana[j]*s[i])
		}
		start += mx
	}
	return int64(start + mana[m-1]*s[n])
}
