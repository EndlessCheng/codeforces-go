**前置知识**：前缀和，原理请看 [前缀和及其扩展，附题单！](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)

定义一个长为 $n-1$ 的数组 $a$，其中

$$
a[i] =
\begin{cases} 
0,\ &\textit{nums}[i]\bmod 2 \ne \textit{nums}[i+1]\bmod 2\\
1,\ &\textit{nums}[i]\bmod 2 = \textit{nums}[i+1]\bmod 2
\end{cases}
$$

那么 $a$ 的从 $\textit{from}$ 到 $\textit{to}-1$ 的子数组和等于 $0$，就说明询问的子数组是特殊数组。

计算 $a$ 的前缀和 $s$，可以快速判断子数组和是否为 $0$。

代码实现时，相邻两数的异或和的最低位取反，即为 $a[i]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV19D421G7mw/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def isArraySpecial(self, nums: List[int], queries: List[List[int]]) -> List[bool]:
        s = list(accumulate(((x ^ y ^ 1) & 1 for x, y in pairwise(nums)), initial=0))
        return [s[from_] == s[to] for from_, to in queries]
```

```java [sol-Java]
class Solution {
    public boolean[] isArraySpecial(int[] nums, int[][] queries) {
        int[] s = new int[nums.length];
        for (int i = 1; i < nums.length; i++) {
            s[i] = s[i - 1] + ((nums[i] ^ nums[i - 1] ^ 1) & 1);
        }
        boolean[] ans = new boolean[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            ans[i] = s[q[0]] == s[q[1]];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> isArraySpecial(vector<int>& nums, vector<vector<int>>& queries) {
        vector<int> s(nums.size());
        for (int i = 1; i < nums.size(); i++) {
            s[i] = s[i - 1] + ((nums[i] ^ nums[i - 1] ^ 1) & 1);
        }
        vector<bool> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            auto& q = queries[i];
            ans[i] = s[q[0]] == s[q[1]];
        }
        return ans;
    }
};
```

```go [sol-Go]
func isArraySpecial(nums []int, queries [][]int) []bool {
	s := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1] + (nums[i]^nums[i-1]^1)&1
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = s[q[0]] == s[q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
