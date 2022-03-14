[比赛链接](https://leetcode-cn.com/contest/cnunionpay-2022spring/)

### Q1 贪心

设链表元素值组成的数组为 $a$，其长度为 $n$。在答案存在的情况下，假设删除了位置 $i$ 处的元素，不失一般性，假设 $i$ 位于数组中心左侧，那么 $a[:i]$ 和 $a[n-i:]$ 是对称的，并且 $a[i+1:n-i]$ 是回文的。

如果 $i$ 能尽量靠近数组中心，那么 $a[i+1:n-i]$ 越可能是回文的。

因此可以先从两端开始向中心匹配，找到第一个不同的位置 $i$，然后枚举是删除 $a[i]$ 还是 $a[n-1-i]$，判断剩下的元素是否是回文的。

```go
func isP(a []int) bool {
	for i, n := 0, len(a); i < n/2; i++ {
		v, w := a[i], a[n-1-i]
		if v != w {
			return false
		}
	}
	return true
}

func isPalindrome(head *ListNode) (ans bool) {
	a := []int{}
	for o := head; o != nil; o = o.Next {
		a = append(a, o.Val)
	}
	i, j := 0, len(a)-1
	for i < j && a[i] == a[j] {
		i++
		j--
	}
	return i >= j || isP(a[i:j]) || isP(a[i+1:j+1])
}
```

### Q2 模拟

由于至多有 $1001$ 个活动，操作数也至多有 $1000$ 个，枚举每个活动找到优惠减免最大的活动。

```go
type activity struct {
	priceLimit, discount, left, userLimit int
	userCnt map[int]int
}

type DiscountSystem []*activity

func Constructor() DiscountSystem {
	return make([]*activity, 1001)
}

func (acts DiscountSystem) AddActivity(actId, priceLimit, discount, number, userLimit int) {
	acts[actId] = &activity{priceLimit, discount, number, userLimit, map[int]int{}}
}

func (acts DiscountSystem) RemoveActivity(actId int) {
	acts[actId] = nil
}

func (acts DiscountSystem) Consume(userId, cost int) int {
	maxDiscount := -1
	var best *activity
	for _, a := range acts {
		if a != nil && a.left > 0 && a.discount > maxDiscount && a.priceLimit <= cost && a.userCnt[userId] < a.userLimit {
			maxDiscount, best = a.discount, a
		}
	}
	if best != nil {
		best.left--
		best.userCnt[userId]++
		cost -= best.discount
	}
	return cost
}
```

### Q3 二分最后一次投资的价格

如果一个一个投资，我们有一个贪心的策略：每次选择价格最高的项目投资。这可以用堆来模拟，但是本题 $\textit{limit}$ 高达 $10^9$，模拟是无法在时限内通过的。

不妨设最后一次投资的价格为 $\textit{price}$，由于 $\textit{price}$ 越大投资次数越少，因此 $\textit{price}$ 可以二分出来。

对于每个理财项目，我们知道其初始投资价格和最后一次投资的价格，那么用等差数列之和就能算出在该理财项目上的投资总和。

需要注意的是，由于 $\textit{limit}$ 的限制，可能有的项目最后一次投资的价格为 $\textit{price}+1$，因此我们需要对每个项目算出其初始投资价格到 $\textit{price}+1$ 的投资总和，然后再计算剩下的投资价格之和。

```go
func maxInvestment(product []int, limit int) (ans int) {
	price := sort.Search(1e7, func(price int) bool {
		cnt := 0
		for _, p := range product {
			if p > price {
				cnt += p - price
			}
		}
		return cnt <= limit
	})

	for _, p := range product {
		if p > price {
			cnt := p - price
			ans = (ans + (price+1+p)*cnt/2) % (1e9 + 7)
			limit -= cnt
		}
	}
	if limit > len(product) {
		limit = len(product)
	}
	return (ans + limit*price) % (1e9 + 7)
}
```

### Q4

#### 提示 1 

逆向思考：考虑不满足合作开发的成员对数。这些成员需要满足什么条件？

#### 提示 2

每个成员的技能列表长度不超过 $4$。

#### 提示 3

枚举技能列表的子集，存入一个哈希表中。

---

考虑不满足合作开发的成员对数，这在其中一个成员的技能列表是另一个成员的技能列表的子集时成立。

遍历 $\textit{skills}$，由于每个成员的技能列表长度不超过 $4$，我们可以枚举技能列表的所有非空子集，存入一个集合（哈希表）$\textit{cnt}$ 中，这样对于第 $i$ 个成员，$\textit{cnt}$ 中 $\textit{skills}[i]$ 的个数就是下标在 $i$ 前面的无法与 $i$ 合作开发的成员个数了。

为了保证后面遍历到的技能列表是前面某个技能列表的子集，我们需要将 $\textit{skills}$ 按照 $\textit{skills}[i]$ 的大小降序排序。

设 $\textit{skills}$ 的长度为 $n$，从 $\dfrac{n(n-1)}{2}$ 中减去这些无法合作开发的成员个数，即为答案。

代码实现时，由于技能值小于 $2^{10}$，$4$ 个技能仅需 $40$ 个比特存储，因此我们可以用一个 $64$ 位的整数存下每个子集。 

```go
func coopDevelop(a [][]int) int {
	ans := len(a) * (len(a) - 1) / 2
	sort.Slice(a, func(i, j int) bool { return len(a[i]) > len(a[j]) }) // 按照数组大小降序排序
	cnt := map[int]int{}
	for _, skill := range a {
		s := 0
		for _, v := range skill {
			s = s<<10 | v
		}
		ans -= cnt[s]
		// 枚举 skill 的所有非空子集
		for i := 1; i < 1<<len(skill); i++ {
			s := 0
			for j, v := range skill {
				if i>>j&1 > 0 {
					s = s<<10 | v
				}
			}
			cnt[s]++ // 加到哈希表中
		}
	}
	return ans % (1e9 + 7)
}
```
