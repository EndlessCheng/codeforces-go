例如 $\textit{nums}=[4,3,1,2,1]$，在从左到右遍历的过程中，由于 $2$ 的出现，左边的 $1$ 永远不可能与右边的 $1$ 组成一个题目要求的子数组。所以当遍历到 $2$ 时，左边的 $1$ 就是**无用数据**了，可以清除。清除后我们会得到一个从左到右递减的数据结构。

这个性质和 [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)很像，启发我们用单调栈思考。具体来说：

1. 初始化答案等于 $n$，因为每个元素可以单独组成一个长为 $1$ 的子数组，满足题目要求。
2. 维护一个底大顶小的单调栈，记录元素及其出现次数。
3. 从左到右遍历 $\textit{nums}$。
4. 只要 $x=\textit{nums}[i]$ 大于栈顶，就把栈顶出栈。
5. 如果 $x$ 小于栈顶，把 $x$ 及其出现次数 $1$ 入栈。
6. 如果 $x$ 等于栈顶，设栈顶记录的出现次数为 $\textit{cnt}$，那么 $x$ 可以和左边 $\textit{cnt}$ 个 $x$ 组成 $\textit{cnt}$ 个满足要求的子数组，把答案增加 $\textit{cnt}$，然后把 $\textit{cnt}$ 加一。

注意可能出现某个元素 $v$ 出栈后，又重新入栈的情况，此时 $v$ 的出现次数会重置成 $1$。

代码实现时，可以往栈底加入一个 $\infty$ 哨兵，从而简化判断逻辑。

请看 [视频讲解](https://www.bilibili.com/video/BV1et42177VM/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfSubarrays(self, nums: List[int]) -> int:
        ans = len(nums)
        st = [[inf, 0]]  # 无穷大哨兵
        for x in nums:
            while x > st[-1][0]:
                st.pop()
            if x == st[-1][0]:
                ans += st[-1][1]
                st[-1][1] += 1
            else:
                st.append([x, 1])
        return ans
```

```java [sol-Java]
class Solution {
    public long numberOfSubarrays(int[] nums) {
        long ans = nums.length;
        Deque<int[]> st = new ArrayDeque<>();
        st.push(new int[]{Integer.MAX_VALUE, 0}); // 无穷大哨兵
        for (int x : nums) {
            while (x > st.peek()[0]) {
                st.pop();
            }
            if (x == st.peek()[0]) {
                ans += st.peek()[1]++;
            } else {
                st.push(new int[]{x, 1});
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfSubarrays(vector<int>& nums) {
        long long ans = nums.size();
        stack<pair<int, int>> st;
        st.emplace(INT_MAX, 0); // 无穷大哨兵
        for (int x : nums) {
            while (x > st.top().first) {
                st.pop();
            }
            if (x == st.top().first) {
                ans += st.top().second++;
            } else {
                st.emplace(x, 1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubarrays(nums []int) int64 {
	ans := len(nums)
	type pair struct{ x, cnt int }
	st := []pair{{math.MaxInt, 0}} // 无穷大哨兵
	for _, x := range nums {
		for x > st[len(st)-1].x {
			st = st[:len(st)-1]
		}
		if x == st[len(st)-1].x {
			ans += st[len(st)-1].cnt
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{x, 1})
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面的单调栈题单。

## 思考题

1. 改成子数组第一个 **或** 最后一个元素是最大值，要怎么做？
2. 改成子数组第一个元素是最大值，最后一个元素是 **最小值**，要怎么做？
3. 改成树上路径问题，见 [2421. 好路径的数目](https://leetcode.cn/problems/number-of-good-paths/)。本题相当于把 2421 的树特化成一条链。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
