## 分治

回顾示例 3 $\textit{nums}=[1,2,1,2,1,2]$ 的操作过程：

- 首先，只需要一次操作（选择整个数组），就可以把所有的最小值 $1$ 都变成 $0$。现在数组是 $[0,2,0,2,0,2]$。
- 这些被 $0$ 分割开的 $2$，无法合在一起操作（因为子数组会包含 $0$，导致 $2$ 无法变成 $0$），只能一个一个操作。

一般地：

1. 先通过一次操作，把 $\textit{nums}$ 的最小值都变成 $0$（如果最小值已经是 $0$ 则跳过这步）。
2. 此时 $\textit{nums}$ 被这些 $0$ 划分成了若干段，后续操作只能在每段内部，不能跨段操作（否则子数组会包含 $0$）。每一段是规模更小的子问题，可以用第一步的方法解决。这样我们可以写一个递归去处理。递归边界：如果操作后全为 $0$，直接返回。

找最小值可以用 ST 表或者线段树，但这种做法很麻烦。有没有简单的做法呢？

## 单调栈

从左往右遍历数组，只在「必须要操作」的时候，才把答案加一。

什么时候必须要操作一个数？

示例 3 $\textit{nums}=[1,2,1,2,1,2]$，因为 $2$ 左右两侧都有小于 $2$ 的数，都需要单独操作。

又例如 $\textit{nums}=[1,2,3,2,1]$：

- 遍历到第二个 $2$ 时，可以知道 $3$ 左右两侧都有小于 $3$ 的数，所以 $3$ 必须要操作一次，答案加一。注意这不表示第一次操作的是 $3$，而是某次操作会把 $3$ 变成 $0$。
- 遍历到末尾 $1$ 时，可以知道中间的两个 $2$，左边有 $1$，右边也有 $1$，必须操作一次，答案加一。比如选择 $[2,3,2]$ 可以把这两个 $2$ 都变成 $0$。
- 最后，数组中的 $1$ 需要操作一次都变成 $0$。

我们怎么知道「$3$ 左右两侧都有小于 $3$ 的数」？

遍历数组的同时，把遍历过的元素用栈记录：

- 如果当前元素比栈顶大（或者栈为空），那么直接入栈。
- 如果当前元素比栈顶小，那么对于栈顶来说，左边（栈顶倒数第二个数）比栈顶小（原因后面解释），右边（当前元素）也比栈顶小，所以栈顶必须操作一次。然后弹出栈顶。
- 如果当前元素等于栈顶，可以在同一次操作中把当前元素与栈顶都变成 $0$，所以无需入栈。注意这保证了栈中没有重复元素。

如果当前元素比栈顶小，就弹出栈顶，我们会得到一个底小顶大的单调栈，这就保证了「对于栈顶来说，左边（栈顶倒数第二个数）比栈顶小」。

遍历结束后，因为栈是严格递增的，所以栈中每个非零数字都需要操作一次。

代码实现时，可以直接把 $\textit{nums}$ 当作栈。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        st = []
        for x in nums:
            while st and x < st[-1]:
                st.pop()
                ans += 1
            # 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
            if not st or x != st[-1]:
                st.append(x)
        return ans + len(st) - (st[0] == 0)  # 0 不需要操作
```

```py [sol-Python3 原地]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        top = -1  # 栈顶下标（把 nums 当作栈）
        for x in nums:
            while top >= 0 and x < nums[top]:
                top -= 1  # 出栈
                ans += 1
            # 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
            if top < 0 or x != nums[top]:
                top += 1
                nums[top] = x  # 入栈
        return ans + top + (nums[0] > 0)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int ans = 0;
        int top = -1; // 栈顶下标（把 nums 当作栈）
        for (int x : nums) {
            while (top >= 0 && x < nums[top]) {
                top--; // 出栈
                ans++;
            }
            // 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
            if (top < 0 || x != nums[top]) {
                nums[++top] = x; // 入栈
            }
        }
        return ans + top + (nums[0] > 0 ? 1 : 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int ans = 0;
        int top = -1; // 栈顶下标（把 nums 当作栈）
        for (int x : nums) {
            while (top >= 0 && x < nums[top]) {
                top--; // 出栈
                ans++;
            }
            // 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
            if (top < 0 || x != nums[top]) {
                nums[++top] = x; // 入栈
            }
        }
        return ans + top + (nums[0] > 0);
    }
};
```

```go [sol-Go]
func minOperations(nums []int) (ans int) {
	st := nums[:0] // 原地
	for _, x := range nums {
		for len(st) > 0 && x < st[len(st)-1] {
			st = st[:len(st)-1]
			ans++
		}
		// 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
		if len(st) == 0 || x != st[len(st)-1] {
			st = append(st, x)
		}
	}
	if st[0] == 0 { // 0 不需要操作
		ans--
	}
	return ans + len(st)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。每个元素至多入栈出栈各一次，所以二重循环的循环次数是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面单调栈题单。

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
