package main

// github.com/EndlessCheng/codeforces-go
func cellsInRange(s string) []string {
	ans := make([]string, 0, (s[3]-s[0]+1)*(s[4]-s[1]+1))
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			ans = append(ans, string(i)+string(j))
		}
	}
	return ans
}
