为方便描述，把 $\textit{nums}$ 记作 $a$。

如果 $1\le i \le n-2$ 且 $a[i-1] < a[i] > a[i+1]$，那么把 $a[i]$ 视作 $1$，否则视作 $0$。

如此转换后，操作 1 相当于计算从 $l+1$ 到 $r-1$ 的子数组的元素和。请注意，题目说的是「子数组」的第一个和最后一个不是峰值元素，注意是子数组，不是整个数组。

由于操作 2 要动态修改数组，我们可以用**树状数组**维护，具体请看 [带你发明树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)。

具体来说：

1. 先把区间 $[\max(i-1,1),\min(i+1,n-2)]$ 中的峰值元素从树状数组中去掉。
2. 然后修改 $a[i]=\textit{val}$。
3. 最后再把区间 $[\max(i-1,1),\min(i+1,n-2)]$ 中的峰值元素加入树状数组。

[本题视频讲解](https://www.bilibili.com/video/BV1T1421k7Hi/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Fenwick:
    __slots__ = 'f'

    def __init__(self, n: int):
        self.f = [0] * n

    def update(self, i: int, val: int) -> None:
        while i < len(self.f):
            self.f[i] += val
            i += i & -i

    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.f[i]
            i &= i - 1
        return res

    def query(self, l: int, r: int) -> int:
        if r < l:
            return 0
        return self.pre(r) - self.pre(l - 1)

class Solution:
    def countOfPeaks(self, nums, queries):
        n = len(nums)
        f = Fenwick(n - 1)
        def update(i: int, val: int) -> None:
            if nums[i - 1] < nums[i] and nums[i] > nums[i + 1]:
                f.update(i, val)
        for i in range(1, n - 1):
            update(i, 1)

        ans = []
        for op, i, val in queries:
            if op == 1:
                ans.append(f.query(i + 1, val - 1))
                continue
            for j in range(max(i - 1, 1), min(i + 2, n - 1)):
                update(j, -1)
            nums[i] = val
            for j in range(max(i - 1, 1), min(i + 2, n - 1)):
                update(j, 1)
        return ans
```

```java [sol-Java]
class Fenwick {
    private final int[] f;

    Fenwick(int n) {
        f = new int[n];
    }

    void update(int i, int val) {
        for (; i < f.length; i += i & -i) {
            f[i] += val;
        }
    }

    private int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res += f[i];
        }
        return res;
    }

    int query(int l, int r) {
        if (r < l) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    public List<Integer> countOfPeaks(int[] nums, int[][] queries) {
        int n = nums.length;
        Fenwick f = new Fenwick(n - 1);
        for (int i = 1; i < n - 1; i++) {
            update(f, nums, i, 1);
        }

        List<Integer> ans = new ArrayList<>();
        for (int[] q : queries) {
            if (q[0] == 1) {
                ans.add(f.query(q[1] + 1, q[2] - 1));
                continue;
            }
            int i = q[1];
            for (int j = Math.max(i - 1, 1); j <= Math.min(i + 1, n - 2); j++) {
                update(f, nums, j, -1);
            }
            nums[i] = q[2];
            for (int j = Math.max(i - 1, 1); j <= Math.min(i + 1, n - 2); j++) {
                update(f, nums, j, 1);
            }
        }
        return ans;
    }

    private void update(Fenwick f, int[] nums, int i, int val) {
        if (nums[i - 1] < nums[i] && nums[i] > nums[i + 1]) {
            f.update(i, val);
        }
    }
}
```

```cpp [sol-C++]
class Fenwick {
    vector<int> f;

public:
    Fenwick(int n) : f(n) {}

    void update(int i, int val) {
        for (; i < f.size(); i += i & -i) {
            f[i] += val;
        }
    }

    int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res += f[i];
        }
        return res;
    }

    int query(int l, int r) {
        if (r < l) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
};

class Solution {
public:
    vector<int> countOfPeaks(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        Fenwick f(n - 1);
        auto update = [&](int i, int val) {
            if (nums[i - 1] < nums[i] && nums[i] > nums[i + 1]) {
                f.update(i, val);
            }
        };
        for (int i = 1; i < n - 1; i++) {
            update(i, 1);
        }

        vector<int> ans;
        for (auto& q : queries) {
            if (q[0] == 1) {
                ans.push_back(f.query(q[1] + 1, q[2] - 1));
                continue;
            }
            int i = q[1];
            for (int j = max(i - 1, 1); j <= min(i + 1, n - 2); j++) {
                update(j, -1);
            }
            nums[i] = q[2];
            for (int j = max(i - 1, 1); j <= min(i + 1, n - 2); j++) {
                update(j, 1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return res
}

func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func countOfPeaks(nums []int, queries [][]int) (ans []int) {
	n := len(nums)
	f := make(fenwick, n-1)
	update := func(i, val int) {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			f.update(i, val)
		}
	}
	for i := 1; i < n-1; i++ {
		update(i, 1)
	}

	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, f.query(q[1]+1, q[2]-1))
			continue
		}
		i := q[1]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			update(j, -1)
		}
		nums[i] = q[2]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			update(j, 1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

**注**：可以用 [带你发明树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/) 中的技巧，$\mathcal{O}(n)$ 初始化树状数组，从而做到 $\mathcal{O}(n + q\log n)$ 的时间复杂度。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
