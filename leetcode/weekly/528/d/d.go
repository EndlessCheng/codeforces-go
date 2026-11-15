package main

import "sort"

// https://space.bilibili.com/206214
func countConsistent1(nums []int) (ans int64) {
	type pair struct{ and, left int } // 子数组与值，最小左端点
	data := []pair{}
	st := []int{}

	for i, x := range nums {
		// x 入单调栈
		for len(st) > 0 && nums[st[len(st)-1]] >= x {
			st = st[:len(st)-1]
		}
		st = append(st, i)

		// 计算右端点为 i 的子数组与值
		for j := range data {
			data[j].and &= x
		}
		// x 单独一个数作为子数组
		data = append(data, pair{x, i})

		// 原地去重（相同与值只保留最左边的）
		idx := 1
		for j := 1; j < len(data); j++ {
			if data[j].and != data[j-1].and {
				data[idx] = data[j]
				idx++
			}
		}
		data = data[:idx]

		m := len(st)
		for k, p := range data {
			andVal := p.and
			// 判断子数组最小值是否为 andVal
			j := sort.Search(m, func(j int) bool { return nums[st[j]] >= andVal })
			if j == m || nums[st[j]] != andVal {
				continue
			}

			right := i
			if k < len(data)-1 {
				right = data[k+1].left - 1
			}
			right = min(right, st[j])
			// 注：由于 nums[st[j-1]] < andVal，假如子数组包含 st[j-1]，那么与值会比 andVal 还小，矛盾
			// 所以子数组不可能包含 st[j-1]，所以无需考虑 st[j-1]
			ans += int64(max(right-p.left+1, 0))
		}
	}

	return
}

func countConsistent2(nums []int) (ans int64) {
	type pair struct{ and, left int } // 子数组与值，最小左端点
	data := []pair{}
	st := []int{}
	pos := map[int]int{}

	for i, x := range nums {
		// x 入单调栈
		for len(st) > 0 && st[len(st)-1] >= x {
			delete(pos, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, x)
		pos[x] = i

		// 计算右端点为 i 的子数组与值
		for j := range data {
			data[j].and &= x
		}
		// x 单独一个数作为子数组
		data = append(data, pair{x, i})

		// 原地去重（相同与值只保留最左边的）
		idx := 1
		for j := 1; j < len(data); j++ {
			if data[j].and != data[j-1].and {
				data[idx] = data[j]
				idx++
			}
		}
		data = data[:idx]

		for k, p := range data {
			andVal := p.and
			// 判断子数组最小值是否为 andVal
			j, ok := pos[andVal]
			if !ok {
				continue
			}

			right := i
			if k < len(data)-1 {
				right = data[k+1].left - 1
			}
			right = min(right, j)
			ans += int64(max(right-p.left+1, 0))
		}
	}

	return
}

func countConsistent(nums []int) (ans int64) {
	n := len(nums)
	left := make([]int, n)
	st := []int{-1} // 哨兵
	for i, x := range nums {
		for len(st) > 1 && nums[st[len(st)-1]]&x == x {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1] // nums[left[i]] 不是 x 的超集
		st = append(st, i)
	}

	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(st) > 1 && nums[st[len(st)-1]] != x && nums[st[len(st)-1]]&x == x {
			st = st[:len(st)-1]
		}
		right := st[len(st)-1] // nums[right] 不是 x 的超集
		st = append(st, i)
		ans += int64(i-left[i]) * int64(right-i)
	}

	return
}
