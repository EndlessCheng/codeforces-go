#### 提示 1

思考方向？

看到题目给的数据范围，先想想能否用 DP 做出来。（DP 可以认为是一种更高级的暴力）

#### 提示 2

如何定义 DP 的状态？

一般来说，题目给了什么就用什么定义：地板长度和地毯个数。而地毯长度更适合去划分状态。

只用地板长度一个维度够吗？

不够，状态定义没有体现出所使用的地毯的个数。因此需要两个维度。

#### 提示 3

状态的值及其转移如何设计？

一般来说，题目求什么就定义什么：定义 $f[i][j]$ 表示用 $i$ 条地毯覆盖前 $j$ 块板砖时，没被覆盖的白色砖块的最少数目。

转移时可以考虑是否用第 $i$ 条地毯的末尾覆盖第 $j$ 块板砖：

- 不覆盖：$f[i][j] = f[i][j-1] + [\textit{floor}[j]=\text{`1'}]$；
- 覆盖：$f[i][j] = f[i-1][j-\textit{carpetLen}]$

取二者最小值。

注意 $i=0$ 的时候只能不覆盖，需要单独计算。

最后答案为 $f[\textit{numCarpets}][\textit{floor.length}-1]$。

```Python [sol1-Python3]
class Solution:
    def minimumWhiteTiles(self, floor: str, n: int, carpetLen: int) -> int:
        m = len(floor)
        f = [[0] * m for _ in range(n + 1)]
        f[0][0] = ord(floor[0]) % 2
        for i in range(1, m):
            f[0][i] = f[0][i - 1] + ord(floor[i]) % 2
        for i in range(1, n + 1):
            # j < carpetLen 的 f[i][j] 均为 0
            for j in range(carpetLen, m):
                f[i][j] = min(f[i][j - 1] + ord(floor[j]) % 2, f[i - 1][j - carpetLen])
        return f[n][m - 1]
```

```go [sol1-Go]
func minimumWhiteTiles(floor string, n, carpetLen int) int {
	m := len(floor)
	f := make([][]int, n+1)
	f[0] = make([]int, m)
	f[0][0] = int(floor[0] & 1)
	for i := 1; i < m; i++ {
		f[0][i] = f[0][i-1] + int(floor[i]&1)
	}
	for i := 1; i <= n; i++ {
		f[i] = make([]int, m)
		// j < carpetLen 的 f[i][j] 均为 0
		for j := carpetLen; j < m; j++ {
			f[i][j] = min(f[i][j-1]+int(floor[j]&1), f[i-1][j-carpetLen])
		}
	}
	return f[n][m-1]
}

func min(a, b int) int { if a > b { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int minimumWhiteTiles(string &floor, int n, int carpetLen) { // 默认代码没加引用，这里补上
        int m = floor.length();
        vector<vector<int>> f(n + 1, vector<int>(m));
        f[0][0] = floor[0] % 2;
        for (int i = 1; i < m; ++i)
            f[0][i] = f[0][i - 1] + floor[i] % 2;
        for (int i = 1; i <= n; ++i)
            // j < carpetLen 的 f[i][j] 均为 0
            for (int j = carpetLen; j < m; ++j)
                f[i][j] = min(f[i][j - 1] + floor[j] % 2, f[i - 1][j - carpetLen]);
        return f[n][m - 1];
    }
};
```

```java [sol1-Java]
class Solution {
    public int minimumWhiteTiles(String floor, int n, int carpetLen) {
        var m = floor.length();
        var f = new int[n + 1][m];
        f[0][0] = floor.charAt(0) % 2;
        for (var i = 1; i < m; ++i)
            f[0][i] = f[0][i - 1] + floor.charAt(i) % 2;
        for (var i = 1; i <= n; ++i)
            // j < carpetLen 的 f[i][j] 均为 0
            for (var j = carpetLen; j < m; ++j)
                f[i][j] = Math.min(f[i][j - 1] + floor.charAt(j) % 2, f[i - 1][j - carpetLen]);
        return f[n][m - 1];
    }
}
```
