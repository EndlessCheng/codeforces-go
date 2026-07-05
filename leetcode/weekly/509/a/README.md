遍历的同时，维护数字范围的最大值 $\textit{maxRange}$。

- 如果 $\textit{nums}[i]$ 的数字范围比 $\textit{maxRange}$ 还大，那么更新 $\textit{maxRange}$，同时重置答案为 $\textit{nums}[i]$（重新累加）。
- 如果 $\textit{nums}[i]$ 的数字范围等于 $\textit{maxRange}$，那么把 $\textit{nums}[i]$ 加到答案中。
- 如果 $\textit{nums}[i]$ 的数字范围小于 $\textit{maxRange}$，不变。

[本题视频讲解](https://www.bilibili.com/video/BV1ioTC6BECj/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxDigitRange(self, nums: list[int]) -> int:
        ans = max_range = 0

        for x in nums:
            s = str(x)  # 不用字符串的写法，见另一份代码
            r = ord(max(s)) - ord(min(s))
            if r > max_range:
                max_range = r
                ans = x  # 重新累加
            elif r == max_range:
                ans += x

        return ans
```

```py [sol-Python3 不用字符串]
class Solution:
    def maxDigitRange(self, nums: list[int]) -> int:
        ans = max_range = 0

        for x in nums:
            mn, mx = 9, 0
            v = x
            while v > 0:
                v, d = divmod(v, 10)
                mn = min(mn, d)
                mx = max(mx, d)

            r = mx - mn
            if r > max_range:
                max_range = r
                ans = x  # 重新累加
            elif r == max_range:
                ans += x

        return ans
```

```java [sol-Java]
class Solution {
    public int maxDigitRange(int[] nums) {
        int maxRange = 0;
        int ans = 0;

        for (int x : nums) {
            int mn = 9;
            int mx = 0;
            for (int v = x; v > 0; v /= 10) {
                int d = v % 10;
                mn = Math.min(mn, d);
                mx = Math.max(mx, d);
            }

            int r = mx - mn;
            if (r > maxRange) {
                maxRange = r;
                ans = x; // 重新累加
            } else if (r == maxRange) {
                ans += x;
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDigitRange(vector<int>& nums) {
        int max_range = 0;
        int ans = 0;

        for (int x : nums) {
            int mn = 9, mx = 0;
            for (int v = x; v > 0; v /= 10) {
                int d = v % 10;
                mn = min(mn, d);
                mx = max(mx, d);
            }

            int r = mx - mn;
            if (r > max_range) {
                max_range = r;
                ans = x; // 重新累加
            } else if (r == max_range) {
                ans += x;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxDigitRange(nums []int) (ans int) {
	maxRange := 0
	for _, x := range nums {
		mn, mx := 9, 0
		for v := x; v > 0; v /= 10 {
			d := v % 10
			mn = min(mn, d)
			mx = max(mx, d)
		}

		r := mx - mn
		if r > maxRange {
			maxRange = r
			ans = x // 重新累加
		} else if r == maxRange {
			ans += x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

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
