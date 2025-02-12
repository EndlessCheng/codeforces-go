package main

type BrowserHistory struct {
	history []string
	cur     int // 当前页面是 history[cur]
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.cur++
	bh.history = bh.history[:bh.cur]     // 把浏览历史前进的记录全部删除
	bh.history = append(bh.history, url) // 从当前页跳转访问 url 对应的页面
}

func (bh *BrowserHistory) Back(steps int) string {
	bh.cur = max(bh.cur-steps, 0) // 后退 steps 步
	return bh.history[bh.cur]
}

func (bh *BrowserHistory) Forward(steps int) string {
	bh.cur = min(bh.cur+steps, len(bh.history)-1) // 前进 steps 步
	return bh.history[bh.cur]
}
