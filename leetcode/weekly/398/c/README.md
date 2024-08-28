横看成岭侧成峰，换一个角度，把每一位拆开：

- 计算个位数中的不同数对个数；
- 计算十位数中的不同数对个数；
- 计算百位数中的不同数对个数；
- ……

单独考虑每个数位，此时问题变成：

- 给你一个长为 $n$ 的数组 $a$，只包含数字 $0$ 到 $9$，其中有多少个**不同的数对**？

做法有多种，一次遍历的做法如下。

遍历 $a$，同时用一个长为 $10$ 的数组 $\textit{cnt}$ 统计 $0$ 到 $9$ 每个数字的出现次数。假设现在遍历到 $d=a[k]$，那么前面有 $k$ 个数字，其中有 $\textit{cnt}[d]$ 个数和 $d$ 是一样的，所以有

$$
k - \textit{cnt}[d]
$$

个数和 $d$ 是不一样的，这正是我们要统计的，加入答案。

代码实现时，可以外层循环枚举个位数、十位数、百位数等，内层循环枚举 $\textit{nums}$；也可以外层循环枚举 $\textit{nums}$，内层循环枚举个位数、十位数、百位数等。下面代码用的后者。

附：[视频讲解](https://www.bilibili.com/video/BV19D421G7mw/) 第三题，欢迎点赞关注！

### 答疑

**问**：为什么代码要先更新 $\textit{ans}$，再更新 $\textit{cnt}$？

**答**：如果先更新 $\textit{cnt}$，再更新 $\textit{ans}$ 的话，假设 $\textit{nums}=[1]$，这样写会算出 $-1$。或者说，题目要求两个数的下标是不同的，如果先更新 $\textit{cnt}$，就把下标相同的数对也考虑进来了。

```py [sol-Python3]
class Solution:
    def sumDigitDifferences(self, nums: List[int]) -> int:
        ans = 0
        cnt = [[0] * 10 for _ in str(nums[0])]
        for k, x in enumerate(nums):
            i = 0
            while x:
                x, d = divmod(x, 10)
                ans += k - cnt[i][d]
                cnt[i][d] += 1
                i += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long sumDigitDifferences(int[] nums) {
        long ans = 0;
        int[][] cnt = new int[Integer.toString(nums[0]).length()][10];
        for (int k = 0; k < nums.length; k++) {
            int x = nums[k];
            for (int i = 0; x > 0; x /= 10, i++) {
                ans += k - cnt[i][x % 10]++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long sumDigitDifferences(vector<int>& nums) {
        long long ans = 0;
        vector<array<int, 10>> cnt(to_string(nums[0]).length());
        for (int k = 0; k < nums.size(); k++) {
            int x = nums[k];
            for (int i = 0; x; x /= 10, i++) {
                ans += k - cnt[i][x % 10]++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumDigitDifferences(nums []int) (ans int64) {
	cnt := make([][10]int, len(strconv.Itoa(nums[0])))
	for k, x := range nums {
		for i := 0; x > 0; x /= 10 {
			d := x % 10
			ans += int64(k - cnt[i][d])
			cnt[i][d]++
			i++
		}
	}
	return
}
```

也可以逆向思考，设 $\textit{nums}[0]$ 的十进制长度为 $m$，那么总共有

$$
\dfrac{mn(n-1)}{2}
$$

个数对。

在此基础上，减去**相同的数对**。这就和 [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/) 完全一样了。

```py [sol-Python3]
class Solution:
    def sumDigitDifferences(self, nums: List[int]) -> int:
        n, m = len(nums), len(str(nums[0]))
        ans = m * n * (n - 1) // 2
        cnt = [[0] * 10 for _ in range(m)]
        for x in nums:
            i = 0
            while x:
                x, d = divmod(x, 10)
                ans -= cnt[i][d]
                cnt[i][d] += 1
                i += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long sumDigitDifferences(int[] nums) {
        int n = nums.length;
        int m = Integer.toString(nums[0]).length();
        long ans = (long) m * n * (n - 1) / 2;
        int[][] cnt = new int[m][10];
        for (int x : nums) {
            for (int i = 0; x > 0; x /= 10) {
                ans -= cnt[i++][x % 10]++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long sumDigitDifferences(vector<int>& nums) {
        long long n = nums.size(), m = to_string(nums[0]).length();
        long long ans = m * n * (n - 1) / 2;
        vector<array<int, 10>> cnt(m);
        for (int x : nums) {
            for (int i = 0; x; x /= 10) {
                ans -= cnt[i++][x % 10]++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumDigitDifferences(nums []int) int64 {
	n, m := len(nums), len(strconv.Itoa(nums[0]))
	ans := m * n * (n - 1) / 2
	cnt := make([][10]int, m)
	for _, x := range nums {
		for i := 0; x > 0; x /= 10 {
			d := x % 10
			ans -= cnt[i][d]
			cnt[i][d]++
			i++
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\textit{nums}[0]$。
- 空间复杂度：$\mathcal{O}(D\log U)$，其中 $D=10$。

## 相关题目

- [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/) 第零章，枚举右维护左。
- [位运算题单](https://leetcode.cn/circle/discuss/dHn9Vk/) 第四章，拆位/贡献法。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
