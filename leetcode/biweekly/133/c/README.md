由于 $\textit{nums}[i]$ 会被其左侧元素的操作影响，所以我们先从最左边的 $\textit{nums}[0]$ 开始思考。

分类讨论：

- 如果 $\textit{nums}[0]=1$，无需反转，问题变成剩下 $n-1$ 个数如何操作。接下来考虑 $\textit{nums}[1]$。
- 如果 $\textit{nums}[0]=0$，反转次数加一，问题变成剩下 $n-1$ 个数（在修改次数是奇数的情况下）如何操作。接下来考虑 $\textit{nums}[1]$。

对后续元素来说，由于反转偶数次等于没反转，所以只需考虑反转次数的奇偶性。

一般地，设反转次数的奇偶性为 $k$，分类讨论：

- 如果 $\textit{nums}[i]\ne k$，无需反转，接下来考虑 $\textit{nums}[i+1]$。
- 如果 $\textit{nums}[i]=k$，反转次数加一，接下来考虑 $\textit{nums}[i+1]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17w4m1e7Nw/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        for x in nums:
            if x == ans % 2:
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int ans = 0;
        for (int x : nums) {
            if (x == ans % 2) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int ans = 0;
        for (int x : nums) {
            ans += x == ans % 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) (ans int) {
	for _, x := range nums {
		if x == ans%2 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
