package main

type pair struct{ v, i int }
type StockSpanner struct {
	stack []pair
	days  int
}

func Constructor() (s StockSpanner) {
	s.stack = []pair{{2e9, -1}}
	return
}

func (s *StockSpanner) Next(v int) (ans int) {
	for {
		if top := s.stack[len(s.stack)-1]; top.v > v {
			ans = s.days - top.i
			break
		}
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, pair{v, s.days})
	s.days++
	return
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
