package main

import "fmt"

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	n := len(nums)
	cnt, x := make([]int, n*2), n
	cnt[x] = 1
	for i := pos - 1; i >= 0; i-- {
		if nums[i] < k {
			x++
		} else {
			x--
		}
		cnt[x]++
	}

	ans, x := cnt[n]+cnt[n-1], n
	for _, v := range nums[pos+1:] {
		if v > k {
			x++
		} else {
			x--
		}
		ans += cnt[x] + cnt[x-1]
	}
	return ans
}

/*
2 1 [0 1] 0
3 2 [0 1 2] 1
4 3 [0 2 3 1] 2
5 5 [0 3 2 1 4] 2
6 6 [0 4 1 3 5 2] 3
7 8 [0 4 1 3 5 2 6] 3
8 10 [0 5 1 4 6 2 7 3] 4
9 13 [0 5 1 6 4 2 7 3 8] 4
10 15 [0 6 1 7 2 5 8 3 9 4] 5
11 18 [0 6 1 7 2 5 8 3 9 4 10] 5

2 2 [0 1] 0
3 3 [0 1 2] 1
4 5 [0 2 1 3] 1
5 8 [0 3 2 4 1] 2
6 10 [0 3 2 4 1 5] 2
7 13 [0 4 1 5 3 6 2] 3
8 16 [0 4 1 5 3 6 2 7] 3
9 21 [0 5 1 6 4 7 2 8 3] 4
10 24 [0 5 1 6 4 7 2 8 3 9] 4
11 29 [0 6 1 7 2 8 5 9 3 10 4] 5

*/
func main() {
	for n := 2; n <= 9; n++ {
		a := make([]int, n)
		k := (n - 1) / 2
		a[k] = k
		cur := 0
		for i := 0; i < n; i += 2 {
			if i != k {
				a[i] = cur
				cur++
			}
		}
		cur = k + 1
		for i := 1; i < n; i += 2 {
			a[i] = cur
			cur++
		}
		for i := range a {
			a[i]++
		}
		fmt.Println(n, countSubarrays(a, k))
	}
	return

	permutations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		if do(ids[:r]) {
			return
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := ids[i]
					copy(ids[i:], ids[i+1:])
					ids[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					ids[i], ids[n-j] = ids[n-j], ids[i]
					if do(ids[:r]) {
						return
					}
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}
	for n := 5; n <= 5; n++ {
		mx := -1
		mxA := [][]int{}
		mxK := 0
		permutations(n, n, func(ids []int) (Break bool) {
			for k := 0; k < n; k++ {
				res := countSubarrays(ids, k)
				if res > mx {
					mx = res
					mxA = [][]int{append(ids[:0:0], ids...)}
					mxK = k
				} else if res == mx {
					mxA = append(mxA, append(ids[:0:0], ids...))
				}
			}
			return
		})
		fmt.Println(n, mx, mxA, mxK)
	}

}
