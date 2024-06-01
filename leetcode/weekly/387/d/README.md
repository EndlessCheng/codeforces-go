将元素**离散化**成 $[1,m]$ 中的元素，其中 $m$ 为 $\textit{nums}$ 中的不同元素个数。

这可以对 $\textit{nums}$ 排序去重后，在数组中二分查找得到。

记 $\textit{arr}_1$ 为 $a$，记 $\textit{arr}_2$ 为 $b$。用两棵**树状数组**分别维护 $a$ 和 $b$ 中的每个元素的出现次数，即可快速计算 $\texttt{greaterCount}$。然后按照题目要求模拟即可。

[带你发明树状数组！附数学证明](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)

[视频讲解](https://www.bilibili.com/video/BV14r421W7oR)

## 写法一：两棵树状数组

```py [sol-Python3]
class Fenwick:
    __slots__ = 'tree'

    def __init__(self, n: int):
        self.tree = [0] * n

    # 把下标为 i 的元素增加 1
    def add(self, i: int) -> None:
        while i < len(self.tree):
            self.tree[i] += 1
            i += i & -i

    # 返回下标在 [1,i] 的元素之和
    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.tree[i]
            i &= i - 1
        return res

class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        sorted_nums = sorted(set(nums))
        m = len(sorted_nums)
        a = [nums[0]]
        b = [nums[1]]
        t1 = Fenwick(m + 1)
        t2 = Fenwick(m + 1)
        t1.add(bisect_left(sorted_nums, nums[0]) + 1)
        t2.add(bisect_left(sorted_nums, nums[1]) + 1)
        for x in nums[2:]:
            v = bisect_left(sorted_nums, x) + 1
            gc1 = len(a) - t1.pre(v)  # greaterCount(a, v)
            gc2 = len(b) - t2.pre(v)  # greaterCount(b, v)
            if gc1 > gc2 or gc1 == gc2 and len(a) <= len(b):
                a.append(x)
                t1.add(v)
            else:
                b.append(x)
                t2.add(v)
        return a + b
```

```java [sol-Java]
class Fenwick {
    private final int[] tree;

    public Fenwick(int n) {
        tree = new int[n];
    }

    // 把下标为 i 的元素增加 1
    public void add(int i) {
        while (i < tree.length) {
            tree[i]++;
            i += i & -i;
        }
    }

    // 返回下标在 [1,i] 的元素之和
    public int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i &= i - 1;
        }
        return res;
    }
}

class Solution {
    public int[] resultArray(int[] nums) {
        int[] sorted = nums.clone();
        Arrays.sort(sorted); // 只排序不去重
        int n = nums.length;

        List<Integer> a = new ArrayList<>(n); // 预分配空间
        List<Integer> b = new ArrayList<>();
        a.add(nums[0]);
        b.add(nums[1]);

        Fenwick t1 = new Fenwick(n + 1);
        Fenwick t2 = new Fenwick(n + 1);
        t1.add(Arrays.binarySearch(sorted, nums[0]) + 1);
        t2.add(Arrays.binarySearch(sorted, nums[1]) + 1);

        for (int i = 2; i < nums.length; i++) {
            int x = nums[i];
            int v = Arrays.binarySearch(sorted, x) + 1;
            int gc1 = a.size() - t1.pre(v); // greaterCount(a, v)
            int gc2 = b.size() - t2.pre(v); // greaterCount(b, v)
            if (gc1 > gc2 || gc1 == gc2 && a.size() <= b.size()) {
                a.add(x);
                t1.add(v);
            } else {
                b.add(x);
                t2.add(v);
            }
        }
        a.addAll(b);
        for (int i = 0; i < n; i++) {
            nums[i] = a.get(i);
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Fenwick {
    vector<int> tree;

public:
    Fenwick(int n) : tree(n) {}

    // 把下标为 i 的元素增加 1
    void add(int i) {
        while (i < tree.size()) {
            tree[i]++;
            i += i & -i;
        }
    }

    // 返回下标在 [1,i] 的元素之和
    int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i &= i - 1;
        }
        return res;
    }
};

class Solution {
public:
    vector<int> resultArray(vector<int> &nums) {
        auto sorted = nums;
        ranges::sort(sorted);
        sorted.erase(unique(sorted.begin(), sorted.end()), sorted.end());
        int m = sorted.size();

        vector<int> a{nums[0]}, b{nums[1]};
        Fenwick t1(m + 1), t2(m + 1);
        t1.add(ranges::lower_bound(sorted, nums[0]) - sorted.begin() + 1);
        t2.add(ranges::lower_bound(sorted, nums[1]) - sorted.begin() + 1);
        for (int i = 2; i < nums.size(); i++) {
            int x = nums[i];
            int v = ranges::lower_bound(sorted, x) - sorted.begin() + 1;
            int gc1 = a.size() - t1.pre(v); // greaterCount(a, v)
            int gc2 = b.size() - t2.pre(v); // greaterCount(b, v)
            if (gc1 > gc2 || gc1 == gc2 && a.size() <= b.size()) {
                a.push_back(x);
                t1.add(v);
            } else {
                b.push_back(x);
                t2.add(v);
            }
        }
        a.insert(a.end(), b.begin(), b.end());
        return a;
    }
};
```

```go [sol-Go]
type fenwick []int

// 把下标为 i 的元素增加 1
func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

// 返回下标在 [1,i] 的元素之和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func resultArray(nums []int) (ans []int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	a := nums[:1]
	b := []int{nums[1]}
	t1 := make(fenwick, m+1)
	t2 := make(fenwick, m+1)
	t1.add(sort.SearchInts(sorted, nums[0]) + 1)
	t2.add(sort.SearchInts(sorted, nums[1]) + 1)
	for _, x := range nums[2:] {
		v := sort.SearchInts(sorted, x) + 1
		gc1 := len(a) - t1.pre(v) // greaterCount(a, v)
		gc2 := len(b) - t2.pre(v) // greaterCount(b, v)
		if gc1 > gc2 || gc1 == gc2 && len(a) <= len(b) {
			a = append(a, x)
			t1.add(v)
		} else {
			b = append(b, x)
			t2.add(v)
		}
	}
	return append(a, b...)
}
```

## 写法二：一棵树状数组

把元素 $v$ 添加到 $\textit{t}_2$ 的操作，可以改成把元素 $v$ 在 $\textit{t}_1$ 中的出现次数减一。

也就是说，用一棵树状数组维护 $a$ 和 $b$ 元素出现次数的**差值**。

同时，为了方便调用 $\texttt{pre}$，离散化的 $v$ 改成 $m-j$，其中 $j$ 是二分的下标。这样问题就转换成了求小于 $v$ 的元素个数（之差）。

```py [sol-Python3]
class Fenwick:
    __slots__ = 'tree'

    def __init__(self, n: int):
        self.tree = [0] * n

    # 把下标为 i 的元素增加 v
    def add(self, i: int, v: int) -> None:
        while i < len(self.tree):
            self.tree[i] += v
            i += i & -i

    # 返回下标在 [1,i] 的元素之和
    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.tree[i]
            i &= i - 1
        return res

class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        sorted_nums = sorted(set(nums))
        m = len(sorted_nums)
        a = [nums[0]]
        b = [nums[1]]
        t = Fenwick(m + 1)
        t.add(m - bisect_left(sorted_nums, nums[0]), 1)
        t.add(m - bisect_left(sorted_nums, nums[1]), -1)
        for x in nums[2:]:
            v = m - bisect_left(sorted_nums, x)
            d = t.pre(v - 1)  # 转换成 < v 的元素个数之差
            if d > 0 or d == 0 and len(a) <= len(b):
                a.append(x)
                t.add(v, 1)
            else:
                b.append(x)
                t.add(v, -1)
        return a + b
```

```java [sol-Java]
class Fenwick {
    private final int[] tree;

    public Fenwick(int n) {
        tree = new int[n];
    }

    // 把下标为 i 的元素增加 v
    public void add(int i, int v) {
        while (i < tree.length) {
            tree[i] += v;
            i += i & -i;
        }
    }

    // 返回下标在 [1,i] 的元素之和
    public int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i &= i - 1;
        }
        return res;
    }
}

class Solution {
    public int[] resultArray(int[] nums) {
        int[] sorted = nums.clone();
        Arrays.sort(sorted); // 只排序不去重
        int n = nums.length;

        List<Integer> a = new ArrayList<>(n); // 预分配空间
        List<Integer> b = new ArrayList<>();
        a.add(nums[0]);
        b.add(nums[1]);

        Fenwick t = new Fenwick(n + 1);
        t.add(n - Arrays.binarySearch(sorted, nums[0]), 1);
        t.add(n - Arrays.binarySearch(sorted, nums[1]), -1);

        for (int i = 2; i < nums.length; i++) {
            int x = nums[i];
            int v = n - Arrays.binarySearch(sorted, x);
            int d = t.pre(v - 1); // 转换成 < v 的元素个数之差
            if (d > 0 || d == 0 && a.size() <= b.size()) {
                a.add(x);
                t.add(v, 1);
            } else {
                b.add(x);
                t.add(v, -1);
            }
        }
        a.addAll(b);
        for (int i = 0; i < n; i++) {
            nums[i] = a.get(i);
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Fenwick {
    vector<int> tree;

public:
    Fenwick(int n) : tree(n) {}

    // 把下标为 i 的元素增加 v
    void add(int i, int v) {
        while (i < tree.size()) {
            tree[i] += v;
            i += i & -i;
        }
    }

    // 返回下标在 [1,i] 的元素之和
    int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i &= i - 1;
        }
        return res;
    }
};

class Solution {
public:
    vector<int> resultArray(vector<int> &nums) {
        auto sorted = nums;
        ranges::sort(sorted);
        sorted.erase(unique(sorted.begin(), sorted.end()), sorted.end());
        int m = sorted.size();

        vector<int> a{nums[0]}, b{nums[1]};
        Fenwick t(m + 1);
        t.add(sorted.end() - ranges::lower_bound(sorted, nums[0]), 1);
        t.add(sorted.end() - ranges::lower_bound(sorted, nums[1]), -1);
        for (int i = 2; i < nums.size(); i++) {
            int x = nums[i];
            int v = sorted.end() - ranges::lower_bound(sorted, x);
            int d = t.pre(v - 1); // 转换成 < v 的元素个数之差
            if (d > 0 || d == 0 && a.size() <= b.size()) {
                a.push_back(x);
                t.add(v, 1);
            } else {
                b.push_back(x);
                t.add(v, -1);
            }
        }
        a.insert(a.end(), b.begin(), b.end());
        return a;
    }
};
```

```go [sol-Go]
type fenwick []int

// 把下标为 i 的元素增加 v
func (f fenwick) add(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

// 返回下标在 [1,i] 的元素之和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func resultArray(nums []int) (ans []int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	a := nums[:1]
	b := []int{nums[1]}
	t := make(fenwick, m+1)
	t.add(m-sort.SearchInts(sorted, nums[0]), 1)
	t.add(m-sort.SearchInts(sorted, nums[1]), -1)
	for _, x := range nums[2:] {
		v := m - sort.SearchInts(sorted, x)
		d := t.pre(v - 1) // 转换成 < v 的元素个数之差
		if d > 0 || d == 0 && len(a) <= len(b) {
			a = append(a, x)
			t.add(v, 1)
		} else {
			b = append(b, x)
			t.add(v, -1)
		}
	}
	return append(a, b...)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 练习

见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/) 中的树状数组。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
