本场比赛 [视频讲解](https://www.bilibili.com/video/BV1fP4y1d7Mn) 已出炉，欢迎素质三连，在评论区分享你对这场比赛的看法~

---

# [银联-4. 设计自动售货机](https://leetcode.cn/contest/cnunionpay2022/problems/NyZD2B/)

这题数据范围比较小，暴力枚举是可以通过的。

但是本题有插入元素和查找最小元素的操作，这非常适合用最小堆来做，那么是否有时间复杂度为 $O(q\log q)$ 的做法呢？（$q$ 为调用次数）

麻烦之处在于，我们除了需要维护按照（价格，过期时间）双关键字排序的堆以外，还需要维护一个按照过期时间排序的堆，能否做到同时维护这两个堆呢？

是可以的，关键之处在于，让这两个堆存储同一份元素（Python/Java 可以用对象，C++/Go 可以用指针），这样可以在移除超过过期时间的商品和购买商品时，修改商品的个数，同时修改另一个堆中的元素。

具体细节见代码。

```py [sol1-Python3]
class ExpireList:
    __slots__ = 'lst'

    def __init__(self, lst):
        self.lst = lst

    def __lt__(self, other):
        return self.lst[1] < other.lst[1]

class Good:
    __slots__ = 'data', 'expire', 'left'

    def __init__(self):
        self.data = []
        self.expire = []
        self.left = 0

class VendingMachine:
    def __init__(self):
        self.good = defaultdict(Good)
        self.discount = defaultdict(lambda: 100)

    def addItem(self, time: int, number: int, item: str, price: int, duration: int) -> None:
        it = self.good[item]
        lst = [price, time + duration, number]
        heappush(it.data, lst)
        heappush(it.expire, ExpireList(lst))
        it.left += number

    def sell(self, time: int, customer: str, item: str, number: int) -> int:
        it = self.good[item]
		# 清除过期商品
        while it.expire and it.expire[0].lst[1] < time:
            lst = heappop(it.expire).lst
            it.left -= lst[2]
            lst[2] = 0  # 懒删除 it.data 中的元素

        if it.left < number:
            return -1
        it.left -= number

        # 计算花费
        cost = 0
        while it.data:
            lst = it.data[0]
            if lst[2] >= number:
                cost += number * lst[0]
                lst[2] -= number
                break
            cost += lst[2] * lst[0]
            number -= lst[2]
            lst[2] = 0  # 懒删除 it.expire 中的元素
            heappop(it.data)

        # 计算折扣
        ans = (cost * self.discount[customer] + 99) // 100
        if self.discount[customer] > 70:
            self.discount[customer] -= 1
        return ans
```

```go [sol1-Go]
type Good struct {
	data   hp
	expire hp2
	left   int
}

type VendingMachine struct {
	good     map[string]*Good
	discount map[string]int
}

func Constructor() VendingMachine {
	return VendingMachine{map[string]*Good{}, map[string]int{}}
}

func (v VendingMachine) AddItem(time int, number int, item string, price int, duration int) {
	if v.good[item] == nil {
		v.good[item] = &Good{}
	}
	it := v.good[item]
	t := &tuple{price, time + duration, number}
	heap.Push(&it.data, t)
	heap.Push(&it.expire, t)
	it.left += number
}

func (v VendingMachine) Sell(time int, customer string, item string, number int) int64 {
	it := v.good[item]
	if it == nil {
		return -1
	}

	// 清除过期商品
	for len(it.expire) > 0 && it.expire[0].expire < time {
		t := heap.Pop(&it.expire).(*tuple)
		it.left -= t.left
		t.left = 0 // 懒删除 it.data 中的元素
	}

	if it.left < number {
		return -1
	}
	it.left -= number

	// 计算花费
	cost := 0
	for len(it.data) > 0 {
		t := it.data[0]
		if t.left >= number {
			cost += number * t.price
			t.left -= number
			break
		}
		cost += t.left * t.price
		number -= t.left
		t.left = 0 // 懒删除 it.expire 中的元素
		heap.Pop(&it.data)
	}

	// 计算折扣
	if v.discount[customer] == 0 {
		v.discount[customer] = 100
	}
	ans := (cost*v.discount[customer] + 99) / 100
	if v.discount[customer] > 70 {
		v.discount[customer]--
	}
	return int64(ans)
}

type tuple struct{ price, expire, left int }
type hp []*tuple
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.price < b.price || a.price == b.price && a.expire < b.expire }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(*tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
type hp2 []*tuple
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].expire < h[j].expire }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(*tuple)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```