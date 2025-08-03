三段式子数组必须满足「严格递增 - 严格递减 - 严格递增」，一共三段，每一段**至少要有两个数**。

利用 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)，我们可以遍历所有的极大三段式子数组。极大的意思是子数组不能再往左右延长。

⚠**注意**：第三段的起点也是下一个极大三段式子数组的第一段的起点。

定义：

- 第一段的范围为 $[\textit{start},\textit{peak}]$。
- 第二段的范围为 $[\textit{peak},\textit{bottom}]$。
- 第三段从 $\textit{bottom}$ 开始。

⚠**注意**：第一二段之间的峰顶 $\textit{peak}$ 是第一二段共享的，第二三段之间的谷底 $\textit{bottom}$ 是第二三段共享的。

该三段式子数组中的最大三段式子数组和，由三部分组成：

1. 必须包含从 $\textit{peak}-1$ 到 $\textit{bottom}+1$ 的所有元素。（每一段至少有两个数）
2. 从第一段的倒数第三个数（下标 $\textit{peak}-2$）开始，往左的最大元素和。如果不存在则为 $0$。
3. 从第三段的第三个数（下标 $\textit{bottom}+2$）开始，往右的最大元素和。如果不存在则为 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1BEh3zZEoM/?t=37m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxSumTrionic(self, nums: List[int]) -> int:
        n = len(nums)
        ans = -inf
        i = 0
        while i < n:
            # 第一段
            start = i
            i += 1
            while i < n and nums[i - 1] < nums[i]:
                i += 1
            if i == start + 1:  # 第一段至少要有两个数
                continue

            # 第二段
            peak = i - 1
            res = nums[peak - 1] + nums[peak]  # 第一段的最后两个数必选
            while i < n and nums[i - 1] > nums[i]:
                res += nums[i]  # 第二段的所有元素必选
                i += 1
            if i == peak + 1 or i == n:  # 第二段至少要有两个数，第三段至少要有两个数
                continue

            # 第三段
            bottom = i - 1
            res += nums[i]  # 第三段的前两个数必选（第一个数在上面的循环中加了）
            # 从第三段的第三个数往右，计算最大元素和
            max_s = s = 0
            i += 1
            while i < n and nums[i - 1] < nums[i]:
                s += nums[i]
                max_s = max(max_s, s)
                i += 1
            res += max_s

            # 从第一段的倒数第三个数往左，计算最大元素和
            max_s = s = 0
            j = peak - 2
            while j >= start:
                s += nums[j]
                max_s = max(max_s, s)
                j -= 1
            res += max_s
            ans = max(ans, res)

            i = bottom  # 第三段的起点也是下一个极大三段式子数组的第一段的起点
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSumTrionic(int[] nums) {
        int n = nums.length;
        long ans = Long.MIN_VALUE;
        for (int i = 0; i < n;) {
            // 第一段
            int start = i;
            for (i++; i < n && nums[i - 1] < nums[i]; i++);
            if (i == start + 1) { // 第一段至少要有两个数
                continue;
            }

            // 第二段
            int peak = i - 1;
            long res = nums[peak - 1] + nums[peak]; // 第一段的最后两个数必选
            for (; i < n && nums[i - 1] > nums[i]; i++) {
                res += nums[i]; // 第二段的所有元素必选
            }
            if (i == peak + 1 || i == n) { // 第二段至少要有两个数，第三段至少要有两个数
                continue;
            }

            // 第三段
            int bottom = i - 1;
            res += nums[i]; // 第三段的前两个数必选（第一个数在上面的循环中加了）
            // 从第三段的第三个数往右，计算最大元素和
            long maxS = 0;
            long s = 0;
            for (i++; i < n && nums[i - 1] < nums[i]; i++) {
                s += nums[i];
                maxS = Math.max(maxS, s);
            }
            res += maxS;

            // 从第一段的倒数第三个数往左，计算最大元素和
            maxS = 0;
            s = 0;
            for (int j = peak - 2; j >= start; j--) {
                s += nums[j];
                maxS = Math.max(maxS, s);
            }
            res += maxS;
            ans = Math.max(ans, res);

            i = bottom; // 第三段的起点也是下一个极大三段式子数组的第一段的起点
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSumTrionic(vector<int>& nums) {
        int n = nums.size();
        long long ans = LLONG_MIN;
        for (int i = 0; i < n;) {
            // 第一段
            int start = i;
            for (i++; i < n && nums[i - 1] < nums[i]; i++);
            if (i == start + 1) { // 第一段至少要有两个数
                continue;
            }

            // 第二段
            int peak = i - 1;
            long long res = nums[peak - 1] + nums[peak]; // 第一段的最后两个数必选
            for (; i < n && nums[i - 1] > nums[i]; i++) {
                res += nums[i]; // 第二段的所有元素必选
            }
            if (i == peak + 1 || i == n) { // 第二段至少要有两个数，第三段至少要有两个数
                continue;
            }

            // 第三段
            int bottom = i - 1;
            res += nums[i]; // 第三段的前两个数必选（第一个数在上面的循环中加了）
            // 从第三段的第三个数往右，计算最大元素和
            long long max_s = 0, s = 0;
            for (i++; i < n && nums[i - 1] < nums[i]; i++) {
                s += nums[i];
                max_s = max(max_s, s);
            }
            res += max_s;

            // 从第一段的倒数第三个数往左，计算最大元素和
            max_s = 0; s = 0;
            for (int j = peak - 2; j >= start; j--) {
                s += nums[j];
                max_s = max(max_s, s);
            }
            res += max_s;
            ans = max(ans, res);

            i = bottom; // 第三段的起点也是下一个极大三段式子数组的第一段的起点
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	ans := math.MinInt
	for i := 0; i < n; {
		// 第一段
		start := i
		for i++; i < n && nums[i-1] < nums[i]; i++ {
		}
		if i == start+1 { // 第一段至少要有两个数
			continue
		}

		// 第二段
		peak := i - 1
		res := nums[peak-1] + nums[peak] // 第一段的最后两个数必选
		for ; i < n && nums[i-1] > nums[i]; i++ {
			res += nums[i] // 第二段的所有元素必选
		}
		if i == peak+1 || i == n { // 第二段至少要有两个数，第三段至少要有两个数
			continue
		}

		// 第三段
		bottom := i - 1
		res += nums[i] // 第三段的前两个数必选（第一个数在上面的循环中加了）
		// 从第三段的第三个数往右，计算最大元素和
		maxS, s := 0, 0
		for i++; i < n && nums[i-1] < nums[i]; i++ {
			s += nums[i]
			maxS = max(maxS, s)
		}
		res += maxS

		// 从第一段的倒数第三个数往左，计算最大元素和
		maxS, s = 0, 0
		for j := peak - 2; j >= start; j-- {
			s += nums[j]
			maxS = max(maxS, s)
		}
		res += maxS
		ans = max(ans, res)

		i = bottom // 第三段的起点也是下一个极大三段式子数组的第一段的起点
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。见 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)。对于本题，同一个元素可以在两个相交的极大三段式子数组中各遍历一次。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面滑动窗口与双指针题单的「**六、分组循环**」。

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
