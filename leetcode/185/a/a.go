package main

func reformat(S string) (ans string) {
	d, s := []byte{}, []byte{}
	for i := range S {
		if b := S[i]; b <= '9' {
			d = append(d, b)
		} else {
			s = append(s, b)
		}
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	if abs(len(d)-len(s)) > 1 {
		return
	}
	a := []byte{}
	if len(d) > len(s) {
		v := d[0]
		d = d[1:]
		a = append(a, v)
		d, s = s, d
	} else if len(d) < len(s) {
		v := s[0]
		s = s[1:]
		a = append(a, v)
	}
	for i, b := range d {
		a = append(a, b, s[i])
	}
	return string(a)
}
