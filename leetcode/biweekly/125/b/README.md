每次选择**最小**的两个数操作，这可以用**最小堆**模拟。

> 部分语言可以直接修改堆顶：弹出 $x$ 后，把堆顶增加 $2x$。

如果最小值 $\ge k$，那么所有数都 $\ge k$。所以循环直到堆顶 $\ge k$ 为止。

> 注意题目保证答案一定存在。也就是说，堆中只有一个数且这个数小于 $k$ 的情况是不存在的。

```py [sol-Python3]
class Solution:
    def minOperations(self, h: List[int], k: int) -> int:
        heapify(h)
        ans = 0
        while h[0] < k:
            x = heappop(h)
            heapreplace(h, h[0] + x * 2)
            ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int ans = 0;
        PriorityQueue<Long> pq = new PriorityQueue<>();
        for (int x : nums) {
            pq.offer((long) x);
        }

        while (pq.peek() < k) {
            long x = pq.poll();
            long y = pq.poll();
            pq.offer(x * 2 + y);
            ans++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        priority_queue<long long, vector<long long>, greater<>> pq(nums.begin(), nums.end());
        int ans = 0;
        while (pq.top() < k) {
            long long x = pq.top(); pq.pop();
            long long y = pq.top(); pq.pop();
            pq.push(x * 2 + y);
            ans++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) (ans int) {
	h := &hp{nums}
	heap.Init(h)
	for h.IntSlice[0] < k {
		x := heap.Pop(h).(int)
		h.IntSlice[0] += x * 2
		heap.Fix(h, 0)
		ans++
	}
	return
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)    {}
func (h *hp) Pop() any { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

```js [sol-JavaScript]
var minOperations = function(nums, k) {
    const pq = new MinPriorityQueue(); // datastructures-js/priority-queue@5.4.0
    for (const x of nums) {
        pq.enqueue(x);
    }

    let ans = 0;
    while (pq.front().element < k) {
        const x = pq.dequeue().element;
        const y = pq.dequeue().element;
        pq.enqueue(x * 2 + y);
        ans++;
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::BinaryHeap;

impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        let mut h = BinaryHeap::new();
        for x in nums {
            h.push(-x as i64); // 取负号变成最小堆
        }

        let mut ans = 0;
        while -h.peek().unwrap() < k as i64 {
            let x = h.pop().unwrap();
            let y = h.pop().unwrap();
            h.push(x * 2 + y);
            ans += 1;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于每次循环都会出堆一个元素，所以循环次数是 $\mathcal{O}(n)$ 的。每次操作堆需要 $\mathcal{O}(\log n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。Python 和 Go 可以直接把输入的 $\textit{nums}$ 当作堆。

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
