package main

// github.com/EndlessCheng/codeforces-go
func minOperations(logs []string) (d int) {
	for _, s := range logs {
		if s == "../" {
			if d > 0 {
				d--
			}
		} else if s != "./" {
			d++
		}
	}
	return
}
