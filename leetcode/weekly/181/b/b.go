package main

func sumFourDivisors(nums []int) (ans int) {
	doDivisors2 := func(n int, do func(d1, d2 int)) {
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				do(d, n/d)
			}
		}
		return
	}
	for _, v := range nums {
		cnt := 0
		sum := 0
		doDivisors2(v, func(d1, d2 int) {
			if d1 != d2 {
				sum += d1 + d2
				cnt += 2
			} else {
				sum += d1
				cnt++
			}
		})
		if cnt == 4 {
			ans += sum
		}
	}
	return
}
