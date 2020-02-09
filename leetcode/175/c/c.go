package main

type node struct {
	lr       [2]*node
	priority uint
	msz      int
	key      int
	dupCnt   int
}

func (o *node) mSize() int {
	if o != nil {
		return o.msz
	}
	return 0
}

func (o *node) pushUp() {
	msz := o.dupCnt
	if ol := o.lr[0]; ol != nil {
		msz += ol.msz
	}
	if or := o.lr[1]; or != nil {
		msz += or.msz
	}
	o.msz = msz
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	x.msz = o.msz
	o.pushUp()
	return x
}

type treap struct {
	rd   uint
	root *node
}

func newTreap() *treap {
	return &treap{rd: 1}
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) compare(a, b int) int {
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
		return &node{priority: t.fastRand(), msz: 1, key: key, dupCnt: 1}
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	} else {
		o.dupCnt++
	}
	o.pushUp()
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) mRank(key int) (cnt int) {
	for o := t.root; o != nil; {
		switch cmp := t.compare(key, o.key); {
		case cmp == 0:
			o = o.lr[0]
		case cmp > 0:
			cnt += o.dupCnt + o.lr[0].mSize()
			o = o.lr[1]
		default:
			cnt += o.lr[0].mSize()
			return
		}
	}
	return
}

type TweetCounts struct {
	userTweets map[string]*treap
}

func Constructor() TweetCounts {
	return TweetCounts{
		userTweets: map[string]*treap{},
	}
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
	prevCnt := t.mRank(startTime)
	for time := startTime + delta; time < endTime; time += delta {
		cnt := t.mRank(time)
		ans = append(ans, cnt-prevCnt)
		prevCnt = cnt
	}
	cnt := t.mRank(endTime)
	ans = append(ans, cnt-prevCnt)
	return
}
