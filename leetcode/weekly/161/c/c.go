package main

import "strings"

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
