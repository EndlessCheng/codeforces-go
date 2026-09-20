**题意**：重排 $\textit{nums}$，让二进制最高位的前缀 $1$ 尽可能长；在这个前提下，让次高位的前缀 $1$ 尽可能长；依此类推。

从最高位开始思考。既然要把 $1$ 排在前面，那么直接把 $\textit{nums}$ 从大到小排序，然后统计最高位的前缀 $1$ 的长度。

例如 $[7,5,2,1]$ 这四个数的二进制如下，最高位的最长前缀 $1$（竖着看）的长度为 $2$。

```
111
101
010
001
```

在最高位的前缀 $1$ 尽可能长的前提下，次高位的最长前缀 $1$（竖着看）的长度只能是 $1$。其他元素的次高位是 $0$ 还是 $1$ 不重要，可以一视同仁。为方便排序，其他元素的次高位可以都置为 $0$。

置 $0$ 后，重新从大到小排序，得到

```
111
101
001
000
```

这样**既可以让低位的 $1$ 尽量往前排，又保证了高位的 $1$ 的位置不变，高位的最长前缀 $1$ 的长度不变**。

重复上述过程，直到二进制的最低位。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def largestPower(self, nums: list[int]) -> list[int]:
        n = len(nums)
        ans = [0] * 15
        nums.sort(reverse=True)
        for i in range(nums[0].bit_length() - 1, -1, -1):
            # 找最长前缀连续 1
            j = 0
            while j < n and nums[j] >> i & 1:
                j += 1
            ans[14 - i] = j

            # [0, j-1] 这一位都是 1，其余元素无关紧要，为方便排序，全置为 0
            mask = ~(1 << i)
            for k in range(j, n):
                nums[k] &= mask  # nums[k] 的 i 位置为 0
            nums.sort(reverse=True)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] largestPower(int[] nums) {
        int n = nums.length;
        int[] ans = new int[15];
        Arrays.sort(nums);
        int maxWidth = 32 - Integer.numberOfLeadingZeros(nums[n - 1]);
        for (int i = maxWidth - 1; i >= 0; i--) {
            // 找最长前缀连续 1
            int j = n - 1;
            while (j >= 0 && (nums[j] >> i & 1) > 0) {
                j--;
            }
            ans[14 - i] = n - 1 - j;

            // [j+1, n-1] 这一位都是 1，其余元素无关紧要，为方便排序，全置为 0
            for (; j >= 0; j--) {
                nums[j] &= ~(1 << i);
            }
            Arrays.sort(nums);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> largestPower(vector<int> nums) {
        vector<int> ans(15);
        ranges::sort(nums, greater());
        int max_width = bit_width(1u * nums[0]);
        for (int i = max_width - 1; i >= 0; i--) {
            // 找最长前缀连续 1
            int j = 0;
            while (j < nums.size() && (nums[j] >> i & 1)) {
                j++;
            }
            ans[14 - i] = j;

            // [0, j-1] 这一位都是 1，其余元素无关紧要，为方便排序，全置为 0
            for (; j < nums.size(); j++) {
                nums[j] &= ~(1 << i);
            }
            ranges::sort(nums, greater());
        }
        return ans;
    }
};
```

```go [sol-Go]
func largestPower(nums []int) []int {
	ans := [15]int{}
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	maxWidth := bits.Len(uint(nums[0]))
	for i := maxWidth - 1; i >= 0; i-- {
		// 找最长前缀连续 1
		j := 0
		for j < len(nums) && nums[j]>>i&1 > 0 {
			j++
		}
		ans[14-i] = j

		// [0, j-1] 这一位都是 1，其余元素无关紧要，为方便排序，全置为 0
		for ; j < len(nums); j++ {
			nums[j] &^= 1 << i
		}
		slices.SortFunc(nums, func(a, b int) int { return b - a })
	}
	return ans[:]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。忽略排序的栈开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
