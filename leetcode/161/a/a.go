package main

import "strings"

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
