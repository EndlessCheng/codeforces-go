**前置题目**：[3355. 零数组变换 I](https://leetcode.cn/problems/zero-array-transformation-i/)

以 $\textit{nums}=[2,0,2,0,2]$ 为例。

从左到右遍历数组。对于 $\textit{nums}[0]=2$ 来说，我们必须从 $\textit{queries}$ 中选两个左端点为 $0$ 的区间。选哪两个呢？

贪心地想，**区间的右端点越大越好**，后面的元素越小，后续操作就越少。所以选两个左端点为 $0$，且右端点最大的区间。

继续，由于 $\textit{nums}[1]=0$，可以直接跳过。

继续，遍历到 $\textit{nums}[2]=2$，我们需要知道左端点 $\le 2$ 的剩余未选区间中，右端点最大的两个区间。

这启发我们用**最大堆**维护左端点 $\le i$ 的未选区间的右端点。

剩下的就是用**差分数组**去维护区间减一了。你需要先完成 [3355. 零数组变换 I](https://leetcode.cn/problems/zero-array-transformation-i/) 这题。

最终堆的大小（剩余没有使用的区间个数）就是答案。

代码实现时，还需要把 $\textit{queries}$ 按照左端点排序，这样我们可以用双指针遍历 $\textit{nums}$ 和左端点 $\le i$ 的区间。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1uzBxYoEJC/?t=4m40s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxRemoval(self, nums: List[int], queries: List[List[int]]) -> int:
        queries.sort(key=lambda q: q[0])
        h = []
        diff = [0] * (len(nums) + 1)
        sum_d = j = 0
        for i, x in enumerate(nums):
            sum_d += diff[i]
            while j < len(queries) and queries[j][0] <= i:
                heappush(h, -queries[j][1])  # 取相反数表示最大堆
                j += 1
            while sum_d < x and h and -h[0] >= i:
                sum_d += 1
                diff[-heappop(h) + 1] -= 1
            if sum_d < x:
                return -1
        return len(h)
```

```java [sol-Java]
class Solution {
    public int maxRemoval(int[] nums, int[][] queries) {
        Arrays.sort(queries, (a, b) -> a[0] - b[0]);
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> b - a);
        int n = nums.length;
        int[] diff = new int[n + 1];
        int sumD = 0;
        int j = 0;
        for (int i = 0; i < n; i++) {
            sumD += diff[i];
            while (j < queries.length && queries[j][0] <= i) {
                pq.add(queries[j][1]);
                j++;
            }
            while (sumD < nums[i] && !pq.isEmpty() && pq.peek() >= i) {
                sumD++;
                diff[pq.poll() + 1]--;
            }
            if (sumD < nums[i]) {
                return -1;
            }
        }
        return pq.size();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxRemoval(vector<int>& nums, vector<vector<int>>& queries) {
        ranges::sort(queries, {}, [](auto& q) { return q[0]; });
        int n = nums.size(), j = 0, sum_d = 0;
        vector<int> diff(n + 1, 0);
        priority_queue<int> pq;
        for (int i = 0; i < n; i++) {
            sum_d += diff[i];
            while (j < queries.size() && queries[j][0] <= i) {
                pq.push(queries[j][1]);
                j++;
            }
            while (sum_d < nums[i] && !pq.empty() && pq.top() >= i) {
                sum_d++;
                diff[pq.top() + 1]--;
                pq.pop();
            }
            if (sum_d < nums[i]) {
                return -1;
            }
        }
        return pq.size();
    }
};
```

```go [sol-Go]
func maxRemoval(nums []int, queries [][]int) int {
	slices.SortFunc(queries, func(a, b []int) int { return a[0] - b[0] })
	h := hp{}
	diff := make([]int, len(nums)+1)
	sumD, j := 0, 0
	for i, x := range nums {
		sumD += diff[i]
		for ; j < len(queries) && queries[j][0] <= i; j++ {
			heap.Push(&h, queries[j][1])
		}
		for sumD < x && h.Len() > 0 && h.IntSlice[0] >= i {
			sumD++
			diff[heap.Pop(&h).(int)+1]--
		}
		if sumD < x {
			return -1
		}
	}
	return h.Len()
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 相似题目

- [871. 最低加油次数](https://leetcode.cn/problems/minimum-number-of-refueling-stops/)

另外，可以做做下面贪心题单中的「**§1.9 反悔贪心**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
