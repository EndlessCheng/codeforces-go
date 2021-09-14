package main

// 注：本题也可以 O(1) 做，维护元素的 lazy 标记，
// Increment 时 s[k-1].lazy += val
// Pop 时结算，并传递给下一个

type CustomStack struct {
	maxSize int
	s       []int
}

func Constructor(maxSize int) (c CustomStack) {
	c.maxSize = maxSize
	return
}

func (c *CustomStack) Push(x int) {
	if len(c.s) < c.maxSize {
		c.s = append(c.s, x)
	}
}

func (c *CustomStack) Pop() (ans int) {
	if len(c.s) == 0 {
		return -1
	}
	c.s, ans = c.s[:len(c.s)-1], c.s[len(c.s)-1]
	return
}

func (c *CustomStack) Increment(k int, val int) {
	if k >= len(c.s) {
		k = len(c.s)
	}
	for i := range c.s[:k] {
		c.s[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
