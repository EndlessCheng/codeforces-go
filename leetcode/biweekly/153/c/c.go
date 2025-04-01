package main

import (
	"math"
)

// https://space.bilibili.com/206214
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

func minimumCost(nums, cost []int, k int) int64 {
	totalCost := 0
	for _, c := range cost {
		totalCost += c
	}

	q := []vec{{}}
	sumNum, sumCost := 0, 0
	for i, x := range nums {
		sumNum += x
		sumCost += cost[i]

		p := vec{-sumNum - k, 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}

		// 一边算 DP 一边构建下凸包
		p = vec{sumCost, p.dot(q[0]) + sumNum*sumCost + k*totalCost}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	return int64(q[len(q)-1].y)
}

func minimumCost1(nums, cost []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, c := range cost {
		s[i+1] = s[i] + c
	}

	f := make([]int, n+1)
	sumNum := 0
	for i, x := range nums {
		i++
		sumNum += x
		res := math.MaxInt
		for j := range i {
			res = min(res, f[j]+sumNum*(s[i]-s[j])+k*(s[n]-s[j]))
		}
		f[i] = res
	}
	return int64(f[n])
}
