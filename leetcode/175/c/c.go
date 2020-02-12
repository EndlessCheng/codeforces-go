package main

type node struct {
	lr       [2]*node
	priority uint
	key      int
	dupCnt   int
	sz       int
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) pushUp() { o.sz = o.dupCnt + o.lr[0].size() + o.lr[1].size() }

func (o *node) rotate(d int8) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.pushUp()
	x.pushUp()
	return x
}

type treap struct {
	rd   uint
	root *node
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) less(a, b int) int8 {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, dupCnt: 1, sz: 1}
	}
	if d := t.less(key, o.key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.dupCnt++
	}
	o.pushUp()
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) rank(key int) (kth int) {
	for o := t.root; o != nil; {
		switch d := t.less(key, o.key); {
		case d == 0:
			o = o.lr[0]
		case d > 0:
			kth += o.dupCnt + o.lr[0].size()
			o = o.lr[1]
		default:
			kth += o.lr[0].size()
			return
		}
	}
	return
}

func newTreap() *treap { return &treap{rd: 1} }

type TweetCounts struct {
	userTweets map[string]*treap
}

func Constructor() TweetCounts {
	return TweetCounts{map[string]*treap{}}
}

func (tc *TweetCounts) RecordTweet(tweetName string, time int) {
	t, ok := tc.userTweets[tweetName]
	if !ok {
		t = newTreap()
		tc.userTweets[tweetName] = t
	}
	t.put(time)
}

func (tc *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) (ans []int) {
	var delta int
	switch freq {
	case "minute":
		delta = 60
	case "hour":
		delta = 60 * 60
	default:
		delta = 24 * 60 * 60
	}
	ansLen := (endTime-startTime)/delta + 1
	endTime++

	t, ok := tc.userTweets[tweetName]
	if !ok {
		return make([]int, ansLen)
	}

	ans = make([]int, 0, ansLen)
	prevCnt := t.rank(startTime)
	for time := startTime + delta; time < endTime; time += delta {
		cnt := t.rank(time)
		ans = append(ans, cnt-prevCnt)
		prevCnt = cnt
	}
	cnt := t.rank(endTime)
	ans = append(ans, cnt-prevCnt)
	return
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
