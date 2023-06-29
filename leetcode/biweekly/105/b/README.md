## 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

> APP 用户需要分享到微信打开视频链接。

## 一、寻找子问题

> 为了方便转成递推，从后往前考虑。

设 $n$ 为 $s$ 的长度。我们可以：

- 直接跳过 $s$ 的最后一个字符，那么问题变成 $s$ 的前 $n-1$ 个字符的子问题。
- 考虑「枚举选哪个」，如果从 $s[j]$ 开始的后缀在 $\textit{dictionary}$ 中，那么问题变成 $s$ 的前 $j-1$ 个字符的子问题。

## 二、记忆化搜索

根据上面的讨论，定义 $\textit{dfs}(i)$ 表示 $s$ 的前 $i$ 个字符的子问题。

- 跳过 $s$ 的最后一个字符，有 $\textit{dfs}(i)=\textit{dfs}(i-1)+1$。
- 考虑「枚举选哪个」，如果从 $s[j]$ 到 $s[i]$ 的子串在 $\textit{dictionary}$ 中，有

$$
\textit{dfs}(i)=\min_{j=0}^{i}\textit{dfs}(j-1)
$$

这两种情况取最小值。

递归边界：$\textit{dfs}(-1)=0$。

答案：$\textit{dfs}(n-1)$。

```py [sol-Python3]
class Solution:
    def minExtraChar(self, s: str, dictionary: List[str]) -> int:
        d = set(dictionary)
        @cache
        def dfs(i: int) -> int:
            if i < 0: return 0
            res = dfs(i - 1) + 1  # 不选
            for j in range(i + 1):  # 枚举选哪个
                if s[j:i + 1] in d:
                    res = min(res, dfs(j - 1))
            return res
        return dfs(len(s) - 1)
```

```go [sol-Go]
func minExtraChar(s string, dictionary []string) int {
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		// 不选
		res := dfs(i-1) + 1

		// 枚举选哪个
		for j := 0; j <= i; j++ {
			if has[s[j:i+1]] {
				res = min(res, dfs(j-1))
			}
		}

		*p = res
		return res
	}
	return dfs(n - 1)
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n^3)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。预处理哈希表需要 $\mathcal{O}(L)$ 的时间。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n^2)$，因此时间复杂度为 $\mathcal{O}(n^3)$。所以总的时间复杂度为 $\mathcal{O}(L + n^3)$。
- 空间复杂度：$\mathcal{O}(n+L)$。

## 三、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

做法：

- $\textit{dfs}$ 改成 $f$ 数组；
- 递归改成循环（每个参数都对应一层循环）；
- 递归边界改成 $f$ 数组的初始值。

具体来说，$f[i]$ 的含义和 $\textit{dfs}(i)$ 的含义是一样的，都表示 $s$ 的前 $i$ 个字符的子问题。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 的一致：

- 跳过 $s$ 的最后一个字符，有 $f[i]=f[i-1]+1$。
- 考虑「枚举选哪个」，如果从 $s[j]$ 到 $s[i]$ 的子串在 $\textit{dictionary}$ 中，有

$$
f[i]=\min_{j=0}^{i}f[j-1]
$$

这两种情况取最小值。

但当 $i=0$ 或 $j=0$ 时，等号右边会出现负数下标。或者说，**这种定义方式没有状态能表示递归边界**，即出界的情况。

解决办法：在 $f$ 数组左边添加一个状态来表示 $i=-1$，把原来的 $f[i]$ 改成 $f[i+1]$，$f[j-1]$ 改成 $f[j]$。

相应的递推式为 $f[i+1]=f[i]+1$ 以及 $f[i+1]=\min\limits_{j=0}^{i}f[j]$。

初始值 $f[i]=0$。（翻译自 $\textit{dfs}(-1)=0$。）

答案为 $f[n]$。（翻译自 $\textit{dfs}(n-1)$。）

```py [sol-Python3]
class Solution:
    def minExtraChar(self, s: str, dictionary: List[str]) -> int:
        d = set(dictionary)
        n = len(s)
        f = [0] * (n + 1)
        for i in range(n):
            f[i + 1] = f[i] + 1  # 不选
            for j in range(i + 1):  # 枚举选哪个
                if s[j:i + 1] in d:
                    f[i + 1] = min(f[i + 1], f[j])
        return f[n]
```

```java [sol-Java]
class Solution {
    public int minExtraChar(String s, String[] dictionary) {
        var set = new HashSet<String>(dictionary.length);
        for (var str : dictionary) set.add(str);
        int n = s.length();
        var f = new int[n + 1];
        for (int i = 0; i < n; i++) {
            f[i + 1] = f[i] + 1; // 不选
            for (int j = 0; j <= i; j++) { // 枚举选哪个
                if (set.contains(s.substring(j, i + 1))) {
                    f[i + 1] = Math.min(f[i + 1], f[j]);
                }
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minExtraChar(string s, vector<string> &dictionary) {
        unordered_set<string> set(dictionary.begin(), dictionary.end());
        int n = s.size();
        vector<int> f(n + 1);
        for (int i = 0; i < n; i++) {
            f[i + 1] = f[i] + 1; // 不选
            for (int j = 0; j <= i; j++) { // 枚举选哪个
                if (set.count(s.substr(j, i - j + 1))) {
                    f[i + 1] = min(f[i + 1], f[j]);
                }
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func minExtraChar(s string, dictionary []string) int {
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		f[i+1] = f[i] + 1 // 不选
		for j := 0; j <= i; j++ { // 枚举选哪个
			if has[s[j:i+1]] {
				f[i+1] = min(f[i+1], f[j])
			}
		}
	}
	return f[n]
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n^3)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。预处理哈希表需要 $\mathcal{O}(L)$ 的时间。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n^2)$，因此时间复杂度为 $\mathcal{O}(n^3)$。所以总的时间复杂度为 $\mathcal{O}(L + n^3)$。
- 空间复杂度：$\mathcal{O}(n+L)$。
