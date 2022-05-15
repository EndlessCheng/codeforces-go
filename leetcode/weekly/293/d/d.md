思路同 [715. Range 模块](https://leetcode.cn/problems/range-module/)。

一种做法是 [动态开点线段树](https://zhuanlan.zhihu.com/p/246255556)。

对于本题来说，线段树的每个节点可以保存对应范围的左右端点 $l$ 和 $r$，以及范围内 `add` 过的整数个数 $\textit{sum}$。

代码实现时，无需记录 lazy tag，这是因为被覆盖的范围无需再次覆盖，因此若 $\textit{sum}$ 等于范围的长度 $r-l+1$，则可直接返回。

```Python [sol1-Python3]
class CountIntervals:
    def __init__(self, l=1, r=10 ** 9):
        self.left = self.right = None
        self.l, self.r, self.sum = l, r, 0

    def add(self, l: int, r: int) -> None:
        if self.sum == self.r - self.l + 1: return  # self 已被完整覆盖，无需执行任何操作
        if l <= self.l and self.r <= r:  # self 已被区间 [l,r] 完整覆盖，不再继续递归
            self.sum = self.r - self.l + 1
            return
        mid = (self.l + self.r) // 2
        if self.left is None: self.left = CountIntervals(self.l, mid)  # 动态开点
        if self.right is None: self.right = CountIntervals(mid + 1, self.r)  # 动态开点
        if l <= mid: self.left.add(l, r)
        if mid < r: self.right.add(l, r)
        self.sum = self.left.sum + self.right.sum

    def count(self) -> int:
        return self.sum
```

```java [sol1-Java]
class CountIntervals {
    CountIntervals left, right;
    int l, r, sum;

    public CountIntervals() {
        l = 1;
        r = (int) 1e9;
    }

    CountIntervals(int l, int r) {
        this.l = l;
        this.r = r;
    }

    public void add(int L, int R) { // 为方便区分变量名，将递归中始终不变的入参改为大写（视作常量）
        if (sum == r - l + 1) return; // 当前节点已被完整覆盖，无需执行任何操作
        if (L <= l && r <= R) { // 当前节点已被区间 [L,R] 完整覆盖，不再继续递归
            sum = r - l + 1;
            return;
        }
        int mid = (l + r) / 2;
        if (left == null) left = new CountIntervals(l, mid); // 动态开点
        if (right == null) right = new CountIntervals(mid + 1, r); // 动态开点
        if (L <= mid) left.add(L, R);
        if (mid < R) right.add(L, R);
        sum = left.sum + right.sum;
    }

    public int count() {
        return sum;
    }
}
```

```C++ [sol1-C++]
class CountIntervals {
    CountIntervals *left = nullptr, *right = nullptr;
    int l, r, sum = 0;

public:
    CountIntervals() : l(1), r(1e9) {}

    CountIntervals(int l, int r) : l(l), r(r) {}

    void add(int L, int R) { // 为方便区分变量名，将递归中始终不变的入参改为大写（视作常量）
        if (sum == r - l + 1) return; // 当前节点已被完整覆盖，无需执行任何操作
        if (L <= l && r <= R) { // 当前节点已被区间 [L,R] 完整覆盖，不再继续递归
            sum = r - l + 1;
            return;
        }
        int mid = (l + r) / 2;
        if (left == nullptr) left = new CountIntervals(l, mid); // 动态开点
        if (right == nullptr) right = new CountIntervals(mid + 1, r); // 动态开点
        if (L <= mid) left->add(L, R);
        if (mid < R) right->add(L, R);
        sum = left->sum + right->sum;
    }

    int count() { return sum; }
};
```

```go [sol1-Go]
type CountIntervals struct {
	left, right *CountIntervals
	l, r, sum   int
}

func Constructor() CountIntervals { return CountIntervals{l: 1, r: 1e9} }

func (o *CountIntervals) Add(l, r int) {
	if o.sum == o.r-o.l+1 { return } // o 已被完整覆盖，无需执行任何操作
	if l <= o.l && o.r <= r { // 当前节点已被区间 [l,r] 完整覆盖，不再继续递归
		o.sum = o.r - o.l + 1
		return
	}
	mid := (o.l + o.r) >> 1
	if o.left == nil { o.left = &CountIntervals{l: o.l, r: mid} } // 动态开点
	if o.right == nil { o.right = &CountIntervals{l: mid + 1, r: o.r} } // 动态开点
	if l <= mid { o.left.Add(l, r)}
	if mid < r { o.right.Add(l, r) }
	o.sum = o.left.sum + o.right.sum
}

func (o *CountIntervals) Count() int { return o.sum }
```
