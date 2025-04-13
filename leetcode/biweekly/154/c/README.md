虽然题目要求 $i\le j\le k$，但因为异或运算满足交换律 $a\oplus b = b\oplus a$，实际上我们可以随意选。所以本质上，这题就是从 $\textit{nums}$ 中（可重复地）选三个数。

首先，算出任意两数异或的所有可能值，在本题的数据范围下，这不会超过 $2^{11}-1=2047$。

然后遍历两数异或的所有可能值，再与 $\textit{nums}$ 中的数计算异或，就得到了三数异或的所有可能值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        nums = list(set(nums))  # 优化：去重可以减少循环次数
        st = {x ^ y for x, y in combinations(nums, 2)} | {0}
        return len({xy ^ z for xy in st for z in nums})
```

```java [sol-Java]
class Solution {
    public int uniqueXorTriplets(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        int u = 1 << (32 - Integer.numberOfLeadingZeros(mx));

        boolean[] has = new boolean[u];
        for (int i = 0; i < nums.length; i++) {
            for (int j = i; j < nums.length; j++) {
                has[nums[i] ^ nums[j]] = true;
            }
        }

        boolean[] has3 = new boolean[u];
        for (int xy = 0; xy < u; xy++) {
            if (!has[xy]) {
                continue;
            }
            for (int z : nums) {
                has3[xy ^ z] = true;
            }
        }

        int ans = 0;
        for (boolean b : has3) {
            if (b) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int uniqueXorTriplets(vector<int>& nums) {
        int n = nums.size();
        int u = 1 << bit_width((unsigned) ranges::max(nums));

        vector<int> has(u);
        for (int i = 0; i < n; i++) {
            for (int j = i; j < n; j++) {
                has[nums[i] ^ nums[j]] = true;
            }
        }

        vector<int> has3(u);
        for (int xy = 0; xy < u; xy++) {
            if (has[xy]) {
                for (int z : nums) {
                    has3[xy ^ z] = true;
                }
            }
        }

        return reduce(has3.begin(), has3.end());
    }
};
```

```go [sol-Go]
func uniqueXorTriplets(nums []int) (ans int) {
	u := 1 << bits.Len(uint(slices.Max(nums)))

	has := make([]bool, u)
	for i, x := range nums {
		for _, y := range nums[i:] {
			has[x^y] = true
		}
	}

	has3 := make([]bool, u)
	for xy, b := range has {
		if !b {
			continue
		}
		for _, z := range nums {
			has3[xy^z] = true
		}
	}

	for _, b := range has3 {
		if b {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+U))$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。

## 附：异或 FWT

原理见 [快速沃尔什变换（FWT）](https://oi-wiki.org/math/poly/fwt/)。

FWT 模板。甚至还能求出一个数组 $\textit{cnt}_3$，其中 $\textit{cnt}_3[i]$ 表示三数异或恰好等于 $i$ 的三元组**个数**。

本题相当于统计有多少个 $\textit{cnt}_3[i]>0$。

```go
func fwtXOR(a []int, rsh int) {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				a[i+j], a[i+j+k] = (a[i+j]+a[i+j+k])>>rsh, (a[i+j]-a[i+j+k])>>rsh
			}
		}
	}
}

func fwtXOR3(a []int) []int {
	fwtXOR(a, 0)
	for i, x := range a {
		a[i] *= x * x
	}
	fwtXOR(a, 1)
	return a
}

func uniqueXorTriplets(nums []int) (ans int) {
	cnt := make([]int, 1<<bits.Len(uint(slices.Max(nums))))
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range fwtXOR3(cnt) {
		if c > 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + U\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
