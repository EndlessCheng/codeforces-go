[视频讲解](https://www.bilibili.com/video/BV1MT411u7fW) 已出炉，包括本题的末尾列出的部分题目，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

## 方法一：利用或运算的性质

我们可以把二进制数看成集合，二进制数第 $i$ 位为 $1$ 表示 $i$ 在集合中。两个二进制数的或，就可以看成是两个集合的**并集**。

对于两个二进制数 $a$ 和 $b$，如果 $a\ |\ b=a$，从集合的角度上看，$b$ 对应的集合是 $a$ 对应的集合的**子集**。

据此我们可以提出如下算法：

从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$：
- 如果 $\textit{nums}[j]\ |\ x\ne\textit{nums}[j]$，说明 $\textit{nums}[j]$ 可以变大（集合元素增多），更新 $\textit{nums}[j]=\textit{nums}[j]\ |\ x$；
- 如果 $\textit{nums}[j]\ |\ x=\textit{nums}[j]$，从集合的角度看，此时 $x$ 不仅是 $\textit{nums}[j]$ 的子集，同时也是 $\textit{nums}[k]\ (k<j)$ 的子集（因为循环保证了每个集合都是其左侧相邻集合的子集），那么后续的循环都无法让元素变大，退出循环；
- 在循环中，如果 $\textit{nums}[j]$ 可以变大，则更新 $\textit{ans}[j]=i-j+1$。

```py [sol11-Python3]
class Solution:
    def smallestSubarrays(self, nums: List[int]) -> List[int]:
        ans = [0] * len(nums)
        for i, x in enumerate(nums):
            ans[i] = 1
            for j in range(i - 1, -1, -1):
                if (nums[j] | x) == nums[j]:
                    break
                nums[j] |= x
                ans[j] = i - j + 1
        return ans
```

```java [sol11-Java]
class Solution {
    public int[] smallestSubarrays(int[] nums) {
        var n = nums.length;
        var ans = new int[n];
        for (var i = 0; i < n; ++i) {
            ans[i] = 1;
            for (var j = i - 1; j >= 0 && (nums[j] | nums[i]) != nums[j]; --j) {
                nums[j] |= nums[i];
                ans[j] = i - j + 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol11-C++]
class Solution {
public:
    vector<int> smallestSubarrays(vector<int> &nums) {
        int n = nums.size();
        vector<int> ans(n);
        for (int i = 0; i < n; ++i) {
            ans[i] = 1;
            for (int j = i - 1; j >= 0 && (nums[j] | nums[i]) != nums[j]; --j) {
                nums[j] |= nums[i];
                ans[j] = i - j + 1;
            }
        }
        return ans;
    }
};
```

```go [sol11-Go]
func smallestSubarrays(nums []int) []int {
    ans := make([]int, len(nums))
    for i, x := range nums {
        ans[i] = 1
        for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
            nums[j] |= x
            ans[j] = i - j + 1
        }
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。由于 $2^{29}-1<10^9<2^{30}-1$，二进制数对应集合的大小不会超过 $29$，因此在或运算下，每个数字至多可以增大 $29$ 次。总体上看，二重循环的次数等于每个数字可以增大的次数之和，即 $O(n\log U)$。
- 空间复杂度：$O(1)$。

## 方法二：更加通用的模板

该模板可以做到

1. 求出**所有**子数组的按位或的结果，以及值等于该结果的子数组的个数。
2. 求按位或结果等于**任意给定数字**的子数组的最短长度/最长长度。

末尾列出了一些题目，均可以用该模板秒杀。

思考：对于起始位置为 $i$ 的子数组的按位或，至多有多少种不同的结果？

根据或运算的性质，我们可以从 $x=\textit{nums}[i]$ 开始，不断往右扩展子数组，按位或的结果要么使 $x$ 不变，要么让 $x$ 的某些比特位的值由 $0$ 变 $1$。最坏情况下从 $x=0$ 出发，每次改变一个比特位，最终得到 $2^{29}-1<10^9$，因此至多有 $30$ 种不同的结果。

另一个结论是，相同的按位或对应的子数组右端点会形成一个连续的区间，这可以用来统计按位或结果及其对应的子数组的个数。

据此，我们可以倒着遍历 $\textit{nums}$，在遍历的同时，用一个数组 $\textit{ors}$ 维护从 $\textit{nums}[i]$ 开始的子数组的按位或的结果，及其对应的子数组右端点的最小值。继续遍历到 $\textit{nums}[i-1]$ 时，我们可以把 $\textit{nums}[i-1]$ 和 $\textit{ors}$ 中的每个值按位或，合并值相同的结果。

这样在遍历时，$\textit{ors}$ 中值最大的元素对应的子数组右端点的最小值，就是要求的最短子数组的右端点。

注：下面代码用到了**原地去重**的技巧，如果你对此并不熟悉，可以先做做 [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)。

```py [sol1-Python3]
class Solution:
    def smallestSubarrays(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n
        ors = []  # 按位或的值 + 对应子数组的右端点的最小值
        for i in range(n - 1, -1, -1):
            num = nums[i]
            ors.append([0, i])
            k = 0
            for p in ors:
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

```java [sol1-Java]
class Solution {
    public int[] smallestSubarrays(int[] nums) {
        var n = nums.length;
        var ans = new int[n];
        var ors = new ArrayList<int[]>(); // 按位或的值 + 对应子数组的右端点的最小值
        for (int i = n - 1; i >= 0; --i) {
            ors.add(new int[]{0, i});
            var k = 0;
            for (var or : ors) {
                or[0] |= nums[i];
                if (ors.get(k)[0] == or[0])
                    ors.get(k)[1] = or[1]; // 合并相同值，下标取最小的
                else ors.set(++k, or);
            }
            ors.subList(k + 1, ors.size()).clear();
            // 本题只用到了 ors[0]，如果题目改成任意给定数值，可以在 ors 中查找
            ans[i] = ors.get(0)[1] - i + 1;
        }
        return ans;
    }
}
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
            // 本题只用到了 ors[0]，如果题目改成任意给定数字，可以在 ors 中查找
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
        // 本题只用到了 ors[0]，如果题目改成任意给定数字，可以在 ors 中查找
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

#### 思考题

如果是**异或**要怎么做？

依然是倒序遍历，求后缀异或和，然后可以用 [421. 数组中两个数的最大异或值](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/) 的字典树方法，需要额外存后缀异或和对应的下标，如果有多个相同的，存下标最小的。
