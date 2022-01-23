两种做法：倒序 DP / 正序 DP

#### 解法一：倒序 DP（填表法）

填表法适用于大多数 DP：通过当前状态所依赖的状态，来计算当前状态。

设有 $n$ 个问题，定义 $f[i]$ 表示解决区间 $[i,n-1]$ 内的问题可以获得的最高分数。

倒序遍历问题列表，对于第 $i$ 个问题，我们有两种决策：跳过或解决。

若跳过，则有 $f[i]=f[i+1]$。

若解决，则需要跳过后续 $\textit{brainpower}[i]$ 个问题。记 $j=i+\textit{brainpower}[i]+1$，则有

$$
f[i] =
\begin{cases} 
\textit{point}[i]+f[j],&j<n\\
\textit{point}[i],&j\ge n
\end{cases}
$$

这两种决策取最大值。

最后答案为 $f[0]$。

```go [sol1-Go]
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		q := questions[i]
		if j := i + q[1] + 1; j < n {
			f[i] = max(f[i+1], q[0]+f[j])
		} else {
			f[i] = max(f[i+1], q[0])
		}
	}
	return int64(f[0])
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    long long mostPoints(vector<vector<int>> &questions) {
        int n = questions.size();
        vector<long long> f(n + 1);
        for (int i = n - 1; i >= 0; --i) {
            auto &q = questions[i];
            int j = i + q[1] + 1;
            f[i] = max(f[i + 1], q[0] + (j < n ? f[j] : 0));
        }
        return f[0];
    }
};
```

```Python [sol1-Python3]
class Solution:
    def mostPoints(self, questions: List[List[int]]) -> int:
        n = len(questions)
        f = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            q = questions[i]
            j = i + q[1] + 1
            f[i] = max(f[i + 1], q[0] + (f[j] if j < n else 0))
        return f[0]
```

```java [sol1-Java]
class Solution {
    public long mostPoints(int[][] questions) {
        var n = questions.length;
        var f = new long[n + 1];
        for (var i = n - 1; i >= 0; --i) {
            var q = questions[i];
            var j = i + q[1] + 1;
            f[i] = Math.max(f[i + 1], q[0] + (j < n ? f[j] : 0));
        }
        return f[0];
    }
}
```

#### 解法二：正序 DP（刷表法）

另一种做法是刷表法：用当前状态，去更新当前状态所影响的状态。

定义 $f[i]$ 表示解决区间 $[0,i)$ 内的问题可以获得的最高分数。

对于问题 $i$，若跳过，则可以更新 $f[i+1]=\max(f[i+1],f[i])$。

若不跳过，记 $j=i+\textit{brainpower}[i]+1$，则可以更新 $f[j]=\max(f[j],f[i]+\textit{point}[i])$。

对于 $i=n-1$ 和 $j\ge n$ 的情况，为了简化代码逻辑，我们可以将其更新到 $f[n]$ 中。

最后答案为 $f[n]$。

```go [sol2-Go]
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n+1)
	for i, q := range questions {
		f[i+1] = max(f[i+1], f[i])
		j := i + q[1] + 1
		if j > n {
			j = n
		}
		f[j] = max(f[j], f[i]+q[0])
	}
	return int64(f[n])
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol2-C++]
class Solution {
public:
    long long mostPoints(vector<vector<int>> &questions) {
        int n = questions.size();
        vector<long long> f(n + 1);
        for (int i = 0; i < n; ++i) {
            f[i + 1] = max(f[i + 1], f[i]);
            auto &q = questions[i];
            int j = min(i + q[1] + 1, n);
            f[j] = max(f[j], f[i] + q[0]);
        }
        return f[n];
    }
};
```

```Python [sol2-Python3]
class Solution:
    def mostPoints(self, questions: List[List[int]]) -> int:
        n = len(questions)
        f = [0] * (n + 1)
        for i, q in enumerate(questions):
            f[i + 1] = max(f[i + 1], f[i])
            j = min(i + q[1] + 1, n)
            f[j] = max(f[j], f[i] + q[0])
        return f[n]
```

```java [sol2-Java]
class Solution {
    public long mostPoints(int[][] questions) {
        var n = questions.length;
        var f = new long[n + 1];
        for (var i = 0; i < n; i++) {
            f[i + 1] = Math.max(f[i + 1], f[i]);
            var q = questions[i];
            var j = Math.min(i + q[1] + 1, n);
            f[j] = Math.max(f[j], f[i] + q[0]);
        }
        return f[n];
    }
}
```
