package main

// https://space.bilibili.com/206214
const mx = 50_001

var np = [mx]bool{1: true}

func init() {
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubarray(nums []int, k int) (ans int) {
	var minQ, maxQ []int
	last, last2 := -1, -1
	left := 0

	for i, x := range nums {
		// 1. 入
		if !np[x] {
			last2 = last
			last = i

			for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
				minQ = minQ[:len(minQ)-1]
			}
			minQ = append(minQ, i)

			for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
				maxQ = maxQ[:len(maxQ)-1]
			}
			maxQ = append(maxQ, i)

			// 2. 出
			for nums[maxQ[0]]-nums[minQ[0]] > k {
				left++
				if minQ[0] < left {
					minQ = minQ[1:]
				}
				if maxQ[0] < left {
					maxQ = maxQ[1:]
				}
			}
		}

		// 3. 更新答案
		ans += max(last2-left+1, 0)
	}

	return
}
