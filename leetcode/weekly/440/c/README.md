具体请看 [视频讲解](https://www.bilibili.com/video/BV15gRaYZE5o/)，从线段树二分的角度，带你发明线段树。

用线段树维护 $\textit{baskets}$ 的区间最大值。

对于 $x=\textit{fruits}[i]$，在线段树上二分找第一个 $\ge x$ 的数。

- 如果整个区间的最大值都小于 $x$，那么没有这样的数，返回 $-1$。
- 如果能递归到叶子，返回叶子对应的区间端点。
- 先递归左子树。
- 如果左子树没找到，再递归右子树。

如果没有找到这样的数，把答案加一。

否则，把对应的位置改成 $-1$，表示不能放水果。

> 完整的线段树模板见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。

```py [sol-Python3]
class SegmentTree:
    def __init__(self, a: List[int]):
        n = len(a)
        self.max = [0] * (2 << (n - 1).bit_length())
        self.build(a, 1, 0, n - 1)

    def maintain(self, o: int):
        self.max[o] = max(self.max[o * 2], self.max[o * 2 + 1])

    # 初始化线段树
    def build(self, a: List[int], o: int, l: int, r: int):
        if l == r:
            self.max[o] = a[l]
            return
        m = (l + r) // 2
        self.build(a, o * 2, l, m)
        self.build(a, o * 2 + 1, m + 1, r)
        self.maintain(o)

    # 找区间内的第一个 >= x 的数，并更新为 -1，返回这个数的下标（没有则返回 -1）
    def find_first_and_update(self, o: int, l: int, r: int, x: int) -> int:
        if self.max[o] < x:  # 区间没有 >= x 的数
            return -1
        if l == r:
            self.max[o] = -1  # 更新为 -1，表示不能放水果
            return l
        m = (l + r) // 2
        i = self.find_first_and_update(o * 2, l, m, x)  # 先递归左子树
        if i < 0:  # 左子树没找到
            i = self.find_first_and_update(o * 2 + 1, m + 1, r, x)  # 再递归右子树
        self.maintain(o)
        return i


class Solution:
    def numOfUnplacedFruits(self, fruits: List[int], baskets: List[int]) -> int:
        t = SegmentTree(baskets)
        n = len(baskets)
        ans = 0
        for x in fruits:
            if t.find_first_and_update(1, 0, n - 1, x) < 0:
                ans += 1
        return ans
```

```java [sol-Java]
class SegmentTree {
    private final int[] max;

    public SegmentTree(int[] a) {
        int n = a.length;
        max = new int[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    // 找区间内的第一个 >= x 的数，并更新为 -1，返回这个数的下标（没有则返回 -1）
    public int findFirstAndUpdate(int o, int l, int r, int x) {
        if (max[o] < x) { // 区间没有 >= x 的数
            return -1;
        }
        if (l == r) {
            max[o] = -1; // 更新为 -1，表示不能放水果
            return l;
        }
        int m = (l + r) / 2;
        int i = findFirstAndUpdate(o * 2, l, m, x); // 先递归左子树
        if (i < 0) { // 左子树没找到
            i = findFirstAndUpdate(o * 2 + 1, m + 1, r, x); // 再递归右子树
        }
        maintain(o);
        return i;
    }

    private void maintain(int o) {
        max[o] = Math.max(max[o * 2], max[o * 2 + 1]);
    }

    // 初始化线段树
    private void build(int[] a, int o, int l, int r) {
        if (l == r) {
            max[o] = a[l];
            return;
        }
        int m = (l + r) / 2;
        build(a, o * 2, l, m);
        build(a, o * 2 + 1, m + 1, r);
        maintain(o);
    }
}

class Solution {
    public int numOfUnplacedFruits(int[] fruits, int[] baskets) {
        SegmentTree t = new SegmentTree(baskets);
        int n = baskets.length;
        int ans = 0;
        for (int x : fruits) {
            if (t.findFirstAndUpdate(1, 0, n - 1, x) < 0) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class SegmentTree {
    vector<int> mx;

    void maintain(int o) {
        mx[o] = max(mx[o * 2], mx[o * 2 + 1]);
    }

    // 初始化线段树
    void build(const vector<int>& a, int o, int l, int r) {
        if (l == r) {
            mx[o] = a[l];
            return;
        }
        int m = (l + r) / 2;
        build(a, o * 2, l, m);
        build(a, o * 2 + 1, m + 1, r);
        maintain(o);
    }

public:
    SegmentTree(const vector<int>& a) {
        size_t n = a.size();
        mx.resize(2 << bit_width(n - 1));
        build(a, 1, 0, n - 1);
    }

    // 找区间内的第一个 >= x 的数，并更新为 -1，返回这个数的下标（没有则返回 -1）
    int findFirstAndUpdate(int o, int l, int r, int x) {
        if (mx[o] < x) { // 区间没有 >= x 的数
            return -1;
        }
        if (l == r) {
            mx[o] = -1; // 更新为 -1，表示不能放水果
            return l;
        }
        int m = (l + r) / 2;
        int i = findFirstAndUpdate(o * 2, l, m, x); // 先递归左子树
        if (i < 0) { // 左子树没找到
            i = findFirstAndUpdate(o * 2 + 1, m + 1, r, x); // 再递归右子树
        }
        maintain(o);
        return i;
    }
};

class Solution {
public:
    int numOfUnplacedFruits(vector<int>& fruits, vector<int>& baskets) {
        SegmentTree t(baskets);
        int n = baskets.size(), ans = 0;
        for (int x : fruits) {
            if (t.findFirstAndUpdate(1, 0, n - 1, x) < 0) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type seg []int

func (t seg) maintain(o int) {
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 初始化线段树
func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 找区间内的第一个 >= x 的数，并更新为 -1，返回这个数的下标（没有则返回 -1）
func (t seg) findFirstAndUpdate(o, l, r, x int) int {
	if t[o] < x { // 区间没有 >= x 的数
		return -1
	}
	if l == r {
		t[o] = -1 // 更新为 -1，表示不能放水果
		return l
	}
	m := (l + r) >> 1
	i := t.findFirstAndUpdate(o<<1, l, m, x) // 先递归左子树
	if i < 0 { // 左子树没找到
		i = t.findFirstAndUpdate(o<<1|1, m+1, r, x) // 再递归右子树
	}
	t.maintain(o)
	return i
}

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func numOfUnplacedFruits(fruits, baskets []int) (ans int) {
	t := newSegmentTree(baskets)
	for _, x := range fruits {
		if t.findFirstAndUpdate(1, 0, len(baskets)-1, x) < 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{fruits}$ 的长度，也是 $\textit{baskets}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§8.3 线段树（无区间更新）**」，包含线段树的模板。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
