[视频讲解](https://www.bilibili.com/video/BV1FV4y1F7v7/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

如果直接计算好分区的数目，我们可以用 **01 背包**来做，但是背包容量太大，会超时。

**正难则反**，我们可以反过来，计算**坏分区**的数目，即第一个组或第二个组的元素和小于 $k$ 的方案数。根据对称性，我们只需要计算第一个组的元素和小于 $k$ 的方案数，然后乘 $2$ 即可。

> 注意，如果 $\textit{nums}$ 的所有元素之和小于 $2k$，则不存在好分区，我们可以特判这种情况，直接返回 $0$。如果不特判，计算坏分区会重复统计，导致错误的结果。

那么原问题就转换为「从 $\textit{nums}$ 中选择若干元素，使得元素和小于 $k$ 的方案数」，这样用 01 背包就不会超时了。

具体来说，定义 $f[i][j]$ 表示从前 $i$ 个数中选择若干元素，和为 $j$ 的方案数。

分类讨论：

- 不选第 $i$ 个数：$f[i][j] = f[i-1][j]$；
- 选第 $i$ 个数：$f[i][j] = f[i-1][j-\textit{nums}[i]]$。

因此 $f[i][j] = f[i-1][j] + f[i-1][j-\textit{nums}[i]]$。

初始值 $f[0][0] = 1$。

坏分区的数目 $\textit{bad} =(f[n][0]+f[n][1]+\cdots+f[n][k-1])\cdot 2$。

答案为所有分区的数目减去坏分区的数目，即 $2^n-\textit{bad}$，这里 $n$ 为 $\textit{nums}$ 的长度。

代码实现时，可以用倒序循环的技巧来压缩空间。

```py [sol1-Python3]
class Solution:
    def countPartitions(self, nums: List[int], k: int) -> int:
        if sum(nums) < k * 2: return 0
        MOD = 10 ** 9 + 7
        f = [0] * k
        f[0] = 1
        for x in nums:
            for j in range(k - 1, x - 1, -1):
                f[j] = (f[j] + f[j - x]) % MOD
        return (pow(2, len(nums), MOD) - sum(f) * 2) % MOD
```

```java [sol1-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;

    public int countPartitions(int[] nums, int k) {
        var sum = 0L;
        for (var x : nums) sum += x;
        if (sum < k * 2) return 0;
        var ans = 1;
        var f = new int[k];
        f[0] = 1;
        for (var x : nums) {
            ans = ans * 2 % MOD;
            for (var j = k - 1; j >= x; --j)
                f[j] = (f[j] + f[j - x]) % MOD;
        }
        for (var x : f)
            ans = (ans - x * 2 % MOD + MOD) % MOD; // 保证答案非负
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;
public:
    int countPartitions(vector<int> &nums, int k) {
        if (accumulate(nums.begin(), nums.end(), 0L) < k * 2) return 0;
        int ans = 1, f[k]; memset(f, 0, sizeof(f));
        f[0] = 1;
        for (int x : nums) {
            ans = ans * 2 % MOD;
            for (int j = k - 1; j >= x; --j)
                f[j] = (f[j] + f[j - x]) % MOD;
        }
        for (int x : f)
            ans = (ans - x * 2 % MOD + MOD) % MOD; // 保证答案非负
        return ans;
    }
};
```

```go [sol1-Go]
func countPartitions(nums []int, k int) int {
	const mod int = 1e9 + 7
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum < k*2 {
		return 0
	}
	ans := 1
	f := make([]int, k)
	f[0] = 1
	for _, x := range nums {
		ans = ans * 2 % mod
		for j := k - 1; j >= x; j-- {
			f[j] = (f[j] + f[j-x]) % mod
		}
	}
	for _, x := range f {
		ans -= x * 2
	}
	return (ans%mod + mod) % mod // 保证答案非负
}
```

#### 复杂度分析

- 时间复杂度：$O(nk)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(k)$。

#### 相似题目

- [494. 目标和](https://leetcode.cn/problems/target-sum/)
