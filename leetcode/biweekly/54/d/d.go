package main

// github.com/EndlessCheng/codeforces-go
type minFlip struct{ f0, f1 int }

func minOperationsToFlip(s string) int {
	a := []minFlip{}
	ops := []byte{}
	for _, c := range s {
		if c == '0' || c == '1' || c == ')' {
			if c == '0' {
				a = append(a, minFlip{0, 1})
			} else if c == '1' {
				a = append(a, minFlip{1, 0})
			} else {
				ops = ops[:len(ops)-1]
			}
			if len(ops) > 0 && ops[len(ops)-1] != '(' {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				p := a[len(a)-1]
				y0, y1 := p.f0, p.f1
				p = a[len(a)-2]
				x0, x1 := p.f0, p.f1
				a = a[:len(a)-2]
				if op == '&' {
					a = append(a, minFlip{min(x0, y0), min(x1+y1, 1+min(x1, y1))})
				} else {
					a = append(a, minFlip{min(x0+y0, 1+min(x0, y0)), min(x1, y1)})
				}
			}
		} else {
			ops = append(ops, byte(c))
		}
	}
	return max(a[0].f0, a[0].f1) // 大的那个就是翻转的，因为小的那个翻转次数为 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
