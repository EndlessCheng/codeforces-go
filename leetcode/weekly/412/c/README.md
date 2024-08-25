**核心观察**：对于两个数 $x$ 和 $y$，如果 $x<y$ 且 $x\cdot \textit{multiplier} \ge y$，那么 $x\cdot \textit{multiplier} < y\cdot \textit{multiplier}$。

这意味着当两个数接近时，我们会**交替操作**这两个数，而**不会连续操作同一个数**。

对于更多的数的情况也同理，当这些数接近时，我们会按照从小到大的顺序依次操作这些数。

那么，首先用最小堆手动模拟操作，直到原数组的最大值 $\textit{mx}$ 成为这 $n$ 个数的最小值。根据上面的结论，后面的操作就不需要手动模拟了。

假设此时还剩下 $k$ 次操作，那么：

- 对于前 $k\bmod n$ 小的数，还可以再操作 $\left\lfloor\dfrac{k}{n}\right\rfloor+1$ 次。
- 其余元素，还可以再操作 $\left\lfloor\dfrac{k}{n}\right\rfloor$ 次。

用**快速幂**计算操作这么多次后的结果，原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cMW6ePEwC/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        if multiplier == 1:  # 数组不变
            return nums

        MOD = 1_000_000_007
        n = len(nums)
        mx = max(nums)
        h = [(x, i) for i, x in enumerate(nums)]
        heapify(h)

        # 模拟，直到堆顶是 mx
        while k and h[0][0] < mx:
            x, i = h[0]
            heapreplace(h, (x * multiplier, i))
            k -= 1

        # 剩余的操作可以直接用公式计算
        h.sort()
        for i, (x, j) in enumerate(h):
            nums[j] = x * pow(multiplier, k // n + (i < k % n), MOD) % MOD
        return nums
```

```py [sol-Python3 写法二]
# 也可以模拟到 k 刚好是 n 的倍数时才停止，这样最后无需排序
class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        if multiplier == 1:  # 数组不变
            return nums

        MOD = 1_000_000_007
        n = len(nums)
        mx = max(nums)
        h = [(x, i) for i, x in enumerate(nums)]
        heapify(h)

        # 模拟，直到堆顶是 mx
        while k and (h[0][0] < mx or k % n):
            x, i = h[0]
            heapreplace(h, (x * multiplier, i))
            k -= 1

        # 剩余的操作可以直接用公式计算
        for x, j in h:
            nums[j] = x * pow(multiplier, k // n, MOD) % MOD
        return nums
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int[] getFinalState(int[] nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return nums;
        }

        int n = nums.length;
        int mx = 0;
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> a[0] != b[0] ? Long.compare(a[0], b[0]) : Long.compare(a[1], b[1]));
        for (int i = 0; i < n; i++) {
            mx = Math.max(mx, nums[i]);
            pq.offer(new long[]{nums[i], i});
        }

        // 模拟，直到堆顶是 mx
        for (; k > 0 && pq.peek()[0] < mx; k--) {
            long[] p = pq.poll();
            p[0] *= multiplier;
            pq.offer(p);
        }

        // 剩余的操作可以直接用公式计算
        for (int i = 0; i < n; i++) {
            long[] p = pq.poll();
            nums[(int) p[1]] = (int) (p[0] % MOD * pow(multiplier, k / n + (i < k % n ? 1 : 0)) % MOD);
        }
        return nums;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    vector<int> getFinalState(vector<int>& nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return nums;
        }

        int n = nums.size();
        int mx = ranges::max(nums);
        vector<pair<long long, int>> h(n);
        for (int i = 0; i < n; i++) {
            h[i] = {nums[i], i};
        }
        ranges::make_heap(h, greater<>()); // 最小堆，O(n) 堆化

        // 模拟，直到堆顶是 mx
        for (; k && h[0].first < mx; k--) {
            ranges::pop_heap(h, greater<>());
            h.back().first *= multiplier;
            ranges::push_heap(h, greater<>());
        }

        // 剩余的操作可以直接用公式计算
        ranges::sort(h); // 可以换成快速选择
        for (int i = 0; i < n; i++) {
            auto& [x, j] = h[i];
            nums[j] = x % MOD * pow(multiplier, k / n + (i < k % n)) % MOD;
        }
        return nums;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	heap.Init(&h)

	// 模拟，直到堆顶是 mx
	for ; k > 0 && h[0].x < mx; k-- {
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}

	// 剩余的操作可以直接用公式计算
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.i] = p.x % mod * pow(multiplier, e) % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。瓶颈在模拟那，每个数至多操作 $\mathcal{O}(\log U)$ 次。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**堆**」。

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
