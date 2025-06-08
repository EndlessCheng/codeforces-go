判断能否都变成 $1$，或者都变成 $-1$。

假设都变成 $\textit{target}$。

从左到右遍历 $\textit{nums}$，如果 $\textit{nums}[i]\ne \textit{target}$，那么需要操作：把 $\textit{nums}[i+1]$ 乘上 $-1$，把剩余操作次数减一。注意 $\textit{nums}[i]$ 不需要乘上 $-1$，因为后面不会再遍历 $\textit{nums}[i]$。

如果要操作的时候，剩余操作次数等于 $0$，或者 $i=n-1$，那么无法操作，返回 $\texttt{false}$。

如果正常遍历结束，返回 $\texttt{true}$。

代码实现时，可以用一个变量 $\textit{mul}$ 表示 $\textit{nums}[i]$ 要乘以的数，从而避免修改 $\textit{nums}$（注意我们要判断两次，第一次的修改不能影响第二次）。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def canMakeEqual(self, nums: List[int], k: int) -> bool:
        def check(target: int) -> bool:
            n = len(nums)
            left = k
            mul = 1
            for i, x in enumerate(nums):
                if x * mul == target:
                    mul = 1  # 下一个数不用乘 -1
                    continue
                if left == 0 or i == n - 1:
                    return False
                left -= 1
                mul = -1  # 下一个数要乘 -1
            return True
        return check(-1) or check(1)
```

```java [sol-Java]
class Solution {
    public boolean canMakeEqual(int[] nums, int k) {
        return check(nums, k, -1) || check(nums, k, 1);
    }

    private boolean check(int[] nums, int k, int target) {
        int n = nums.length;
        int left = k;
        int mul = 1;
        for (int i = 0; i < n; i++) {
            if (nums[i] * mul == target) {
                mul = 1; // 下一个数不用乘 -1
                continue;
            }
            if (left == 0 || i == n - 1) {
                return false;
            }
            left--;
            mul = -1; // 下一个数要乘 -1
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canMakeEqual(vector<int>& nums, int k) {
        auto check = [&](int target) -> bool {
            int left = k;
            int mul = 1;
            for (int i = 0; i < nums.size(); i++) {
                if (nums[i] * mul == target) {
                    mul = 1; // 下一个数不用乘 -1
                    continue;
                }
                if (left == 0 || i + 1 == nums.size()) {
                    return false;
                }
                left--;
                mul = -1; // 下一个数要乘 -1
            }
            return true;
        };
        return check(-1) || check(1);
    }
};
```

```go [sol-Go]
func canMakeEqual(nums []int, k int) bool {
	check := func(target int) bool {
		left := k
		mul := 1
		for i, x := range nums {
			if x*mul == target {
				mul = 1 // 下一个数不用乘 -1
				continue
			}
			if left == 0 || i == len(nums)-1 {
				return false
			}
			left--
			mul = -1 // 下一个数要乘 -1
		}
		return true
	}
	return check(-1) || check(1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
