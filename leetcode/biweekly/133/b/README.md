分类讨论：

- 如果 $\textit{nums}[0]=1$，无需修改，问题变成剩下 $n-1$ 个数如何操作。接下来考虑 $\textit{nums}[1]$。
- 如果 $\textit{nums}[0]=0$，修改，问题变成剩下 $n-1$ 个数如何操作。接下来考虑 $\textit{nums}[1]$。

所以从左到右遍历数组，一边遍历一边修改。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17w4m1e7Nw/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        for i in range(len(nums) - 2):
            if nums[i] == 0:
                nums[i + 1] ^= 1
                nums[i + 2] ^= 1
                ans += 1
        return ans if nums[-2] and nums[-1] else -1
```

```java [sol-Java]
public class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n - 2; i++) {
            if (nums[i] == 0) {
                nums[i + 1] ^= 1;
                nums[i + 2] ^= 1;
                ans++;
            }
        }
        return nums[n - 2] != 0 && nums[n - 1] != 0 ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 0; i < n - 2; i++) {
            if (nums[i] == 0) {
                nums[i + 1] ^= 1;
                nums[i + 2] ^= 1;
                ans++;
            }
        }
        return nums[n - 2] && nums[n - 1] ? ans : -1;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums[:n-2] {
		if x == 0 {
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			ans++
		}
	}
	if nums[n-2] == 0 || nums[n-1] == 0 {
		return -1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果把题目中的 $3$ 替换成 $k$，你能想出一个 $\mathcal{O}(n)$ 的做法吗？

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
