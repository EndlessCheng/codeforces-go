## 方法一：暴力枚举

子数组中任意两个数按位与均为 $0$，意味着任意两个数对应的集合的**交集为空**（见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)）。

这意味着子数组中的从低到高的第 $i$ 个比特位上，至多有一个比特 $1$，其余均为比特 $0$。例如子数组不可能有两个奇数（最低位为 $1$)，否则这两个数按位与是大于 $0$ 的。

根据鸽巢原理（抽屉原理），在本题数据范围下，优雅子数组的长度不会超过 $30$。例如子数组为 $[2^0,2^1,2^2,\cdots,2^{29}]$，我们无法再加入一个数 $x$，使 $x$ 与子数组中的任何一个数按位与均为 $0$。

既然长度不会超过 $30$，直接暴力枚举子数组的右端点 $i$ 即可。

代码实现时，可以把在子数组中的元素**按位或**起来（求并集），这样可以 $\mathcal{O}(1)$ 判断当前元素是否与前面的元素按位与的结果为 $0$（交集为空）。

[视频讲解](https://www.bilibili.com/video/BV1Dt4y1j7qh) 第三题。

```py [sol-Python3]
class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        ans = 0
        for i, or_ in enumerate(nums):  # 枚举子数组右端点 i
            j = i - 1
            while j >= 0 and (or_ & nums[j]) == 0:  # nums[j] 与子数组中的任意元素 AND 均为 0
                or_ |= nums[j]  # 加到子数组中
                j -= 1  # 向左扩展
            ans = max(ans, i - j)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestNiceSubarray(int[] nums) {
        int ans = 0;
        for (int i = 0; i < nums.length; i++) { // 枚举子数组右端点 i
            int or = 0;
            int j = i;
            while (j >= 0 && (or & nums[j]) == 0) { // nums[j] 与子数组中的任意元素 AND 均为 0
                or |= nums[j--]; // 加到子数组中
            }
            ans = Math.max(ans, i - j);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestNiceSubarray(vector<int>& nums) {
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) { // 枚举子数组右端点 i
            int or_ = 0, j = i;
            while (j >= 0 && (or_ & nums[j]) == 0) { // nums[j] 与子数组中的任意元素 AND 均为 0
                or_ |= nums[j--]; // 加到子数组中
            }
            ans = max(ans, i - j);
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestNiceSubarray(nums []int) (ans int) {
	for i, or := range nums { // 枚举子数组右端点 i
		j := i - 1
		for ; j >= 0 && or&nums[j] == 0; j-- { // nums[j] 与子数组中的任意元素 AND 均为 0
			or |= nums[j] // 加到子数组中
		}
		ans = max(ans, i-j)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

## 方法二：滑动窗口

不了解滑动窗口的同学请看 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

进一步地，既然这些数对应的集合的交集为空，我们可以用**滑动窗口**优化上述过程，右边加入 $\textit{nums}[\textit{right}]$，左边移出 $\textit{nums}[\textit{left}]$。如果 $\textit{or}$ 与新加入的 $\textit{nums}[\textit{right}]$ 有交集，则不断从 $\textit{or}$ 中去掉集合 $\textit{nums}[\textit{left}]$，直到 $\textit{or}$ 与 $\textit{nums}[\textit{right}]$ 交集为空。

如何把集合语言翻译成位运算代码，见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

```py [sol-Python3]
class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        ans = left = or_ = 0
        for right, x in enumerate(nums):
            while or_ & x:  # 有交集
                or_ ^= nums[left]  # 从 or_ 中去掉集合 nums[left]
                left += 1
            or_ |= x  # 把集合 x 并入 or_ 中
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestNiceSubarray(int[] nums) {
        int ans = 0;
        int or = 0;
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            while ((or & nums[right]) > 0) { // 有交集
                or ^= nums[left++]; // 从 or 中去掉集合 nums[left]
            }
            or |= nums[right]; // 把集合 nums[right] 并入 or 中
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestNiceSubarray(vector<int>& nums) {
        int ans = 0, left = 0, or_ = 0;
        for (int right = 0; right < nums.size(); right++) {
            while (or_ & nums[right]) { // 有交集
                or_ ^= nums[left++]; // 从 or 中去掉集合 nums[left]
            }
            or_ |= nums[right]; // 把集合 nums[right] 并入 or 中
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 { // 有交集
			or ^= nums[left] // 从 or 中去掉集合 nums[left]
			left += 1
		}
		or |= x // 把集合 x 并入 or 中
		ans = max(ans, right-left+1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。虽然写了个二重循环，但 $\textit{left}$ 一直在增大，内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $O(n)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
