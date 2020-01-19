package main

import "strconv"

func maximum69Number(num int) (ans int) {
	s := []byte(strconv.Itoa(num))
	for i, b := range s {
		if b == '6' {
			s[i] = '9'
			break
		}
	}
	ans, _ = strconv.Atoi(string(s))
	return
}
