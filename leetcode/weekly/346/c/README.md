## 视频讲解

见[【周赛 346】](https://www.bilibili.com/video/BV1Qm4y1t7cx/)第三题，欢迎点赞投币！

## 思路

判断 $[1,1000]$ 的每个数字 $i$ 是否符合要求，并预处理 $[1,i]$ 内的符合要求的数字和 $\textit{preSum}$。

对于每个数字 $i$，把它转成字符串 $s$ 后，写一个回溯，枚举第一个子串、第二个子串、……，累加所有子串对应的整数值之和 $\textit{sum}$。如果存在 $\textit{sum}=i$，则说明 $i$ 符合要求。

如果你不清楚怎么写这个回溯，可以看 [回溯算法套路①子集型回溯【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。

```python [sol-Python3]
PRE_SUM = [0] * 1001  # 预处理
for i in range(1, 1001):
    s = str(i * i)
    n = len(s)
    def dfs(p: int, sum: int) -> bool:
        if p == n:  # 递归终点
            return sum == i  # i 符合要求
        x = 0
        for j in range(p, n):  # 从 s[p] 到 s[j] 组成的子串
            x = x * 10 + int(s[j])  # 对应的整数值
            if dfs(j + 1, sum + x):
                return True
        return False
    PRE_SUM[i] = PRE_SUM[i - 1] + (i * i if dfs(0, 0) else 0)

class Solution:
    def punishmentNumber(self, n: int) -> int:
        return PRE_SUM[n]
```

```java [sol-Java]
class Solution {
    private static final int[] PRE_SUM = new int[1001];

    static {
        for (int i = 1; i <= 1000; i++) {
            var s = Integer.toString(i * i).toCharArray();
            PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(s, i, 0, 0) ? i * i : 0);
        }
    }

    private static boolean dfs(char[] s, int i, int p, int sum) {
        if (p == s.length)  // 递归终点
            return sum == i; // i 符合要求
        int x = 0;
        for (int j = p; j < s.length; j++) { // 从 s[p] 到 s[j] 组成的子串
            x = x * 10 + s[j] - '0'; // 对应的整数值
            if (dfs(s, i, j + 1, sum + x))
                return true;
        }
        return false;
    }

    public int punishmentNumber(int n) {
        return PRE_SUM[n];
    }
}
```

```cpp [sol-C++]
int PRE_SUM[1001];

int init = []() {
    for (int i = 1; i <= 1000; i++) {
        auto s = to_string(i * i);
        int n = s.length();
        function<bool(int, int)> dfs = [&](int p, int sum) -> bool {
            if (p == n)  // 递归终点
                return sum == i; // i 符合要求
            int x = 0;
            for (int j = p; j < n; j++) { // 从 s[p] 到 s[j] 组成的子串
                x = x * 10 + int(s[j] - '0'); // 对应的整数值
                if (dfs(j + 1, sum + x))
                    return true;
            }
            return false;
        };
        PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(0, 0) ? i * i : 0);
    }
    return 0;
}();

class Solution {
public:
    int punishmentNumber(int n) {
        return PRE_SUM[n];
    }
};
```

```go [sol-Go]
var preSum [1001]int

func init() { // 预处理
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n { // 递归终点
				return sum == i // i 符合要求
			}
			x := 0
			for j := p; j < n; j++ { // 从 s[p] 到 s[j] 组成的子串
				x = x*10 + int(s[j]-'0') // 对应的整数值
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(0, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
```

#### 复杂度分析

- 时间复杂度：预处理 $\mathcal{O}(U^{1 + 2\log_{10} 2})\approx\mathcal{O}(U^{1.602})$，其中 $U=1000$。对于数字 $i^2$，它转成字符串后的长度为 $m=\lfloor1+2\log_{10} i\rfloor$，所以回溯需要 $\mathcal{O}(2^m)=\mathcal{O}(i^{2\log_{10} 2})$ 的时间，对其积分可知，整个预处理需要 $\mathcal{O}(U^{1 + 2\log_{10} 2})$ 的时间。
- 空间复杂度：预处理 $\mathcal{O}(U)$。
