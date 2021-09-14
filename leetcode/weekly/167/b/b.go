package main

import "strconv"

func sequentialDigits(low int, high int) (ans []int) {
	all := []int{}
	for sz := 2; sz <= 9; sz++ {
		for st := 1; st+sz-1 <= 9; st++ {
			s := []byte{}
			for i := 0; i < sz; i++ {
				s = append(s, byte('0'+st+i))
			}
			v, _ := strconv.Atoi(string(s))
			all = append(all, v)
		}
	}
	for _, v := range all {
		if v >= low && v <= high {
			ans = append(ans, v)
		}
	}
	return
}
