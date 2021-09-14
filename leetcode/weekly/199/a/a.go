package main

func restoreString(s string, indices []int) (ans string) {
	ss := make([]byte, len(s))
	for i, id := range indices {
		ss[id] = s[i]
	}
	return string(ss)
}
