本题是 [80. 删除有序数组中的重复项 II](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/) 更一般的版本，那题相当于 $k=2$。做法见 [我的题解](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/solutions/3060042/yong-zhan-si-kao-yuan-di-shi-xian-python-zw8l/)。

注意本题保证 $k\le n$。

```py [sol-Python3]
class Solution:
    def limitOccurrences(self, nums: list[int], k: int) -> list[int]:
        stack_size = k  # 栈的大小，前 k 个元素默认保留
        for i in range(k, len(nums)):
            if nums[i] != nums[stack_size - k]:  # 和栈的倒数第 k 个数比较
                nums[stack_size] = nums[i]  # 入栈
                stack_size += 1

        del nums[stack_size:]
        return nums
```

```java [sol-Java]
class Solution {
    public int[] limitOccurrences(int[] nums, int k) {
        int stackSize = k; // 栈的大小，前 k 个元素默认保留
        for (int i = k; i < nums.length; i++) {
            if (nums[i] != nums[stackSize - k]) { // 和栈的倒数第 k 个数比较
                nums[stackSize++] = nums[i]; // 入栈
            }
        }
        return Arrays.copyOf(nums, stackSize);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> limitOccurrences(vector<int>& nums, int k) {
        int stack_size = k; // 栈的大小，前 k 个元素默认保留
        for (int i = k; i < nums.size(); i++) {
            if (nums[i] != nums[stack_size - k]) { // 和栈的倒数第 k 个数比较
                nums[stack_size++] = nums[i]; // 入栈
            }
        }

        nums.resize(stack_size);
        return nums;
    }
};
```

```go [sol-Go]
func limitOccurrences(nums []int, k int) []int {
	stackSize := k // 栈的大小，前 k 个元素默认保留
	for i := k; i < len(nums); i++ {
		if nums[i] != nums[stackSize-k] { // 和栈的倒数第 k 个数比较
			nums[stackSize] = nums[i] // 入栈
			stackSize++
		}
	}
	return nums[:stackSize]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n-k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面双指针题单的「**§3.5 原地修改**」。

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
