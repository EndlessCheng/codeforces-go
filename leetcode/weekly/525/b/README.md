把 $\textit{nums}[i]$ 从大到小排列，即为 $\textit{nums}[i]$ 能变成的最大数。

设操作后 $\textit{nums}[i]$ 增大了 $d[i]$。

取增量最大的 $k$ 个数操作，可以得到最优解。

```py [sol-Python3]
class Solution:
    def maximizeSum(self, nums: list[int], k: int) -> int:
        ans = 0

        for i, x in enumerate(nums):
            ans += x
            s = sorted(str(x), reverse=True)
            y = int(''.join(s))
            # 把修改后的增量保存到 nums[i] 中
            nums[i] = y - x

        nums.sort(reverse=True)
        # 选择增量最大的 k 个数
        return ans + sum(nums[:k])
```

```java [sol-Java]
class Solution {
    public long maximizeSum(int[] nums, int k) {
        int n = nums.length;
        long ans = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            ans += x;

            // 把 x 从大到小排列，得到 y
            char[] s = String.valueOf(x).toCharArray();
            Arrays.sort(s);
            int y = 0;
            for (int j = s.length - 1; j >= 0; j--) {
                y = y * 10 + s[j] - '0';
            }

            // 把修改后的增量保存到 nums[i] 中
            nums[i] = y - x;
        }

        Arrays.sort(nums);
        // 选择增量最大的 k 个数
        for (int i = n - 1; i >= n - k; i--) {
            ans += nums[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximizeSum(vector<int>& nums, int k) {
        long long ans = 0;
        for (int& x : nums) {
            ans += x;

            // 把 x 从大到小排列，得到 y
            string s = to_string(x);
            ranges::sort(s, greater());
            int y = stoi(s);

            // 把修改后的增量保存到 nums[i] 中
            x = y - x;
        }

        // ranges::sort(nums, greater());
        ranges::nth_element(nums, nums.begin() + k, greater());
        // 选择增量最大的 k 个数
        return ans + reduce(nums.begin(), nums.begin() + k, 0LL);
    }
};
```

```go [sol-Go]
func maximizeSum(nums []int, k int) (ans int64) {
	for i, x := range nums {
		ans += int64(x)

		// 把 x 从大到小排列，得到 y
		s := []byte(strconv.Itoa(x))
		slices.SortFunc(s, func(a, b byte) int { return int(b) - int(a) })
		y, _ := strconv.Atoi(string(s))

		// 把修改后的增量保存到 nums[i] 中
		nums[i] = y - x
	}

	slices.Sort(nums)
	// 选择增量最大的 k 个数
	for _, d := range nums[len(nums)-k:] {
		ans += int64(d)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\log m + n\log n)$ 或 $\mathcal{O}(nm\log m + n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m=\log \max(\textit{nums})$ 是 $\textit{nums}[i]$ 的十进制长度。第二个排序，用快速选择可以做到 $\mathcal{O}(n)$，见 C++ 代码。
- 空间复杂度：$\mathcal{O}(m)$。忽略排序的栈开销。

## 专题训练

见下面贪心题单的「**§1.1 从最小/最大开始贪心**」。

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
