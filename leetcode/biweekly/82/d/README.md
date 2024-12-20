本题 [视频讲解](https://www.bilibili.com/video/BV1Le4y1R7xu) 已出炉，欢迎点赞三连~

---

## 方法一：并查集

#### 提示 1

数组中的元素越大越好，不妨从大往小考虑 $\textit{nums}[i]$。

#### 提示 2

子数组的长度 $k$ 越大，$\dfrac{\textit{threshold}}{k}$ 就越小，越能满足要求。

#### 提示 3

把考虑过的元素都串起来，这条链的长度就是 $k$。

于是关键在于如何动态维护每条链的长度，高效地串联两条链。

#### 提示 4

用并查集，遍历到 $\textit{nums}[i]$ 时，用并查集合并 $i$ 和 $i+1$，这样可以把连续访问过的位置串起来，同时维护链的长度。

```py [sol1-Python3]
class Solution:
    def validSubarraySize(self, nums: List[int], threshold: int) -> int:
        n = len(nums)
        fa = list(range(n + 1))
        sz = [0] * (n + 1)
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]
        for num, i in sorted(zip(nums, range(n)), reverse=True):
            j = find(i + 1)
            fa[i] = j  # 合并 i 和 i+1
            sz[j] += sz[i] + 1
            if num > threshold // sz[j]: return sz[j]
        return -1
```

```java [sol1-Java]
class Solution {
    int[] fa;

    public int validSubarraySize(int[] nums, int threshold) {
        var n = nums.length;
        fa = new int[n + 1];
        for (var i = 0; i <= n; i++) fa[i] = i;
        var sz = new int[n + 1];

        var ids = IntStream.range(0, n).boxed().toArray(Integer[]::new);
        Arrays.sort(ids, (i, j) -> nums[j] - nums[i]);
        for (var i : ids) {
            var j = find(i + 1);
            fa[i] = j; // 合并 i 和 i+1
            sz[j] += sz[i] + 1;
            if (nums[i] > threshold / sz[j]) return sz[j];
        }
        return -1;
    }

    int find(int x) {
        if (fa[x] != x) fa[x] = find(fa[x]);
        return fa[x];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int validSubarraySize(vector<int> &nums, int threshold) {
        int n = nums.size();
        int fa[n + 1], sz[n + 1];
        iota(fa, fa + n + 1, 0);
        memset(sz, 0, sizeof(sz));
        function<int(int)> find = [&](int x) -> int { return fa[x] == x ? x : fa[x] = find(fa[x]); };

        int ids[n];
        iota(ids, ids + n, 0);
        sort(ids, ids + n, [&](int i, int j) { return nums[i] > nums[j]; });
        for (int i : ids) {
            int j = find(i + 1);
            fa[i] = j; // 合并 i 和 i+1
            sz[j] += sz[i] + 1;
            if (nums[i] > threshold / sz[j]) return sz[j];
        }
        return -1;
    }
};
```

```go [sol1-Go]
func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i, v := range nums {
		a[i] = pair{v, i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	sz := make([]int, n+1)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for _, p := range a {
		i := p.i
		j := find(i + 1)
		fa[i] = j // 合并 i 和 i+1
		sz[j] += sz[i] + 1
		if p.v > threshold/sz[j] {
			return sz[j]
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：单调栈

#### 提示 1

枚举每个元素，假设它是子数组中的最小值。

#### 提示 2

子数组的左右边界最远能到哪？

#### 提示 3

用**单调栈**来计算左右边界。

不了解单调栈的同学可以看一下 [496. 下一个更大元素 I](https://leetcode-cn.com/problems/next-greater-element-i/)。本题求的是更小元素。

知道了左右边界也就知道了子数组的长度 $k$。

```py [sol1-Python3]
class Solution:
    def validSubarraySize(self, nums: List[int], threshold: int) -> int:
        n = len(nums)
        left, st = [-1] * n, []  # left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
        for i, v in enumerate(nums):
            while st and nums[st[-1]] >= v: st.pop()
            if st: left[i] = st[-1]
            st.append(i)

        right, st = [n] * n, []  # right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
        for i in range(n - 1, -1, -1):
            while st and nums[st[-1]] >= nums[i]: st.pop()
            if st: right[i] = st[-1]
            st.append(i)

        for num, l, r in zip(nums, left, right):
            k = r - l - 1
            if num > threshold // k: return k
        return -1
```

```java [sol1-Java]
class Solution {
    public int validSubarraySize(int[] nums, int threshold) {
        var n = nums.length;
        var left = new int[n]; // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
        var st = new ArrayDeque<Integer>();
        for (var i = 0; i < n; i++) {
            while (!st.isEmpty() && nums[st.peek()] >= nums[i]) st.pop();
            left[i] = st.isEmpty() ? -1 : st.peek();
            st.push(i);
        }

        var right = new int[n]; // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
        st = new ArrayDeque<>();
        for (var i = n - 1; i >= 0; i--) {
            while (!st.isEmpty() && nums[st.peek()] >= nums[i]) st.pop();
            right[i] = st.isEmpty() ? n : st.peek();
            st.push(i);
        }

        for (var i = 0; i < n; ++i) {
            var k = right[i] - left[i] - 1;
            if (nums[i] > threshold / k) return k;
        }
        return -1;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int validSubarraySize(vector<int> &nums, int threshold) {
        int n = nums.size();
        int left[n]; // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
        stack<int> s;
        for (int i = 0; i < n; ++i) {
            while (!s.empty() && nums[s.top()] >= nums[i]) s.pop();
            left[i] = s.empty() ? -1 : s.top();
            s.push(i);
        }

        int right[n]; // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
        s = stack<int>();
        for (int i = n - 1; i >= 0; --i) {
            while (!s.empty() && nums[s.top()] >= nums[i]) s.pop();
            right[i] = s.empty() ? n : s.top();
            s.push(i);
        }

        for (int i = 0; i < n; ++i) {
            int k = right[i] - left[i] - 1;
            if (nums[i] > threshold / k) return k;
        }
        return -1;
    }
};
```

```go [sol1-Go]
func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	left := make([]int, n) // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	right := make([]int, n) // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && nums[st[len(st)-1]] >= nums[i] {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	for i, num := range nums {
		k := right[i] - left[i] - 1
		if num > threshold/k {
			return k
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
