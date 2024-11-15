**请记住：有序是一个非常好的性质。**

把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 组合起来，按照 $\textit{nums}_2[i]$ 从大到小排序。枚举 $\textit{nums}_2[i]$ 作为序列的最小值，那么 $\textit{nums}_1$ 就只能在下标 $\le i$ 的数中选了。要选最大的 $k$ 个数。

根据 [703. 数据流中的第 K 大元素](https://leetcode.cn/problems/kth-largest-element-in-a-stream/)，这可以用一个大小固定为 $k$ 的最小堆来做，如果当前元素大于堆顶，就替换堆顶，这样可以让堆中元素之和变大。

[视频讲解](https://www.bilibili.com/video/BV1jG4y197qD/) 第三题。

```py [sol-Python3]
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        a = sorted(zip(nums1, nums2), key=lambda p: -p[1])
        h = [x for x, _ in a[:k]]
        heapify(h)
        s = sum(h)
        ans = s * a[k - 1][1]
        for x, y in a[k:]:
            if x > h[0]:
                s += x - heapreplace(h, x)
                ans = max(ans, s * y)
        return ans
```

```java [sol-Java]
class Solution {
    public long maxScore(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        Integer[] ids = new Integer[n];
        for (int i = 0; i < n; i++) {
            ids[i] = i;
        }
        // 对下标排序，不影响原数组的顺序
        Arrays.sort(ids, (i, j) -> nums2[j] - nums2[i]);

        PriorityQueue<Integer> pq = new PriorityQueue<>();
        long sum = 0;
        for (int i = 0; i < k; i++) {
            sum += nums1[ids[i]];
            pq.offer(nums1[ids[i]]);
        }

        long ans = sum * nums2[ids[k - 1]];
        for (int i = k; i < n; i++) {
            int x = nums1[ids[i]];
            if (x > pq.peek()) {
                sum += x - pq.poll();
                pq.offer(x);
                ans = Math.max(ans, sum * nums2[ids[i]]);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxScore(vector<int> &nums1, vector<int> &nums2, int k) {
        int n = nums1.size();
        vector<int> ids(n);
        iota(ids.begin(), ids.end(), 0);
        // 对下标排序，不影响原数组的顺序
        ranges::sort(ids, [&](int i, int j) { return nums2[i] > nums2[j]; });

        priority_queue<int, vector<int>, greater<>> pq;
        long long sum = 0;
        for (int i = 0; i < k; i++) {
            sum += nums1[ids[i]];
            pq.push(nums1[ids[i]]);
        }

        long long ans = sum * nums2[ids[k - 1]];
        for (int i = k; i < n; i++) {
            int x = nums1[ids[i]];
            if (x > pq.top()) {
                sum += x - pq.top();
                pq.pop();
                pq.push(x);
                ans = max(ans, sum * nums2[ids[i]]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxScore(nums1, nums2 []int, k int) int64 {
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}
	// 对下标排序，不影响原数组的顺序
	slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] })

	h := hp{make([]int, k)}
	sum := 0
	for i, idx := range ids[:k] {
		sum += nums1[idx]
		h.IntSlice[i] = nums1[idx]
	}
	heap.Init(&h)

	ans := sum * nums2[ids[k-1]]
	for _, i := range ids[k:] {
		x := nums1[i]
		if x > h.IntSlice[0] {
			sum += x - h.replace(x)
			ans = max(ans, sum*nums2[i])
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)            {}
func (hp) Pop() (_ any)        { return }
func (h hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
```

```js [sol-JavaScript]
var maxScore = function(nums1, nums2, k) {
    const n = nums1.length;
    const ids = [...Array(n).keys()];
    // 对下标排序，不影响原数组的顺序
    ids.sort((i, j) => nums2[j] - nums2[i]);

    const pq = new MinPriorityQueue();
    let sum = 0;
    for (let i = 0; i < k; i++) {
        sum += nums1[ids[i]];
        pq.enqueue(nums1[ids[i]]);
    }

    let ans = sum * nums2[ids[k - 1]];
    for (let i = k; i < n; i++) {
        const x = nums1[ids[i]];
        if (x > pq.front().element) {
            sum += x - pq.dequeue().element;
            pq.enqueue(x);
            ans = Math.max(ans, sum * nums2[ids[i]]);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::BinaryHeap;

impl Solution {
    pub fn max_score(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i64 {
        let n = nums1.len();
        let k = k as usize;
        let mut ids = (0..n).collect::<Vec<_>>();
        // 对下标排序，不影响原数组的顺序
        ids.sort_unstable_by(|&i, &j| nums2[j].cmp(&nums2[i]));

        let mut h = BinaryHeap::new();
        let mut sum = 0;
        for i in 0..k {
            sum += nums1[ids[i]] as i64;
            h.push(-nums1[ids[i]]); // 加负号变成最小堆
        }

        let mut ans = sum * nums2[ids[k - 1]] as i64;
        for i in k..n {
            let x = nums1[ids[i]];
            if x > -h.peek().unwrap() {
                sum += (x + h.pop().unwrap()) as i64;
                h.push(-x);
                ans = ans.max(sum * nums2[ids[i]] as i64);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
