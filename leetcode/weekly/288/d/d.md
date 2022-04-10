#### 提示 1

将 $\textit{flowers}$ 从小到大排序。

#### 提示 2

贪心：让靠后的 $\textit{flowers}[i]$ 增加至 $\textit{target}$，其余的 $\textit{flowers}$ 的最小值尽量大。

#### 提示 3

枚举 $\textit{flowers}$ 的后缀，让这些花园的花增加至 $\textit{target}$，同时我们需要求出 $\textit{flowers}$ 的最长前缀（设其长为 $x$），满足前缀中的花园的花都能增加至 $\textit{flowers}[x-1]$ 朵。

设在填充后缀之后，剩余 $\textit{leftFlowers}$ 朵花可以种植，且长为 $x$ 的前缀一共有 $\textit{sumFlowers}$ 朵花，那么这些前缀的花的最小值的最大值为

$$
\Big\lfloor\dfrac{\textit{leftFlowers}+\textit{sumFlowers}}{x}\Big\rfloor
$$

注意这个值不能超过 $\textit{target}-1$，否则不满足题目「不完善」的要求。

按照上述方法，枚举后缀的同时计算出对应的总美丽值，所有总美丽值的最大值即为答案。

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{flowers}$ 的长度。瓶颈在排序上。
- 空间复杂度：$O(1)$，仅用到若干变量。如果考虑快排时的栈开销，则空间复杂度为 $O(\log n)$。

```Python [sol1-Python3]
class Solution:
    def maximumBeauty(self, flowers: List[int], newFlowers: int, target: int, full: int, partial: int) -> int:
        flowers.sort()
        n = len(flowers)
        if flowers[0] >= target:  # 剪枝，此时所有花园都是完善的
            return n * full

        leftFlowers = newFlowers - target * n  # 填充后缀后，剩余可以种植的花
        for i in range(n):
            flowers[i] = min(flowers[i], target)
            leftFlowers += flowers[i]

        ans, x, sumFlowers = 0, 0, 0
        for i in range(n + 1):  # 枚举后缀长度 n-i
            if leftFlowers >= 0:
                # 计算最长前缀的长度
                while x < i and flowers[x] * x - sumFlowers <= leftFlowers:
                    sumFlowers += flowers[x]
                    x += 1  # 注意 x 只增不减，这部分的时间复杂度为 O(n)
                beauty = (n - i) * full  # 计算总美丽值
                if x:
                    beauty += min((leftFlowers + sumFlowers) // x, target - 1) * partial
                ans = max(ans, beauty)
            if i < n:
                leftFlowers += target - flowers[i]
        return ans
```

```go [sol1-Go]
func maximumBeauty(flowers []int, newFlowers int64, target, full, partial int) int64 {
	sort.Ints(flowers)
	n := len(flowers)
	if flowers[0] >= target { // 剪枝，此时所有花园都是完善的
		return int64(n * full)
	}

	leftFlowers := int(newFlowers) - target*n // 填充后缀后，剩余可以种植的花
	for i, f := range flowers {
		flowers[i] = min(f, target)
		leftFlowers += flowers[i]
	}

	ans := 0
	for i, x, sumFlowers := 0, 0, 0; i <= n; i++ { // 枚举后缀长度 n-i
		if leftFlowers >= 0 {
			// 计算最长前缀的长度
			for ; x < i && flowers[x]*x-sumFlowers <= leftFlowers; x++ {
				sumFlowers += flowers[x] // 注意 x 只增不减，这部分的时间复杂度为 O(n)
			}
			beauty := (n - i) * full // 计算总美丽值
			if x > 0 {
				beauty += min((leftFlowers+sumFlowers)/x, target-1) * partial
			}
			ans = max(ans, beauty)
		}
		if i < n {
			leftFlowers += target - flowers[i]
		}
	}
	return int64(ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    long long maximumBeauty(vector<int> &flowers, long long newFlowers, int target, int full, int partial) {
        sort(flowers.begin(), flowers.end());
        long n = flowers.size();
        if (flowers[0] >= target) return n * full; // 剪枝，此时所有花园都是完善的

        long leftFlowers = newFlowers - target * n; // 填充后缀后，剩余可以种植的花
        for (int i = 0; i < n; ++i) {
            flowers[i] = min(flowers[i], target);
            leftFlowers += flowers[i];
        }

        long ans = 0L, sumFlowers = 0L;
        for (int i = 0, x = 0; i <= n; ++i) { // 枚举后缀长度 n-i
            if (leftFlowers >= 0) {
                // 计算最长前缀的长度
                while (x < i && (long) flowers[x] * x - sumFlowers <= leftFlowers)
                    sumFlowers += flowers[x++]; // 注意 x 只增不减，这部分的时间复杂度为 O(n)
                long beauty = (n - i) * full; // 计算总美丽值
                if (x > 0) beauty += min((leftFlowers + sumFlowers) / x, (long) target - 1) * partial;
                ans = max(ans, beauty);
            }
            if (i < n) leftFlowers += target - flowers[i];
        }
        return ans;
    }
};
```

```java [sol1-Java]
class Solution {
    public long maximumBeauty(int[] flowers, long newFlowers, int target, int full, int partial) {
        Arrays.sort(flowers);
        long n = flowers.length;
        if (flowers[0] >= target) return n * full; // 剪枝，此时所有花园都是完善的

        var leftFlowers = newFlowers - target * n;// 填充后缀后，剩余可以种植的花
        for (var i = 0; i < n; ++i) {
            flowers[i] = Math.min(flowers[i], target);
            leftFlowers += flowers[i];
        }

        var ans = 0L;
        var sumFlowers = 0L;
        for (int i = 0, x = 0; i <= n; ++i) { // 枚举后缀长度 n-i
            if (leftFlowers >= 0) {
                // 计算最长前缀的长度
                while (x < i && (long) flowers[x] * x - sumFlowers <= leftFlowers)
                    sumFlowers += flowers[x++]; // 注意 x 只增不减，这部分的时间复杂度为 O(n)
                var beauty = (n - i) * full; // 计算总美丽值
                if (x > 0) beauty += Math.min((leftFlowers + sumFlowers) / x, (long) target - 1) * partial;
                ans = Math.max(ans, beauty);
            }
            if (i < n) leftFlowers += target - flowers[i];
        }
        return ans;
    }
}
```
