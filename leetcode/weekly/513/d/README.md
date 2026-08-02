$$
\dfrac{x}{y} \le \dfrac{a}{b}
$$

等价于

$$
ay-bx\ge 0
$$

当 $y=0$ 时，$x>0$（题目要求子数组非空），$ay-bx = -bx < 0$，上式不成立。所以题目的两个条件可以合并成上式。

把 $\textit{nums}$ 中的奇数视作 $a$，偶数视作 $-b$，得到数组 $\textit{arr}$。问题等价于：

- 计算 $\textit{arr}$ 中有多少个元素和 $\ge 0$ 的非空连续子数组。

设 $\textit{arr}$ 的**前缀和**数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

$\textit{arr}$ 的子数组 $[L,R-1]$ 的元素和等于 $s[R] - s[L]$。问题等价于：

- 有多少个下标对 $(L,R)$ 满足 $0\le L<R\le n$ 且 $s[R] - s[L]\ge 0$？

枚举 $R$，我们需要知道在 $R$ 的左边有多少个 $s[L] \le s[R]$。

类似逆序对，这可以用**值域树状数组**或**归并排序**计算。

[本题视频讲解](https://www.bilibili.com/video/BV1gn3R6qEbX/?t=8m53s)，包含两种方法，欢迎点赞关注~

## 方法一：值域树状数组

```py [sol-Python3]
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 1
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def add(self, i: int) -> None:
        t = self.tree
        while i < len(t):
            t[i] += 1
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        while i > 0:
            res += t[i]
            i &= i - 1
        return res


class Solution:
    def countRatioSubarrays(self, nums: list[int], a: int, b: int) -> int:
        s = [0] * (len(nums) + 1)
        for i, x in enumerate(nums):
            s[i + 1] = s[i] + (a if x % 2 else -b)  # 奇数视作 a，偶数视作 -b

        # s 去重排序
        sorted_s = sorted(set(s))

        t = FenwickTree(len(sorted_s) + 1)
        ans = 0
        for x in s:
            x = bisect_left(sorted_s, x) + 1  # 离散化（从 1 开始）
            ans += t.pre(x)  # 计算在 s 左边有多少个 <= s 的数
            t.add(x)
        return ans
```

```java [sol-Java]
class FenwickTree {
    private final int[] tree;

    public FenwickTree(int n) {
        tree = new int[n + 1]; // 使用下标 1 到 n
    }

    // a[i] 增加 1
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public void add(int i) {
        for (; i < tree.length; i += i & -i) {
            tree[i]++;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }
}

class Solution {
    public long countRatioSubarrays(int[] nums, int a, int b) {
        int n = nums.length;
        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums[i] % 2 == 0 ? -b : a); // 偶数视作 -b，奇数视作 a
        }

        // sum 复制排序
        long[] sortedS = sum.clone();
        Arrays.sort(sortedS);

        FenwickTree t = new FenwickTree(n + 1);
        long ans = 0;
        for (long s : sum) {
            int x = Arrays.binarySearch(sortedS, s) + 1; // 离散化（从 1 开始）
            ans += t.pre(x); // 计算在 s 左边有多少个 <= s 的数
            t.add(x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }
};

class Solution {
public:
    long long countRatioSubarrays(vector<int>& nums, int a, int b) {
        int n = nums.size();
        vector<long long> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums[i] & 1 ? a : -b); // 奇数视作 a，偶数视作 -b
        }

        // s 排序去重
        auto sorted_s = sum;
        ranges::sort(sorted_s);
        sorted_s.erase(ranges::unique(sorted_s).begin(), sorted_s.end());

        FenwickTree<int> t(sorted_s.size() + 1);
        long long ans = 0;
        for (auto s : sum) {
            int x = ranges::lower_bound(sorted_s, s) - sorted_s.begin() + 1; // 离散化（从 1 开始）
            ans += t.pre(x); // 计算在 s 左边有多少个 <= s 的数
            t.update(x, 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func (t fenwick) add(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func countRatioSubarrays(nums []int, a, b int) int64 {
	sum := make([]int, len(nums)+1)
	value := [2]int{-b, a}
	for i, x := range nums {
		sum[i+1] = sum[i] + value[x&1] // 偶数视作 -b，奇数视作 a
	}

	// sum 排序去重
	sorted := slices.Clone(sum)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	t := make(fenwick, len(sorted)+1)
	ans := 0
	for _, s := range sum {
		s = sort.SearchInts(sorted, s) + 1 // 离散化（从 1 开始）
		ans += t.pre(s) // 计算在 s 左边有多少个 <= s 的数
		t.add(s)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：归并排序

```py [sol-Python3]
class Solution:
    def countRatioSubarrays(self, nums: list[int], a: int, b: int) -> int:
        s = [0] * (len(nums) + 1)
        for i, x in enumerate(nums):
            s[i + 1] = s[i] + (a if x % 2 else -b)  # 奇数视作 a，偶数视作 -b

        def merge_count(s: list[int]) -> int:
            n = len(s)
            if n <= 1:
                return 0

            left = s[:n // 2]
            right = s[n // 2:]
            cnt = merge_count(left) + merge_count(right)  # left 和 right 各自的合法数对个数

            l = r = 0
            for i in range(n):
                # 计算一个在 left 中，另一个在 right 中的合法数对个数
                if l < len(left) and (r == len(right) or left[l] <= right[r]):
                    s[i] = left[l]
                    l += 1
                else:
                    cnt += l  # left[:l] 中的数都 <= right[r]，这有 l 个
                    s[i] = right[r]
                    r += 1

            return cnt

        return merge_count(s)
```

```java [sol-Java]
class Solution {
    public long countRatioSubarrays(int[] nums, int a, int b) {
        int n = nums.length;
        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums[i] % 2 == 0 ? -b : a); // 偶数视作 -b，奇数视作 a
        }

        return mergeCount(sum);
    }

    private long mergeCount(long[] sum) {
        int n = sum.length;
        if (n <= 1) {
            return 0;
        }

        long[] left = Arrays.copyOfRange(sum, 0, n / 2);
        long[] right = Arrays.copyOfRange(sum, n / 2, n);
        long cnt = mergeCount(left) + mergeCount(right); // left 和 right 各自的合法数对个数

        int l = 0;
        int r = 0;
        for (int i = 0; i < n; i++) {
            // 计算一个在 left 中，另一个在 right 中的合法数对个数
            if (l < left.length && (r == right.length || left[l] <= right[r])) {
                sum[i] = left[l];
                l++;
            } else {
                cnt += l; // left 的 [0,l-1] 中的数都 <= right[r]，这有 l 个
                sum[i] = right[r];
                r++;
            }
        }

        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    long long merge_count(vector<long long>& sum) {
        int n = sum.size();
        if (n <= 1) {
            return 0;
        }

        vector<long long> left(sum.begin(), sum.begin() + n / 2);
        vector<long long> right(sum.begin() + n / 2, sum.end());
        long long cnt = merge_count(left) + merge_count(right); // left 和 right 各自的合法数对个数

        int l = 0, r = 0;
        for (int i = 0; i < n; i++) {
            // 计算一个在 left 中，另一个在 right 中的合法数对个数
            if (l < left.size() && (r == right.size() || left[l] <= right[r])) {
                sum[i] = left[l];
                l++;
            } else {
                cnt += l; // left 的 [0,l-1] 中的数都 <= right[r]，这有 l 个
                sum[i] = right[r];
                r++;
            }
        }

        return cnt;
    }

public:
    long long countRatioSubarrays(vector<int>& nums, int a, int b) {
        int n = nums.size();
        vector<long long> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums[i] & 1 ? a : -b); // 奇数视作 a，偶数视作 -b
        }

        return merge_count(sum);
    }
};
```

```go [sol-Go]
func mergeCount(sum []int) int {
	n := len(sum)
	if n <= 1 {
		return 0
	}

	left := slices.Clone(sum[:n/2])
	right := slices.Clone(sum[n/2:])
	cnt := mergeCount(left) + mergeCount(right) // left 和 right 各自的合法数对个数

	l, r := 0, 0
	for i := range sum {
		// 计算一个在 left 中，另一个在 right 中的合法数对个数
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			sum[i] = left[l]
			l++
		} else {
			cnt += l // left[:l] 中的数都 <= right[r]，这有 l 个
			sum[i] = right[r]
			r++
		}
	}

	return cnt
}

func countRatioSubarrays(nums []int, a, b int) int64 {
	sum := make([]int, len(nums)+1)
	value := [2]int{-b, a}
	for i, x := range nums {
		sum[i+1] = sum[i] + value[x&1] // 偶数视作 -b，奇数视作 a
	}

	return int64(mergeCount(sum))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」和「**§8.2 逆序对**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
