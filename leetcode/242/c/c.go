package main

// github.com/EndlessCheng/codeforces-go
type fenwick struct{ tree []int }

func (f fenwick) inc(i int) {
	for i++; i < len(f.tree); i += i & -i {
		f.tree[i]++
	}
}
func (f fenwick) sum(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}
func (f fenwick) query(l, r int) (res int) {
	return f.sum(r) - f.sum(l-1)
}

func canReach(s string, minJump, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	f := fenwick{make([]int, n)}
	f.inc(0)
	for i := minJump; i < n; i++ {
		if s[i] == '0' && f.query(i-maxJump, i-minJump) > 0 {
			if i == n-1 {
				return true
			}
			f.inc(i)
		}
	}
	return false
}
