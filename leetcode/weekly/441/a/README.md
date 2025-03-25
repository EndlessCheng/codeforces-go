负数肯定不能选，先都去掉。

「子数组中的所有元素互不相同」意味着每个非负数只能选一个，所以答案就是 $\textit{nums}$ 中的非负数去重后的元素和。

⚠**注意**：如果数组中都是负数，由于题目规定不能都去掉，只能选一个最大的负数（绝对值最小的负数）作为答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JYQ8YWEvD/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int]) -> int:
        st = set(x for x in nums if x >= 0)
        return sum(st) if st else max(nums)
```

```py [sol-Python3 写法二]
class Solution:
    def maxSum(self, nums: List[int]) -> int:
        return sum(set(x for x in nums if x >= 0)) or max(nums)
```

```java [sol-Java]
class Solution {
    public int maxSum(int[] nums) {
        Set<Integer> set = new HashSet<>();
        int s = 0;
        int mx = Integer.MIN_VALUE;
        for (int x : nums) { // 一次遍历
            if (x < 0) {
                mx = Math.max(mx, x);
            } else if (set.add(x)) { // x 不在 set 中
                s += x;
            }
        }
        return set.isEmpty() ? mx : s;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSum(vector<int>& nums) {
        unordered_set<int> st;
        int s = 0, mx = INT_MIN;
        for (int x : nums) { // 一次遍历
            if (x < 0) {
                mx = max(mx, x);
            } else if (st.insert(x).second) { // x 不在 set 中
                s += x;
            }
        }
        return st.empty() ? mx : s;
    }
};
```

```go [sol-Go]
func maxSum(nums []int) (ans int) {
	set := map[int]struct{}{}
	mx := math.MinInt
	for _, x := range nums { // 一次遍历
		if x < 0 {
			mx = max(mx, x)
		} else if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			ans += x
		}
	}
	if len(set) == 0 {
		return mx
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
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
