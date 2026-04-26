首先，倒序遍历 $\textit{nums}$，同时用一个布尔数组记录哪些元素，严格大于其右侧所有元素。我们可以用一个变量 $\textit{mx}$ 记录遍历过的元素的最大值。如果 $\textit{nums}[i] > \textit{mx}$，那么 $\textit{nums}[i]$ 严格大于其右侧所有元素。

然后，正序遍历 $\textit{nums}$，做法同上，我们可以知道 $\textit{nums}[i]$ 是否严格大于其左侧所有元素。

如果两个条件其中一个成立，那么把 $\textit{nums}[i]$ 添加到答案中。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def findValidElements(self, nums: list[int]) -> list[int]:
        # 标记严格大于其右侧所有元素的元素
        n = len(nums)
        right_valid = [False] * n
        mx = 0
        for i in range(n - 1, -1, -1):
            x = nums[i]
            right_valid[i] = x > mx
            mx = max(mx, x)

        ans = []
        mx = 0
        for valid, x in zip(right_valid, nums):
            if valid or x > mx:
                ans.append(x)
            mx = max(mx, x)
        return ans
```

```java [sol-Java]
class Solution {
    public List<Integer> findValidElements(int[] nums) {
        // 标记严格大于其右侧所有元素的元素
        int n = nums.length;
        boolean[] rightValid = new boolean[n];
        int mx = 0;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            rightValid[i] = x > mx;
            mx = Math.max(mx, x);
        }

        List<Integer> ans = new ArrayList<>();
        mx = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x > mx || rightValid[i]) {
                ans.add(x);
            }
            mx = Math.max(mx, x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findValidElements(vector<int>& nums) {
        // 标记严格大于其右侧所有元素的元素
        int n = nums.size();
        vector<int8_t> right_valid(n);
        int mx = 0;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            right_valid[i] = x > mx;
            mx = max(mx, x);
        }

        vector<int> ans;
        mx = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x > mx || right_valid[i]) {
                ans.push_back(x);
            }
            mx = max(mx, x);
        }
        return ans;
    }
};
```

```go [sol-Go]
func findValidElements(nums []int) (ans []int) {
	// 标记严格大于其右侧所有元素的元素
	rightValid := make([]bool, len(nums))
	mx := 0
	for i, x := range slices.Backward(nums) {
		rightValid[i] = x > mx
		mx = max(mx, x)
	}

	mx = 0
	for i, x := range nums {
		if x > mx || rightValid[i] {
			ans = append(ans, x)
		}
		mx = max(mx, x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
