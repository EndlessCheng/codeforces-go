package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func digitSum(s string, k int) string {
	for len(s) > k {
		t := []byte{}
		for i := 0; i < len(s); i += k { // k 个一组
			sum := 0
			for _, ch := range s[i:min(i+k, len(s))] {
				sum += int(ch - '0') // 计算数字和
			}
			t = append(t, strconv.Itoa(sum)...)
		}
		s = string(t)
	}
	return s
}
