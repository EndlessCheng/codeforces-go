package main

// github.com/EndlessCheng/codeforces-go
type OrderedStream struct{}

var (
	p  int
	mp map[int]string
)

func Constructor(int) (_ OrderedStream) {
	p = 1
	mp = map[int]string{}
	return
}

func (OrderedStream) Insert(id int, value string) (ans []string) {
	mp[id] = value
	for mp[p] != "" {
		ans = append(ans, mp[p])
		p++
	}
	return
}
