把 $s[i] = \texttt{1}$ 理解成一根红色箭头，指向 $\textit{nums}[i]$。我们需要最大化箭头指向（选择）的元素之和。

一次操作相当于把一根箭头左移一个单位（如果左边有空位的话）。

示例 1 的 $\textit{nums}=[2,1,5,2,3]$，$s=\texttt{01010}$。

由于箭头只能左移，无法右移，所以最右边的箭头右侧的数 $[3]$，无法选择。

从最右边的箭头，到倒数第二个箭头（不含）之间的数呢？在示例 1 中，子数组 $b=[5,2]$ 只能被最右边的箭头选择。贪心地，选择其中最大的元素 $\max(b)=5$。如果前面发现更大的，再替换。

从倒数第二个箭头，到倒数第三个箭头（不含）之间的数呢？在示例 1 中，子数组 $b=[2,1]$ 可以被两个箭头选择：

- 倒着看，我们先选择 $1$。目前这两根箭头指向的数为 $[1,5]$。
- 继续倒着遍历，由于 $2>1$，我们舍弃最小的 $1$，改成 $2$。目前这两根箭头指向的数为 $[2,5]$。
- 假如左边还有数，比如 $9$，由于 $9>2$，我们舍弃最小的 $2$，改成 $9$。目前这两根箭头指向的数为 $[9,5]$。
- 假如左边还有数，比如 $99$，由于 $99>5$，我们舍弃最小的 $5$，改成 $99$。目前这两根箭头指向的数为 $[99,9]$。

这个过程类似维护前 $2$ 大，可以用**最小堆**实现：先把元素入堆，再把堆顶（最小值）弹出。

一般地，在倒着遍历的过程中，维护前 $k$ 大元素，其中 $k$ 是遍历过的 $s[i]=\texttt{1}$ 的个数。

- 先把 $\textit{nums}[i]$ 入堆。
- 如果 $s[i]=\texttt{1}$，无需额外操作（新增一个箭头指向的元素）；否则，把堆顶弹出。

最后返回堆中元素之和。

[本题视频讲解](https://www.bilibili.com/video/BV14LqmBMECK/?t=8m20s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumScore(self, nums: List[int], s: str) -> int:
        h = []
        for x, ch in zip(reversed(nums), reversed(s)):
            heappush(h, x)
            if ch == '0':
                heappop(h)  # 更快的写法见【Python3 写法二】
        return sum(h)
```

```py [sol-Python3 写法二]
class Solution:
    def maximumScore(self, nums: List[int], s: str) -> int:
        h = []
        for x, ch in zip(reversed(nums), reversed(s)):
            if ch == '1':
                heappush(h, x)
            else:
                heappushpop(h, x)
        return sum(h)
```

```java [sol-Java]
class Solution {
    public long maximumScore(int[] nums, String s) {
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        long ans = 0;
        for (int i = nums.length - 1; i >= 0; i--) {
            int x = nums[i];
            ans += x;
            pq.offer(x);
            if (s.charAt(i) == '0') {
                ans -= pq.poll();
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumScore(vector<int>& nums, string s) {
        priority_queue<int, vector<int>, greater<>> pq;
        long long ans = 0;
        for (int i = nums.size() - 1; i >= 0; i--) {
            int x = nums[i];
            ans += x;
            pq.push(x);
            if (s[i] == '0') {
                ans -= pq.top();
                pq.pop();
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumScore(nums []int, s string) (ans int64) {
	h := hp{}
	for i, x := range slices.Backward(nums) {
		ans += int64(x)
		heap.Push(&h, x)
		if s[i] == '0' {
			ans -= int64(heap.Pop(&h).(int))
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

```go [sol-Go 写法二]
func maximumScore(nums []int, s string) (ans int64) {
	h := hp{}
	for i, x := range slices.Backward(nums) {
		if s[i] == '1' {
			ans += int64(x)
			heap.Push(&h, x)
		} else if h.Len() > 0 && x > h.IntSlice[0] {
			ans += int64(x - h.IntSlice[0])
			h.IntSlice[0] = x
			heap.Fix(&h, 0)
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp) Pop() (_ any)  { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 是 $\textit{nums}$ 的长度，$k$ 是 $s$ 中的 $\texttt{1}$ 的个数。
- 空间复杂度：$\mathcal{O}(k)$。

## 专题训练

见下面数据结构题单的「**五、堆（优先队列）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
