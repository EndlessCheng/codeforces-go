**前置题目**：[703. 数据流中的第 K 大元素](https://leetcode.cn/problems/kth-largest-element-in-a-stream/)

根据输入，创建 $n$ 个三元组 $(\textit{nums}_1[i], \textit{nums}_2[i], i)$，然后按照 $\textit{nums}_1[i]$ 从小到大排序。

排序后，小于 $\textit{nums}_1[i]$ 的数，都在 $\textit{nums}_1[i]$ 左边，这样方便我们增量地处理。

遍历三元组列表，同时用一个**最小堆**维护 $\textit{nums}_2[i]$ 的前 $k$ 大元素：

- 把 $\textit{nums}_2[i]$ 入堆。
- 如果堆的大小超过 $k$，弹出堆顶（这 $k+1$ 个数的最小值）。
- 入堆出堆的过程中，用一个变量 $s$ 维护堆中元素之和。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15gRaYZE5o/?t=32m17s)，欢迎点赞关注~

## 写法一（更通用）

可能存在多个 $\textit{nums}_1[i]$ 相同的情况，要把这些相同的数一起处理，原理见 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)。

```py [sol-Python3]
class Solution:
    def findMaxSum(self, nums1: List[int], nums2: List[int], k: int) -> List[int]:
        a = sorted((x, y, i) for i, (x, y) in enumerate(zip(nums1, nums2)))

        n = len(a)
        ans = [0] * n
        h = []
        # 分组循环模板
        s = i = 0
        while i < n:
            start = i
            x = a[start][0]
            # 找到所有相同的 nums1[i]，这些数的答案都是一样的
            while i < n and a[i][0] == x:
                ans[a[i][2]] = s
                i += 1
            # 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
            for j in range(start, i):
                y = a[j][1]
                s += y
                heappush(h, y)
                if len(h) > k:
                    s -= heappop(h)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] findMaxSum(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        int[][] a = new int[n][3];
        for (int i = 0; i < n; i++) {
            a[i] = new int[]{nums1[i], nums2[i], i};
        }
        Arrays.sort(a, (p, q) -> p[0] - q[0]);

        long[] ans = new long[n];
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        long s = 0;
        // 分组循环模板
        for (int i = 0; i < n; ) {
            int start = i;
            int x = a[start][0];
            // 找到所有相同的 nums1[i]，这些数的答案都是一样的
            while (i < n && a[i][0] == x) {
                ans[a[i][2]] = s;
                i++;
            }
            // 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
            for (int j = start; j < i; j++) {
                int y = a[j][1];
                s += y;
                pq.offer(y);
                if (pq.size() > k) {
                    s -= pq.poll();
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> findMaxSum(vector<int>& nums1, vector<int>& nums2, int k) {
        int n = nums1.size();
        vector<tuple<int, int, int>> a(n);
        for (int i = 0; i < n; i++) {
            a[i] = {nums1[i], nums2[i], i};
        }
        ranges::sort(a);

        vector<long long> ans(n);
        priority_queue<int, vector<int>, greater<>> pq;
        long long s = 0;
        // 分组循环模板
        for (int i = 0; i < n;) {
            int start = i;
            int x = get<0>(a[start]);
            // 找到所有相同的 nums1[i]，这些数的答案都是一样的
            while (i < n && get<0>(a[i]) == x) {
                ans[get<2>(a[i])] = s;
                i++;
            }
            // 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
            for (int j = start; j < i; j++) {
                int y = get<1>(a[j]);
                s += y;
                pq.push(y);
                if (pq.size() > k) {
                    s -= pq.top();
                    pq.pop();
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := &hp{}
	s := 0
	// 分组循环模板
	for i := 0; i < n; {
		start := i
		// 找到所有相同的 nums1[i]，这些数的答案都是一样的
		x := a[start].x
		for ; i < n && a[i].x == x; i++ {
			ans[a[i].i] = int64(s)
		}
		// 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
		for ; start < i; start++ {
			y := a[start].y
			s += y
			heap.Push(h, y)
			if h.Len() > k {
				s -= heap.Pop(h).(int)
			}
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

## 写法二（针对本题）

上面是通用写法。针对本题，可以判断 $\textit{nums}_1[i]$ 和 $\textit{nums}_1[i-1]$ 是否相等，如果相等那么 $\textit{nums}_1[i]$ 的答案就是 $\textit{nums}_1[i-1]$ 的答案。

```py [sol-Python3]
class Solution:
    def findMaxSum(self, nums1: List[int], nums2: List[int], k: int) -> List[int]:
        a = sorted((x, y, i) for i, (x, y) in enumerate(zip(nums1, nums2)))
        n = len(a)
        ans = [0] * n
        h = []
        s = 0
        for i, (x, y, idx) in enumerate(a):
            ans[idx] = ans[a[i - 1][2]] if i and x == a[i - 1][0] else s
            s += y
            heappush(h, y)
            if len(h) > k:
                s -= heappop(h)
        return ans
```

```py [sol-Python3 更快写法]
class Solution:
    def findMaxSum(self, nums1: List[int], nums2: List[int], k: int) -> List[int]:
        a = sorted((x, y, i) for i, (x, y) in enumerate(zip(nums1, nums2)))
        n = len(a)
        ans = [0] * n
        h = []
        s = 0
        for i, (x, y, idx) in enumerate(a):
            ans[idx] = ans[a[i - 1][2]] if i and x == a[i - 1][0] else s
            s += y
            if i < k:
                heappush(h, y)
            else:
                s -= heappushpop(h, y)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] findMaxSum(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        int[][] a = new int[n][3];
        for (int i = 0; i < n; i++) {
            a[i] = new int[]{nums1[i], nums2[i], i};
        }
        Arrays.sort(a, (p, q) -> p[0] - q[0]);

        long[] ans = new long[n];
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        long s = 0;
        for (int i = 0; i < n; i++) {
            ans[a[i][2]] = i > 0 && a[i][0] == a[i - 1][0] ? ans[a[i - 1][2]] : s;
            int y = a[i][1];
            s += y;
            pq.offer(y);
            if (pq.size() > k) {
                s -= pq.poll();
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> findMaxSum(vector<int>& nums1, vector<int>& nums2, int k) {
        int n = nums1.size();
        vector<tuple<int, int, int>> a(n);
        for (int i = 0; i < n; i++) {
            a[i] = {nums1[i], nums2[i], i};
        }
        ranges::sort(a);

        vector<long long> ans(n);
        priority_queue<int, vector<int>, greater<>> pq;
        long long s = 0;
        for (int i = 0; i < n; i++) {
            auto& [x, y, idx] = a[i];
            ans[idx] = i > 0 && x == get<0>(a[i - 1]) ? ans[get<2>(a[i - 1])] : s;
            s += y;
            pq.push(y);
            if (pq.size() > k) {
                s -= pq.top();
                pq.pop();
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := &hp{}
	s := 0
	for i, t := range a {
		if i > 0 && t.x == a[i-1].x {
			ans[t.i] = ans[a[i-1].i]
		} else {
			ans[t.i] = int64(s)
		}
		s += t.y
		heap.Push(h, t.y)
		if h.Len() > k {
			s -= heap.Pop(h).(int)
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

```go [sol-Go 更快写法]
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := hp{make([]int, k)}
	s := 0
	for i, t := range a {
		if i > 0 && t.x == a[i-1].x {
			ans[t.i] = ans[a[i-1].i]
		} else {
			ans[t.i] = int64(s)
		}
		y := t.y
		if i < k {
			s += y
			h.IntSlice[i] = y
			continue
		}
		if i == k {
			heap.Init(&h)
		}
		if y > h.IntSlice[0] {
			s += y - h.IntSlice[0]
			h.IntSlice[0] = y
			heap.Fix(&h, 0)
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)     {}
func (hp) Pop() (_ any) { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**§5.1 堆基础**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
