## 解法一：前后缀分解

#### 提示 1

考虑将列车分成左右两部分，枚举分割线，分别计算这两部分的最少时间。

#### 提示 2

考虑 DP。

![](https://pic.leetcode-cn.com/1644496698-FtlVnq-2167-1.drawio%20\(1\).png)

考虑左半部分的最少时间。

定义 $\textit{pre}[i]$ 表示移除从 $s[0]$ 到 $s[i]$ 的所有违禁货物车厢所花费的最少时间。

讨论 $s[i]$：

- 当 $s[i]=0$ 时，无需移除车厢，则有 $\textit{pre}[i]=\textit{pre}[i-1]$；
- 当 $s[i]=1$ 时，可以单独移除第 $i$ 节车厢，也可以移除前 $i$ 个车厢，二者取最小值，即 $\textit{pre}[i]=\min(\textit{pre}[i-1]+2,i+1)$。

对于右半部分，同样定义 $\textit{suf}[i]$ 表示移除从 $s[i]$ 到 $s[n-1]$ 的所有违禁货物车厢所花费的最少时间，有

$$
\textit{suf}[i] =
\begin{cases} 
\textit{suf}[i+1],&s[i]=0\\
\min(\textit{suf}[i+1]+2,n-i),&s[i]=1
\end{cases}
$$

然后枚举分割线，计算所有 $\textit{pre}[i]+\textit{suf}[i+1]$ 的最小值，即为答案。

---

代码实现时，有如下三处优化点：

优化 1：可以先计算 $\textit{suf}$，然后在枚举分割线的同时计算 $\textit{pre}$。

优化 2：由于计算 $\textit{pre}$ 的转移时当前状态只和上一个状态有关，因此可以使用滚动数组优化，即用一个变量来表示 $\textit{pre}$。

优化 3：由于 $s[i]=0$ 时，$\textit{pre}[i]$ 和 $\textit{suf}[i]$ 的值均不会变化，因此仅需要考虑 $s[i]=1$ 时的 $\textit{pre}[i]+\textit{suf}[i+1]$ 的最小值。

```py [sol1-Python3]
class Solution:
    def minimumTime(self, s: str) -> int:
        n = len(s)
        suf = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            suf[i] = suf[i + 1] if s[i] == '0' else min(suf[i + 1] + 2, n - i)
        ans = suf[0]
        pre = 0
        for i, ch in enumerate(s):
            if ch == '1':
                pre = min(pre + 2, i + 1)
                ans = min(ans, pre + suf[i + 1])
        return ans
```

```C++ [sol1-C++]
class Solution {
public:
    int minimumTime(string s) {
        int n = s.length();
        vector<int> suf(n + 1);
        for (int i = n - 1; i >= 0; --i)
            suf[i] = s[i] == '0' ? suf[i + 1] : min(suf[i + 1] + 2, n - i);
        int ans = suf[0], pre = 0;
        for (int i = 0; i < n; ++i)
            if (s[i] == '1') {
                pre = min(pre + 2, i + 1);
                ans = min(ans, pre + suf[i + 1]);
            }
        return ans;
    }
};
```

```java [sol1-Java]
class Solution {
    public int minimumTime(String s) {
        var n = s.length();
        var suf = new int[n + 1];
        for (var i = n - 1; i >= 0; --i)
            suf[i] = s.charAt(i) == '0' ? suf[i + 1] : Math.min(suf[i + 1] + 2, n - i);
        var ans = suf[0];
        var pre = 0;
        for (var i = 0; i < n; ++i)
            if (s.charAt(i) == '1') {
                pre = Math.min(pre + 2, i + 1);
                ans = Math.min(ans, pre + suf[i + 1]);
            }
        return ans;
    }
}
```

```go [sol1-Go]
func minimumTime(s string) int {
	n := len(s)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			suf[i] = suf[i+1]
		} else {
			suf[i] = min(suf[i+1]+2, n-i)
		}
	}
	ans := suf[0]
	pre := 0
	for i, ch := range s {
		if ch == '1' {
			pre = min(pre+2, i+1)
			ans = min(ans, pre+suf[i+1])
		}
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 解法二：进一步优化，一次遍历

由于我们计算的是「移除前缀 + 移除分割线左侧某些车厢 + (分割线) + 移除分割线右侧某些车厢 + 移除后缀」的最少花费，其中「移除分割线左侧某些车厢 + 移除分割线右侧某些车厢」都是在移除中间的某些车厢，因此这是可以合并的，不妨合并到分割线左侧，即计算「移除前缀 + 移除分割线左侧某些车厢 + (分割线) + 移除后缀」的最少花费。

![](https://pic.leetcode-cn.com/1644496724-UEaqty-2167-2.drawio.png)

合并后，计算 $\textit{pre}$ 的过程不变，而 $\textit{suf}$ 就仅为移除后缀所有车厢的花费了，这可以直接用下标计算出来。因此我们可以省略 $\textit{suf}$ 的计算流程，直接一次遍历计算出答案。

```py [sol2-Python3]
class Solution:
    def minimumTime(self, s: str) -> int:
        ans = n = len(s)
        pre = 0
        for i, ch in enumerate(s):
            if ch == '1':
                pre = min(pre + 2, i + 1)
            ans = min(ans, pre + n - 1 - i)
        return ans
```

```C++ [sol2-C++]
class Solution {
public:
    int minimumTime(string s) {
        int n = s.length(), ans = n, pre = 0;
        for (int i = 0; i < n; ++i) {
            if (s[i] == '1') pre = min(pre + 2, i + 1);
            ans = min(ans, pre + n - 1 - i);
        }
        return ans;
    }
};
```



```java [sol2-Java]
class Solution {
    public int minimumTime(String s) {
        var n = s.length();
        var ans = n;
        var pre = 0;
        for (var i = 0; i < n; ++i) {
            if (s.charAt(i) == '1') pre = Math.min(pre + 2, i + 1);
            ans = Math.min(ans, pre + n - 1 - i);
        }
        return ans;
    }
}
```

```go [sol2-Go]
func minimumTime(s string) int {
	n := len(s)
	ans := n
	pre := 0
	for i, ch := range s {
		if ch == '1' {
			pre = min(pre+2, i+1)
		}
		ans = min(ans, pre+n-1-i)
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
