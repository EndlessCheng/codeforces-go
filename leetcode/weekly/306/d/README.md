下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

提供一个一般化的数位 DP 模板。

将 $n$ 转换成字符串 $s$，定义 $f(i,\textit{mask}, \textit{isLimit},\textit{isNum})$ 表示从构造 $n$ 从高到低第 $i$ 及其之后位的方案数，其余参数的含义为：

- 要选的数字不能在 $\textit{mask}$ 集合中。
- $\textit{isLimit}$ 表示当前是否收到了 $n$ 的约束。若为真，则第 $i$ 位填入的数字至多为 $s[i]$。
- $\textit{isNum}$ 表示之前是否填了数字。若为假，则当前位可以跳过（不填数字），或者要填入的数字必须至少为 $1$；若为真，则要填入的数字可以从 $0$ 开始。

枚举要填入的数字，具体实现逻辑见代码。

```py [sol1-Python3]
class Solution:
    def countSpecialNumbers(self, n: int) -> int:
        s = str(n)
        @cache
        def f(i: int, mask: int, is_limit: bool, is_num: bool) -> int:
            if i == len(s):
                return int(is_num)
            res = 0
            if not is_num:  # 可以跳过当前数位
                res = f(i + 1, mask, False, False)
            up = int(s[i]) if is_limit else 9
            for d in range(0 if is_num else 1, up + 1):
                if mask >> d & 1 == 0:  # d 不在 mask 中
                    res += f(i + 1, mask | (1 << d), is_limit and d == up, True)
            return res
        return f(0, 0, True, False)
```

```java [sol1-Java]
class Solution {
    char s[];
    int dp[][];

    public int countSpecialNumbers(int n) {
        s = Integer.toString(n).toCharArray();
        var m = s.length;
        dp = new int[m][1 << 10];
        for (var i = 0; i < m; i++) Arrays.fill(dp[i], -1);
        return f(0, 0, true, false);
    }

    int f(int i, int mask, boolean isLimit, boolean isNum) {
        if (i == s.length) return isNum ? 1 : 0;
        if (!isLimit && isNum && dp[i][mask] >= 0) return dp[i][mask];
        var res = 0;
        if (!isNum) res = f(i + 1, mask, false, false); // 可以跳过当前数位
        for (int d = isNum ? 0 : 1, up = isLimit ? s[i] - '0' : 9; d <= up; ++d)
            if ((mask >> d & 1) == 0) // d 不在 mask 中
                res += f(i + 1, mask | (1 << d), isLimit && d == up, true);
        if (!isLimit && isNum) dp[i][mask] = res;
        return res;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int countSpecialNumbers(int n) {
        auto s = to_string(n);
        int m = s.length(), dp[m][1 << 10];
        memset(dp, -1, sizeof(dp));
        function<int(int, int, bool, bool)> f = [&](int i, int mask, bool isLimit, bool isNum) -> int {
            if (i == m) return isNum;
            if (!isLimit && isNum && dp[i][mask] >= 0) return dp[i][mask];
            int res = 0;
            if (!isNum) res = f(i + 1, mask, false, false); // 可以跳过当前数位
            for (int d = 1 - isNum, up = isLimit ? s[i] - '0' : 9; d <= up; ++d)
                if ((mask >> d & 1) == 0) // d 不在 mask 中
                    res += f(i + 1, mask | (1 << d), isLimit && d == up, true);
            if (!isLimit && isNum) dp[i][mask] = res;
            return res;
        };
        return f(0, 0, true, false);
    }
};
```

```go [sol1-Go]
func countSpecialNumbers(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, 1<<10)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, mask int, isUpper, isNum bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 1
		if isNum {
			d = 0
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return f(0, 0, true, false)
}
```
