思路同 [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)，本题维护前后缀的 $\text{GCD}$ 和 $\text{LCM}$。

细节：

1. 由于 $0$ 可以被任何非零整数整除，所以 $\text{GCD}(0,x)=x$。
2. 由于任何整数都是 $1$ 的倍数，所以 $\text{LCM}(1,x)=x$。

> 注：$[1,30]$ 中的所有元素的 $\text{LCM}$ 等于 $2329089562800$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hn1MYhEtC/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxScore(self, nums: List[int]) -> int:
        n = len(nums)
        suf_gcd = [0] * (n + 1)
        suf_lcm = [0] * n + [1]
        for i in range(n - 1, -1, -1):
            suf_gcd[i] = gcd(suf_gcd[i + 1], nums[i])
            suf_lcm[i] = lcm(suf_lcm[i + 1], nums[i])

        ans = suf_gcd[0] * suf_lcm[0]  # 不移除元素
        pre_gcd, pre_lcm = 0, 1
        for i, x in enumerate(nums):  # 枚举移除 nums[i]
            ans = max(ans, gcd(pre_gcd, suf_gcd[i + 1]) * lcm(pre_lcm, suf_lcm[i + 1]))
            pre_gcd = gcd(pre_gcd, x)
            pre_lcm = lcm(pre_lcm, x)
        return ans
```

```java [sol-Java]
class Solution {
    public long maxScore(int[] nums) {
        int n = nums.length;
        int[] sufGcd = new int[n + 1];
        long[] sufLcm = new long[n + 1];
        sufLcm[n] = 1;
        for (int i = n - 1; i >= 0; i--) {
            sufGcd[i] = (int) gcd(sufGcd[i + 1], nums[i]);
            sufLcm[i] = lcm(sufLcm[i + 1], nums[i]);
        }

        long ans = sufGcd[0] * sufLcm[0]; // 不移除元素
        int preGcd = 0;
        long preLcm = 1;
        for (int i = 0; i < n; i++) { // 枚举移除 nums[i]
            ans = Math.max(ans, gcd(preGcd, sufGcd[i + 1]) * lcm(preLcm, sufLcm[i + 1]));
            preGcd = (int) gcd(preGcd, nums[i]);
            preLcm = lcm(preLcm, nums[i]);
        }
        return ans;
    }

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxScore(vector<int>& nums) {
        int n = nums.size();
        vector<int> suf_gcd(n + 1);
        vector<long long> suf_lcm(n + 1);
        suf_lcm[n] = 1;
        for (int i = n - 1; i >= 0; i--) {
            suf_gcd[i] = gcd(suf_gcd[i + 1], nums[i]);
            suf_lcm[i] = lcm(suf_lcm[i + 1], nums[i]);
        }

        long long ans = suf_gcd[0] * suf_lcm[0]; // 不移除元素
        int pre_gcd = 0;
        long long pre_lcm = 1;
        for (int i = 0; i < n; i++) { // 枚举移除 nums[i]
            ans = max(ans, gcd(pre_gcd, suf_gcd[i + 1]) * lcm(pre_lcm, suf_lcm[i + 1]));
            pre_gcd = gcd(pre_gcd, nums[i]);
            pre_lcm = lcm(pre_lcm, nums[i]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxScore(nums []int) int64 {
	n := len(nums)
	sufGcd := make([]int, n+1)
	sufLcm := make([]int, n+1)
	sufLcm[n] = 1
	for i, x := range slices.Backward(nums) {
		sufGcd[i] = gcd(sufGcd[i+1], x)
		sufLcm[i] = lcm(sufLcm[i+1], x)
	}

	ans := sufGcd[0] * sufLcm[0] // 不移除元素
	preGcd, preLcm := 0, 1
	for i, x := range nums { // 枚举移除 nums[i]
		ans = max(ans, gcd(preGcd, sufGcd[i+1])*lcm(preLcm, sufLcm[i+1]))
		preGcd = gcd(preGcd, x)
		preLcm = lcm(preLcm, x)
	}
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

更多相似题目，见下面动态规划题单中的「**专题：前后缀分解**」。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
