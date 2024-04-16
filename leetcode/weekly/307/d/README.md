## 转换

计算 $\textit{nums}$ 中所有**非负数**的和，记作 $\textit{sum}$。

$\textit{nums}$ 的任意一个子序列的元素和，都等价于从 $\textit{sum}$ 中减去某些**非负数** / 加上某些**负数**得到。

例如 $\textit{nums}=[1,2,3,-4]$，其非负数的和为 $1+2+3=6$，我们可以从 $6$ 中减去 $2$ 得到 $\textit{nums}$ 的子序列 $[1,3]$ 的和 $1+3=4$，也可以把 $6$ 和 $-4$ 相加，得到 $\textit{nums}$ 的子序列 $[1,2,3,-4]$ 的和 $2$。

注意到，「减去非负数」和「加上负数」都相当于减去 $|\textit{nums}[i]|$。

$\textit{sum}$ 减去的数越小，$\textit{nums}$ 的子序列和就越大。

现在要解决的问题是：

- 把每个 $\textit{nums}[i]$ **取绝对值**后，$\textit{nums}$ 的第 $k$ 小的子序列和是多少？

## 方法一：二分答案 + 爆搜

#### 前置知识

1. [二分原理](https://www.bilibili.com/video/BV1AP41137w7/)
2. [子集型回溯](https://www.bilibili.com/video/BV1mG4y1A7Gu/)

二分答案，设当前二分的值为 $\textit{sumLimit}$。

问题变成：判断是否有至少 $k$ 个子序列，其元素和 $s$ 不超过 $\textit{sumLimit}$？

> 注：一道题能否二分答案，得看它有没有单调性。对于本题，$\textit{sumLimit}$ 越大，这样的子序列越多，有单调性，可以二分答案。

爆搜，从小到大考虑每个 $\textit{nums}[i]$ **选或不选**。在递归中，如果发现 $\textit{cnt}=k$ 或者 $s+\textit{nums}[i]>\textit{sumLimit}$，就不再继续递归，因为前者说明我们已经找到 $k$ 个和不超过 $\textit{sumLimit}$ 的子序列，后者说明子序列的和太大。由于每个 $\textit{nums}[i]$ 都取了绝对值，没有负数，$s$ 不会减小，所以可以**剪枝**。

二分下界：$0$。

二分上界：$\sum\limits_{i=0}^{n-1}|\textit{nums}[i]|$，即 $\textit{nums}$ 的所有元素的绝对值的和。

最后，用 $\textit{sum}$ 减去二分得到的值，即为答案。

#### 答疑

**问**：有没有可能，二分得到的值，并不是 $\textit{nums}$ 的子序列和？比如 $\textit{nums}[i]$ 都是偶数，但二分得到的却是一个奇数。

**答**：设二分得到的值为 $x$，那么 $x$ 一定是 $\textit{nums}$ 的子序列和。使用**反证法**证明：

假设 $x$ 不是 $\textit{nums}$ 的子序列和，也就是没有任何子序列的和等于 $x$，这意味着 $s \le x$ 等价于 $s\le x-1$，我们能从 $\textit{nums}$ 中找到 $k$ 个元素和不超过 $x-1$ 的子序列，所以 $\texttt{check}(x-1) = \texttt{true}$。但二分循环结束时，有 $\texttt{check}(x-1) = \texttt{false}$，矛盾，所以原命题成立，$x$ 一定是 $\textit{nums}$ 的子序列和。

```py [sol-Python3]
class Solution:
    def kSum(self, nums: List[int], k: int) -> int:
        s = 0
        for i, x in enumerate(nums):
            if x >= 0:
                s += x
            else:
                nums[i] = -x
        nums.sort()

        def check(sum_limit: int) -> bool:
            cnt = 1  # 空子序列算一个
            def dfs(i: int, s: int) -> None:
                nonlocal cnt
                if cnt == k or i == len(nums) or s + nums[i] > sum_limit:
                    return
                cnt += 1  # s + nums[i] <= sum_limit
                dfs(i + 1, s + nums[i])  # 选
                dfs(i + 1, s)  # 不选
            dfs(0, 0)
            return cnt == k  # 找到 k 个元素和不超过 sum_limit 的子序列
        return s - bisect_left(range(sum(nums)), True, key=check)
```

```java [sol-Java]
class Solution {
    public long kSum(int[] nums, int k) {
        long sum = 0, right = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] >= 0) {
                sum += nums[i];
            } else {
                nums[i] = -nums[i];
            }
            right += nums[i];
        }
        Arrays.sort(nums);

        long left = -1;
        while (left + 1 < right) { // 开区间二分，原理见【前置知识】
            long mid = (left + right) / 2;
            cnt = k - 1; // 空子序列算一个
            dfs(0, mid, nums);
            if (cnt == 0) { // 找到 k 个元素和不超过 mid 的子序列
                right = mid;
            } else {
                left = mid;
            }
        }
        return sum - right;
    }

    private int cnt;

    // 反向递归，增加改成减少，这样可以少传一些参数
    private void dfs(int i, long s, int[] nums) {
        if (cnt == 0 || i == nums.length || s < nums[i]) {
            return;
        }
        cnt--;
        dfs(i + 1, s - nums[i], nums); // 选
        dfs(i + 1, s, nums); // 不选
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long kSum(vector<int> &nums, int k) {
        long sum = 0;
        for (int &x : nums) {
            if (x >= 0) {
                sum += x;
            } else {
                x = -x;
            }
        }
        ranges::sort(nums);

        auto check = [&](long sum_limit) -> bool {
            int cnt = 1; // 空子序列算一个
            function<void(int, long long)> dfs = [&](int i, long long s) {
                if (cnt == k || i == nums.size() || s + nums[i] > sum_limit) {
                    return;
                }
                cnt++; // s + nums[i] <= sum_limit
                dfs(i + 1, s + nums[i]); // 选
                dfs(i + 1, s); // 不选
            };
            dfs(0, 0);
            return cnt == k; // 找到 k 个元素和不超过 sum_limit 的子序列
        };

        long long left = -1, right = accumulate(nums.begin(), nums.end(), 0LL);
        while (left + 1 < right) { // 开区间二分，原理见【前置知识】
            long long mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return sum - right;
    }
};
```

```go [sol-Go]
func kSum(nums []int, k int) int64 {
	sum, total := 0, 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
			total += x
		} else {
			total -= x
			nums[i] = -x
		}
	}
	slices.Sort(nums)

	kthS := sort.Search(total, func(sumLimit int) bool {
		cnt := 1 // 空子序列算一个
		var dfs func(int, int)
		dfs = func(i, s int) {
			if cnt == k || i == len(nums) || s+nums[i] > sumLimit {
				return
			}
			cnt++ // s + nums[i] <= sumLimit
			dfs(i+1, s+nums[i]) // 选
			dfs(i+1, s) // 不选
		}
		dfs(0, 0)
		return cnt == k // 找到 k 个元素和不超过 sumLimit 的子序列
	})
	return int64(sum - kthS)
}
```

```js [sol-JavaScript]
var kSum = function(nums, k) {
    let sum = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] >= 0) {
            sum += nums[i];
        } else {
            nums[i] = -nums[i];
        }
    }
    nums.sort((a, b) => a - b);

    let cnt;
    // 反向递归，增加改成减少，这样可以少传一些参数
    function dfs(i, s) {
        if (cnt === 0 || i === nums.length || s < nums[i]) {
            return;
        }
        cnt--;
        dfs(i + 1, s - nums[i], nums); // 选
        dfs(i + 1, s, nums); // 不选
    }

    let left = -1, right = _.sum(nums);
    while (left + 1 < right) { // 开区间二分，原理见【前置知识】
        const mid = Math.floor((left + right) / 2);
        cnt = k - 1; // 空子序列算一个
        dfs(0, mid, nums);
        if (cnt === 0) { // 找到 k 个元素和不超过 mid 的子序列
            right = mid;
        } else {
            left = mid;
        }
    }
    return sum - right;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn k_sum(mut nums: Vec<i32>, k: i32) -> i64 {
        let mut sum = 0;
        for x in &mut nums {
            if *x >= 0 {
                sum += *x as i64;
            } else {
                *x = -*x;
            }
        }
        nums.sort_unstable();

        // 反向递归，增加改成减少，这样可以少传一些参数
        fn dfs(i: usize, s: i64, nums: &Vec<i32>, cnt: &mut i32) {
            if *cnt == 0 || i == nums.len() || s < nums[i] as i64 {
                return;
            }
            *cnt -= 1;
            dfs(i + 1, s - nums[i] as i64, nums, cnt); // 选
            dfs(i + 1, s, nums, cnt); // 不选
        }

        let mut left = -1;
        let mut right = nums.iter().map(|&x| x as i64).sum::<i64>();
        while left + 1 < right { // 开区间二分，原理见【前置知识】
            let mid = (left + right) / 2;
            let mut cnt = k - 1; // 空子序列算一个
            dfs(0, mid, &nums, &mut cnt); // 找到 k 个元素和不超过 mid 的子序列
            if cnt == 0 {
                right = mid;
            } else {
                left = mid;
            }
        }
        sum - right
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + k\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\sum\limits_{i=0}^{n-1}|\textit{nums}[i]|$。注意 $\texttt{check}$ 的时间复杂度取决于递归中 `cnt++` 的次数，这不超过 $k-1$。
- 空间复杂度：$\mathcal{O}(\min(k,n))$。空间复杂度取决于递归的栈开销，我们至多递归 $\min(k,n)$ 层。

## 方法二：最小堆

如何**不重不漏**地生成 $\textit{nums}$ 的所有子序列？

以**有序非负数组** $\textit{nums}=[1,2,3]$ 为例，有 $2^3=8$ 个子序列，生成的方法如下：

1. 从空子序列 $[]$ 开始。
2. 在 $[]$ 末尾添加 $1$ 得到 $[1]$。
3. 在 $[1]$ 末尾添加 $2$ 得到 $[1,2]$。也可以把末尾的 $1$ 替换成 $2$ 得到 $[2]$。
4. 在 $[2]$ 末尾添加 $3$ 得到 $[2,3]$。也可以把末尾的 $2$ 替换成 $3$ 得到 $[3]$。
5. 在 $[1,2]$ 末尾添加 $3$ 得到 $[1,2,3]$。也可以把末尾的 $2$ 替换成 $3$ 得到 $[1,3]$。

上述过程结合**最小堆**，就可以按照从小到大的顺序生成所有子序列了（堆中维护子序列的和以及下一个要添加/替换的元素下标）。取生成的第 $k$ 个子序列的和，作为 $\textit{sum}$ 要减去的数。

#### 答疑

**问**：这是怎么想到的？为什么这种做法可以生成所有的子序列？

**答**：这种做法本质是把 [78. 子集](https://leetcode.cn/problems/subsets/) 的「枚举选哪个」写法的 `for` 循环去掉，把循环变量 `j` 加到了递归参数中。毕竟 `for` 循环就相当于不断地替换最后一个数，而往下递归则相当于在末尾添加数字。关于「枚举选哪个」可以看我的[【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)中的「答案视角」。

**问**：为什么结合最小堆，就一定是按元素和从小到大的顺序生成的？有没有可能先生成一个大的，再生成一个小的？

**答**：把子序列和它通过添加/替换生成的子序列之间连一条有向边，我们可以得到一棵以空子序列 $[]$ 为根的有向树。把边权定义成相邻节点的子序列元素和的差，由于 $\textit{nums}$ 是有序的且没有负数，所以树是**没有负数边权**的。那么上述算法其实就是在这棵树上跑 [Dijkstra 算法](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。把元素和当作海拔高度，算法执行过程就好比不断上涨的水位，我们会按照海拔高度从低到高淹没节点，所以出堆的元素和是非降的。

```py [sol-Python3]
class Solution:
    def kSum(self, nums: List[int], k: int) -> int:
        sum = 0
        for i, x in enumerate(nums):
            if x >= 0:
                sum += x
            else:
                nums[i] = -x
        nums.sort()

        h = [(0, 0)]  # 空子序列
        for _ in range(k - 1):
            s, i = heappop(h)
            if i < len(nums):
                # 在子序列的末尾添加 nums[i]
                heappush(h, (s + nums[i], i + 1))  # 下一个添加/替换的元素下标为 i+1
                if i:  # 替换子序列的末尾元素为 nums[i]
                    heappush(h, (s + nums[i] - nums[i - 1], i + 1))
        return sum - h[0][0]
```

```java [sol-Java]
class Solution {
    public long kSum(int[] nums, int k) {
        long sum = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] >= 0) {
                sum += nums[i];
            } else {
                nums[i] = -nums[i];
            }
        }
        Arrays.sort(nums);

        PriorityQueue<Pair<Long, Integer>> pq = new PriorityQueue<>((a, b) -> Long.compare(a.getKey(), b.getKey()));
        pq.offer(new Pair<>(0L, 0)); // 空子序列
        while (--k > 0) {
            Pair<Long, Integer> p = pq.poll();
            long s = p.getKey();
            int i = p.getValue();
            if (i < nums.length) {
                // 在子序列的末尾添加 nums[i]
                pq.offer(new Pair<>(s + nums[i], i + 1)); // 下一个添加/替换的元素下标为 i+1
                if (i > 0) { // 替换子序列的末尾元素为 nums[i]
                    pq.offer(new Pair<>(s + nums[i] - nums[i - 1], i + 1));
                }
            }
        }
        return sum - pq.peek().getKey();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long kSum(vector<int> &nums, int k) {
        long long sum = 0;
        for (int &x : nums) {
            if (x >= 0) {
                sum += x;
            } else {
                x = -x;
            }
        }
        ranges::sort(nums);

        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
        pq.emplace(0, 0); // 空子序列
        while (--k) {
            auto [s, i] = pq.top();
            pq.pop();
            if (i < nums.size()) {
                // 在子序列的末尾添加 nums[i]
                pq.emplace(s + nums[i], i + 1); // 下一个添加/替换的元素下标为 i+1
                if (i) { // 替换子序列的末尾元素为 nums[i]
                    pq.emplace(s + nums[i] - nums[i - 1], i + 1);
                }
            }
        }
        return sum - pq.top().first;
    }
};
```

```go [sol-Go]
func kSum(nums []int, k int) int64 {
	n := len(nums)
	sum := 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
		} else {
			nums[i] = -x
		}
	}
	slices.Sort(nums)

	h := hp{{0, 0}} // 空子序列
	for ; k > 1; k-- {
		p := heap.Pop(&h).(pair)
		i := p.i
		if i < n {
		    // 在子序列的末尾添加 nums[i]
			heap.Push(&h, pair{p.sum + nums[i], i + 1}) // 下一个添加/替换的元素下标为 i+1
			if i > 0 { // 替换子序列的末尾元素为 nums[i]
				heap.Push(&h, pair{p.sum + nums[i] - nums[i-1], i + 1})
			}
		}
	}
	return int64(sum - h[0].sum)
}

type pair struct{ sum, i int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

```js [sol-JavaScript]
var kSum = function(nums, k) {
    let sum = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] >= 0) {
            sum += nums[i];
        } else {
            nums[i] = -nums[i];
        }
    }
    nums.sort((a, b) => a - b);

    const pq = new MinPriorityQueue({priority: e => e[0]});
    pq.enqueue([0, 0]); // 空子序列
    while (--k) {
        const [s, i] = pq.dequeue().element;
        if (i < nums.length) {
            // 在子序列的末尾添加 nums[i]
            pq.enqueue([s + nums[i], i + 1]); // 下一个添加/替换的元素下标为 i+1
            if (i > 0) { // 替换子序列的末尾元素为 nums[i]
                pq.enqueue([s + nums[i] - nums[i - 1], i + 1]);
            }
        }
    }
    return sum - pq.front().element[0];
};
```

```rust [sol-Rust]
use std::collections::BinaryHeap;

impl Solution {
    pub fn k_sum(mut nums: Vec<i32>, mut k: i32) -> i64 {
        let mut sum = 0;
        for x in &mut nums {
            if *x >= 0 {
                sum += *x as i64;
            } else {
                *x = -*x;
            }
        }
        nums.sort_unstable();

        // 注意：为方便实现，改成最大堆，在一开始把 sum 入堆，循环中的加法改成减法
        let mut h = BinaryHeap::new();
        h.push((sum, 0)); // 空子序列
        while k > 1 {
            let (s, i) = h.pop().unwrap();
            if i < nums.len() {
                // 在子序列的末尾添加 nums[i]
                h.push((s - nums[i] as i64, i + 1)); // 下一个添加/替换的元素下标为 i+1
                if i > 0 { // 替换子序列的末尾元素为 nums[i]
                    h.push((s - nums[i] as i64 + nums[i - 1] as i64, i + 1));
                }
            }
            k -= 1;
        }
        h.peek().unwrap().0
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + k\log k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 题单：第 K 小/大

- [378. 有序矩阵中第 K 小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/)
- [668. 乘法表中第 K 小的数](https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/)
- [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/)
- [719. 找出第 K 小的数对距离](https://leetcode.cn/problems/find-k-th-smallest-pair-distance/)
- [1201. 丑数 III](https://leetcode.cn/problems/ugly-number-iii/) 2039
- [1439. 有序矩阵中的第 k 个最小数组和](https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/) 2134
- [786. 第 K 个最小的素数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/) 2169
- [2040. 两个有序数组的第 K 小乘积](https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/) 2518
- [1918. 第 K 小的子数组和](https://leetcode.cn/problems/kth-smallest-subarray-sum/)（会员题）

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
