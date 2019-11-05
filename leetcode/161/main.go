package main

import (
	. "fmt"
	"strings"
)

var _ = Print

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func ifElseI(cond bool, a, b int) int {
	if cond {
		return a
	}
	return b
}
func ifElseS(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

const mod int = 1e9 + 7

// 0011010100
// 0110100010
//
// xxyyxyxyxx
// xyyxyxxxyx

// xxyyyyxyxx
// xxyxyxxxyx

// xxyyyxxyxx
// xxyyyxxxyx
func minimumSwap(ss1 string, ss2 string) int {
	cntx := strings.Count(ss1, "x") + strings.Count(ss2, "x")
	cnty := strings.Count(ss1, "y") + strings.Count(ss2, "y")
	if  cntx&1 == 1 || cnty&1 == 1 {
		return -1
	}
	s1 := []byte(ss1)
	s2 := []byte(ss2)
	n := len(s1)
	cnt:=0
	ans :=0
	for i := range s2 {
		if s2[i] != s1[i] {
			cnt++
			for j := i + 1; j < n; j++ {
				if s1[j]==s2[j] {
					continue
				}
				if s1[j] == s1[i] {
					s2[i], s1[j] = s1[j], s2[i]
					cnt--
					ans++
					//fmt.Println(i)
					//fmt.Println(string(s1))
					//fmt.Println(string(s2))
					break
				}
			}
		}
	}
	//fmt.Println(string(s1))
	//fmt.Println(string(s2))
	ans+=cnt
	return ans
}

func numberOfSubarrays(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	cnt := make([]int, len(nums)+5)
	cnt[0]++
	for i, v := range nums {
		if v&1 == 1 {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i]
		}
		cnt[sum[i+1]]++
	}
	ans := 0
	for _, s := range sum {
		if s >= k {
			ans += cnt[s-k]
		}
	}
	return ans
}

func minRemoveToMakeValid(ss string) string {
	s := []byte(ss)
	//n := len(ss)
	//cnt := 0
	//diff := strings.Count(ss, "(") - strings.Count(ss, ")")
	//if diff >= 0 {
	pos := []int{}
	for i, c := range s {
		if c == '(' {
			//cnt++
			pos = append(pos, i)
		} else if c == ')' {
			//cnt--
			if len(pos) <= 0 {
				s[i] = ' '
			} else {
				pos = pos[:len(pos)-1]
			}
			//if cnt < 0 {
			//	s[i] = ' '
			//}
		}
	}
	for _, p := range pos {
		s[p] = ' '
	}
	//} else {
	//	for i := n - 1; i >= 0; i-- {
	//		c := s[i]
	//		if c == ')' {
	//			cnt++
	//		} else {
	//			cnt--
	//			if cnt < 0 {
	//				s[i] = ' '
	//			}
	//		}
	//	}
	//}
	//Println(string(s))
	return strings.Replace(string(s), " ", "", -1)
}

func isGoodArray(nums []int) bool {
	if len(nums) == 1 {
		return nums[0] == 1
	}
	calcGCD := func(a, b int) int {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	val := nums[0]
	for _, v := range nums {
		val = calcGCD(val, v)
	}
	return val == 1
}

func main() {
	Println(minimumSwap("xx","yy"))
	Println(minimumSwap("xy","yx"))
	Println(minimumSwap("xx","xy"))
	Println(minimumSwap("xxyyxyxyxx","xyyxyxxxyx"))
	//fmt.Println(minimumSwap("",""))
	//Println(isGoodArray([]int{12, 5, 7, 23}))
	//Println(isGoodArray([]int{29, 6, 10}))
	//Println(isGoodArray([]int{3, 6}))
	//Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	//Println(minRemoveToMakeValid("a)b(c)d"))
	//Println(minRemoveToMakeValid("))(("))
	//Println(minRemoveToMakeValid("(a(b(c)d)"))
	//Println(minRemoveToMakeValid("(()("))
	//Println(minRemoveToMakeValid(")())"))
	//Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	//Println(numberOfSubarrays([]int{2, 4, 6}, 1))
	//Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
	//Println(numberOfSubarrays([]int{1, 2, 1}, 1))
}
