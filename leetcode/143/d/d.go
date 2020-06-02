package main

func parseBoolExpr(s string) (ans bool) {
	i := 0
	var f func(byte) bool
	f = func(tp byte) bool {
		res := tp == '&'
		g := func(b bool) {
			switch tp {
			case '!':
				res = !b
			case '&':
				res = res && b
			default: // '|'
				res = res || b
			}
		}
	o:
		for ; i < len(s); i++ {
			switch c := s[i]; c {
			case 'f', 't':
				g(c == 't')
			case '!', '&', '|':
				i += 2
				g(f(c))
			case ')':
				//i++ // TODO 注意：这里是易错点
				break o
			}
		}
		return res
	}
	return f('&')
}
