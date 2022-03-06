package main

// github.com/EndlessCheng/codeforces-go
func replaceNonCoprimes(nums []int) []int {
	s := []int{nums[0]}
	for _, num := range nums[1:] {
		s = append(s, num)
		for len(s) > 1 && gcd(s[len(s)-1], s[len(s)-2]) > 1 {
			s[len(s)-2] *= s[len(s)-1] / gcd(s[len(s)-1], s[len(s)-2])
			s = s[:len(s)-1]
		}
	}
	return s
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
