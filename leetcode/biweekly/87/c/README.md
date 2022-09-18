下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

本文将讲述一种处理此类问题的通用模板，可以做到

1. 求出**所有**子数组的按位或的结果，以及值等于该结果的子数组的个数。
2. 求按位或结果等于**任意给定数字**的子数组的最短长度/最长长度。

文章末尾列出了大量题目，均可以用该模板秒杀。

---

思考：对于起始位置为 $i$ 的子数组的按位或，至多有多少种不同的结果？

根据或运算的性质，我们可以从 $x=\textit{nums}[i]$ 开始，不断往右扩展子数组，按位或的结果要么使 $x$ 不变，要么让 $x$ 的某些比特位的值由 $0$ 变 $1$。最坏情况下从 $x=0$ 出发，每次改变一个比特位，最终得到 $2^{29}-1<10^9$，因此至多有 $30$ 种不同的结果。

另一个结论是，相同的按位或对应的子数组右端点会形成一个连续的区间。

据此，我们可以倒着遍历 $\textit{nums}$，在遍历的同时，用一个数组 $\textit{ors}$ 维护从 $\textit{nums}[i]$ 开始的子数组的按位或的结果，及其对应的子数组右端点的最小值。继续遍历到 $\textit{nums}[i-1]$ 时，我们可以把 $\textit{nums}[i-1]$ 和 $\textit{ors}$ 中的每个值按位或，合并值相同的结果。

这样在遍历时，$\textit{ors}$ 中值最大的元素对应的子数组右端点的最小值，就是要求的最短子数组的右端点。

```py [sol1-Python3]
class Solution:
    def smallestSubarrays(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n
        ors = []  # 按位或的值 + 对应子数组的右端点的最小值
        for i in range(n - 1, -1, -1):
            num = nums[i]
            ors.append([0, i])
            ors[0][0] |= num
            k = 0
            for p in ors[1:]:
                p[0] |= num
                if ors[k][0] == p[0]:
                    ors[k][1] = p[1]  # 合并相同值，下标取最小的
                else:
                    k += 1
                    ors[k] = p
            del ors[k + 1:]
            # 本题只用到了 ors[0]，如果题目改成任意给定数值，可以在 ors 中查找
            ans[i] = ors[0][1] - i + 1
        return ans
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> smallestSubarrays(vector<int> &nums) {
        int n = nums.size();
        vector<int> ans(n);
        vector<pair<int, int>> ors; // 按位或的值 + 对应子数组的右端点的最小值
        for (int i = n - 1; i >= 0; --i) {
            ors.emplace_back(0, i);
            ors[0].first |= nums[i];
            int k = 0;
            for (int j = 1; j < ors.size(); ++j) {
                ors[j].first |= nums[i];
                if (ors[k].first == ors[j].first)
                    ors[k].second = ors[j].second; // 合并相同值，下标取最小的
                else ors[++k] = ors[j];
            }
            ors.resize(k + 1);
            // 本题只用到了 ors[0]，如果题目改成任意给定数值，可以在 ors 中查找
            ans[i] = ors[0].second - i + 1;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	type pair struct{ or, i int }
	ors := []pair{} // 按位或的值 + 对应子数组的右端点的最小值
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		ors = append(ors, pair{0, i})
		ors[0].or |= num
		k := 0
		for _, p := range ors[1:] {
			p.or |= num
			if ors[k].or == p.or {
				ors[k].i = p.i // 合并相同值，下标取最小的
			} else {
				k++
				ors[k] = p
			}
		}
		ors = ors[:k+1]
        // 本题只用到了 ors[0]，如果题目改成任意给定数值，可以在 ors 中查找
		ans[i] = ors[0].i - i + 1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(\log U)$。

#### 可以用模板秒杀的题目

按位或：

- [898. 子数组按位或操作](https://leetcode.cn/problems/bitwise-ors-of-subarrays/)

按位与：

- [1521. 找到最接近目标值的函数值](https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target/)

最大公因数（GCD）：

- [Codeforces 475D. CGCDSSQ](https://codeforces.com/problemset/problem/475/D)
- [Codeforces 1632D. New Year Concert](https://codeforces.com/problemset/problem/1632/D)

乘法：

- [蓝桥杯2021年第十二届国赛真题-和与乘积](https://www.dotcpp.com/oj/problem2622.html)
