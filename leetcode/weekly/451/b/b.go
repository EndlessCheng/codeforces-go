package main

// https://space.bilibili.com/206214
func isConsecutive(x, y byte) bool {
	d := abs(int(x) - int(y))
	return d == 1 || d == 25
}

func resultingString(s string) string {
	st := []byte{}
	for _, b := range s {
		st = append(st, byte(b))
		if len(st) > 1 && isConsecutive(st[len(st)-2], st[len(st)-1]) {
			st = st[:len(st)-2]
		}
	}
	return string(st)
}

func abs(x int) int { if x < 0 { return -x }; return x }
