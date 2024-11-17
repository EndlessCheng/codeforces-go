题意可以转换成：

- 把 $[l_i,r_i]$ 中的元素都减一，最终数组中的所有元素是否都 $\le 0$？

如果所有元素都 $\le 0$，那么我们可以撤销一部分元素的减一，使其调整为 $0$，从而满足原始题意的要求。

这可以用**差分数组**计算，[原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)（推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看）。

[本题视频讲解](https://www.bilibili.com/video/BV1yiU6YnEfU/?t=8m17s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def isZeroArray(self, nums: List[int], queries: List[List[int]]) -> bool:
        diff = [0] * (len(nums) + 1)
        for l, r in queries:
            # 区间 [l,r] 中的数都加一
            diff[l] += 1
            diff[r + 1] -= 1

        for x, sum_d in zip(nums, accumulate(diff)):
            # 此时 sum_d 表示 x=nums[i] 要减掉多少
            if x > sum_d:  # x 无法变成 0
                return False
        return True
```

```java [sol-Java]
class Solution {
    public boolean isZeroArray(int[] nums, int[][] queries) {
        int n = nums.length;
        int[] diff = new int[n + 1];
        for (int[] q : queries) {
            // 区间 [l,r] 中的数都加一
            diff[q[0]]++;
            diff[q[1] + 1]--;
        }

        int sumD = 0;
        for (int i = 0; i < n; i++) {
            sumD += diff[i];
            // 此时 sumD 表示 nums[i] 要减掉多少
            if (nums[i] > sumD) { // nums[i] 无法变成 0
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<int> diff(n + 1);
        for (auto& q : queries) {
            // 区间 [l,r] 中的数都加一
            diff[q[0]]++;
            diff[q[1] + 1]--;
        }

        int sum_d = 0;
        for (int i = 0; i < n; i++) {
            sum_d += diff[i];
            // 此时 sum_d 表示 nums[i] 要减掉多少
            if (nums[i] > sum_d) { // nums[i] 无法变成 0
                return false;
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func isZeroArray(nums []int, queries [][]int) bool {
	diff := make([]int, len(nums)+1)
	for _, q := range queries {
		// 区间 [l,r] 中的数都加一
		diff[q[0]]++
		diff[q[1]+1]--
	}

	sumD := 0
	for i, x := range nums {
		sumD += diff[i]
		// 此时 sumD 表示 x=nums[i] 要减掉多少
		if x > sumD { // x 无法变成 0
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**§2.1 一维差分**」。

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
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
