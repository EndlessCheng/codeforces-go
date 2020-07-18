package main

// github.com/EndlessCheng/codeforces-go
func solve(n int, a []int) string {
	min := int(1e9)
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	l := n / min
	if l == 0 {
		return "-1"
	}
	ans := []byte{}
	for ; l >= 0; l-- {
		for i := byte(9); i > 0; i-- {
			if v := a[i-1]; v <= n && (n-v)/min == l {
				n -= v
				ans = append(ans, '0'+i)
				break
			}
		}
	}
	return string(ans)
}
