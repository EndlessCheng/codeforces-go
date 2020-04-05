package main

func numSteps(ss string) (ans int) {
	s := []byte("0" + ss)
	for string(s) != "01" && string(s) != "1" {
		ans++
		if s[len(s)-1] == '0' {
			s = s[:len(s)-1]
		} else {
			s[len(s)-1]--
			for i := len(s) - 2; i >= 0; i-- {
				if s[i] == '1' {
					s[i]--
				} else {
					s[i]++
					break
				}
			}
		}
	}
	return
}
