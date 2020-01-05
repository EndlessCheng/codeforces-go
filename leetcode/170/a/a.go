package main

func freqAlphabets(s string) string {
	isd := make([]bool, len(s))
	for i, c := range s {
		if c == '#' {
			isd[i] = true
			isd[i-1] = true
			isd[i-2] = true
		}
	}
	ans := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !isd[i] {
			ans = append(ans, c-'1'+'a')
		} else {
			v := c - '0'
			v2 := s[i+1] - '0'
			v = v*10 + v2
			ans = append(ans, v+'a'-1)
			i += 2
		}
	}
	return string(ans)
}
