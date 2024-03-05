文字讲解：[带你发明树状数组！附数学证明](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)

视频讲解：[周赛 387](https://www.bilibili.com/video/BV14r421W7oR/) 第四题。

将元素**离散化**成 $[1,m]$ 中的元素，其中 $m$ 为 $\textit{nums}$ 中的不同元素个数。

这可以对 $\textit{nums}$ 排序去重后，在数组中二分查找得到。

记 $\textit{arr}_1$ 为 $a$，记 $\textit{arr}_2$ 为 $b$。用两棵树状数组分别维护 $a$ 和 $b$ 中的每个元素的出现次数，即可快速计算 $\texttt{greaterCount}$。然后按照题目要求模拟即可。

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

- [307. 区域和检索 - 数组可修改](https://leetcode.cn/problems/range-sum-query-mutable/) *模板题
- [315. 计算右侧小于当前元素的个数](https://leetcode.cn/problems/count-of-smaller-numbers-after-self/) *逆序对
- [2426. 满足不等式的数对数目](https://leetcode.cn/problems/number-of-pairs-satisfying-inequality/) 2030
- [493. 翻转对](https://leetcode.cn/problems/reverse-pairs/)
- [327. 区间和的个数](https://leetcode.cn/problems/count-of-range-sum/)

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
