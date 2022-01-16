两种解法：二分答案 / 排序+贪心（附证明）

#### 解法一：二分答案

假设我们可以让 $n$ 台电脑同时运行 $x$ 分钟，那么对于电量大于 $x$ 的电池，其只能被使用 $x$ 分钟。因此每个电池的使用时间为 $\min(\textit{batteries}[i], x)$，我们将其累加起来，记作 $\textit{sum}$。那么要让 $n$ 台电脑同时运行 $x$ 分钟，必要条件是 $n\cdot x\le \textit{sum}$。

下面证明该条件是充分的，即当 $n\cdot x\le \textit{sum}$ 成立时，必然可以让 $n$ 台电脑同时运行 $x$ 分钟。

对于电量不小于 $x$ 的电池，我们可以让其给一台电脑供电 $x$ 分钟。由于一个电池不能同时给多台电脑供电，因此该电池若给一台电脑供电 $x$ 分钟，那它就不能用于其他电脑了。我们可以将所有电量不小于 $x$ 的电池各给一台电脑供电。

对于其余的电池，设其电量和为 $\textit{sum}'$，剩余 $n'$ 台电脑未被供电。我们可以随意选择剩下的电池，供给剩余的第一台电脑，多余的电池电量供给剩余的第二台电脑，依此类推。注意由于这些电池的电量小于 $x$，按照这种做法是不会出现同一个电池在同一时间供给多台电脑的。

由于 $\textit{sum}'=\textit{sum}-(n-n')\cdot x$，结合 $n\cdot x\le \textit{sum}$ 可以得到 $n'\cdot x\le \textit{sum}'$，这意味着剩余电池可以让剩余电脑运行 $x$ 分钟。充分性得证。

如果我们可以让 $n$ 台电脑同时运行 $x$ 分钟，那么必然也可以同时运行低于 $x$ 分钟，因此答案满足单调性，可以二分答案，通过判断 $n\cdot x\le \textit{sum}$ 来求解。

```go [sol1-Go]
func maxRunTime(n int, batteries []int) int64 {
	tot := 0
	for _, b := range batteries {
		tot += b
	}
	return int64(sort.Search(tot/n, func(x int) bool {
		x++
		sum := 0
		for _, b := range batteries {
			sum += min(b, x)
		}
		return sum/n < x
	}))
}

func min(a, b int) int { if a > b { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    long long maxRunTime(int n, vector<int> &batteries) {
        long tot = accumulate(batteries.begin(), batteries.end(), 0L);
        long l = 1, r = tot / n + 1;
        while (l < r) {
            long x = (l + r) / 2, sum = 0;
            for (long b : batteries) {
                sum += min(b, x);
            }
            if (sum / n >= x) {
                l = x + 1;
            } else {
                r = x;
            }
        }
        return r - 1;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        tot = sum(batteries)
        l, r = 1, tot // n + 1
        while l < r:
            x = (l + r) // 2
            if n * x <= sum(min(b, x) for b in batteries):
                l = x + 1
            else:
                r = x
        return r - 1
```

#### 解法二：排序 + 贪心

受解法一的启发，我们可以得出如下贪心策略：

记电池电量和为 $\textit{sum}$，则至多可以供电 $x=\lfloor\dfrac{\textit{sum}}{n}\rfloor$ 分钟。我们对电池电量从大到小排序，然后从电量最大的电池开始遍历：

- 若该电池电量超过 $x$，则将其供给一台电脑，问题缩减为 $n-1$ 台电脑的子问题。
- 若该电池电量不超过 $x$，则其余电池的电量均不超过 $x$，此时有

   $$
   n\cdot x=n\cdot\lfloor\dfrac{\textit{sum}}{n}\rfloor \le \textit{sum}
   $$
   
   根据解法一的结论，这些电池可以供电 $x$ 分钟。

由于随着问题规模减小，$x$ 不会增加，因此若遍历到一个电量不超过 $x$ 的电池时，可以直接返回 $x$ 作为答案。
   
```go [sol2-Go]
func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	sum := 0
	for _, b := range batteries {
		sum += b
	}
	for i := len(batteries) - 1; ; i-- {
		if batteries[i] <= sum/n {
			return int64(sum / n)
		}
		sum -= batteries[i]
		n--
	}
}
```

```C++ [sol2-C++]
class Solution {
public:
    long long maxRunTime(int n, vector<int> &batteries) {
        sort(batteries.begin(), batteries.end());
        long sum = accumulate(batteries.begin(), batteries.end(), 0L);
        for (int i = batteries.size() - 1;; --i) {
            if (batteries[i] <= sum / n) {
                return sum / n;
            }
            sum -= batteries[i];
            n--;
        }
    }
};
```

```Python [sol2-Python3]
class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        batteries.sort(reverse=True)
        tot = sum(batteries)
        for i in range(len(batteries)):
            if batteries[i] <= tot // n:
                return tot // n
            tot -= batteries[i]
            n -= 1
```
