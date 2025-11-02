设 $\textit{nums}$ 的最小值和最大值分别为 $m$ 和 $M$。

枚举在 $[m+1,M-1]$ 中的整数，如果不在 $\textit{nums}$ 中，加入答案。

为了快速判断一个数是否在 $\textit{nums}$ 中，把 $\textit{nums}$ 转成哈希集合。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def findMissingElements(self, nums: List[int]) -> List[int]:
        st = set(nums)
        return [i for i in range(min(nums) + 1, max(nums)) if i not in st]
```

```java [sol-Java]
class Solution {
    public List<Integer> findMissingElements(int[] nums) {
        int mn = Integer.MAX_VALUE;
        int mx = Integer.MIN_VALUE;
        Set<Integer> st = new HashSet<>();
        for (int x : nums) {
            mn = Math.min(mn, x);
            mx = Math.max(mx, x);
            st.add(x);
        }

        List<Integer> ans = new ArrayList<>();
        for (int i = mn + 1; i < mx; i++) {
            if (!st.contains(i)) {
                ans.add(i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findMissingElements(vector<int>& nums) {
        int mn = ranges::min(nums);
        int mx = ranges::max(nums);
        unordered_set<int> st(nums.begin(), nums.end());

        vector<int> ans;
        for (int i = mn + 1; i < mx; i++) {
            if (!st.contains(i)) {
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMissingElements(nums []int) (ans []int) {
	mn, mx := math.MaxInt, math.MinInt
	has := map[int]bool{}
	for _, x := range nums {
		mn = min(mn, x)
		mx = max(mx, x)
		has[x] = true
	}

	for i := mn + 1; i < mx; i++ {
		if !has[i] {
			ans = append(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D)$，其中 $D$ 是 $\textit{nums}$ 的最大值与最小值之差。注意题目保证 $\textit{nums}$ 所有元素互不相同。
- 空间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。返回值不计入。

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
