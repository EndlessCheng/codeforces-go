下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

设分割个数为 $i$，它能「容纳」多长的 $\textit{message}$？

核心思路：设 $\textit{message}$ 的长度为 $n$。枚举分割个数 $i$，不断增大容量 $\textit{cap}$，直到 $\textit{cap} \ge n$ 为止，就可以分割了。

```py [sol1-Python3]
class Solution:
    def splitMessage(self, message: str, limit: int) -> List[str]:
        i = cap = 0
        while True:
            i += 1
            if i < 10:
                tail_len = 5  # 结尾的长度
            elif i < 100:
                if i == 10: cap -= 9  # 前面的结尾的长度都 +1，那么容量就要减小
                tail_len = 7
            elif i < 1000:
                if i == 100: cap -= 99
                tail_len = 9
            else:
                if i == 1000: cap -= 999
                tail_len = 11
            if tail_len >= limit: return []  # cap 无法增大，寄
            cap += limit - tail_len
            if cap < len(message): continue  # 容量没有达到，继续枚举

            ans, k = [], 0
            for j in range(1, i + 1):
                tail = f"<{j}/{i}>"
                if j == i:
                    ans.append(message[k:] + tail)
                else:
                    m = limit - len(tail)
                    ans.append(message[k: k + m] + tail)
                    k += m
            return ans
```

```java [sol1-Java]
class Solution {
    public String[] splitMessage(String message, int limit) {
        var n = message.length();
        for (int i = 1, cap = 0, tail_len; ; ++i) {
            if (i < 10) tail_len = 5; // 结尾的长度
            else if (i < 100) {
                if (i == 10) cap -= 9; // 前面的结尾的长度都 +1，那么容量就要减小
                tail_len = 7;
            } else if (i < 1000) {
                if (i == 100) cap -= 99;
                tail_len = 9;
            } else {
                if (i == 1000) cap -= 999;
                tail_len = 11;
            }
            if (tail_len >= limit) return new String[]{}; // cap 无法增大，寄
            cap += limit - tail_len;
            if (cap < n) continue; // 容量没有达到，继续枚举

            var ans = new String[i];
            for (int j = 0, k = 0; j < i; ++j) {
                var tail = "<" + (j + 1) + "/" + i + ">";
                if (j == i - 1) ans[j] = message.substring(k) + tail;
                else {
                    var m = limit - tail.length();
                    ans[j] = message.substring(k, k + m) + tail;
                    k += m;
                }
            }
            return ans;
        }
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<string> splitMessage(string message, int limit) {
        int n = message.length();
        for (int i = 1, cap = 0, tail_len;; ++i) {
            if (i < 10) tail_len = 5; // 结尾的长度
            else if (i < 100) {
                if (i == 10) cap -= 9; // 前面的结尾的长度都 +1，那么容量就要减小
                tail_len = 7;
            } else if (i < 1000) {
                if (i == 100) cap -= 99;
                tail_len = 9;
            } else {
                if (i == 1000) cap -= 999;
                tail_len = 11;
            }
            if (tail_len >= limit) return {}; // cap 无法增大，寄
            cap += limit - tail_len;
            if (cap < n) continue; // 容量没有达到，继续枚举

            vector<string> ans(i);
            for (int j = 0, k = 0; j < i; ++j) {
                string tail = "<" + to_string(j + 1) + "/" + to_string(i) + ">";
                if (j == i - 1) ans[j] = message.substr(k) + tail;
                else {
                    int m = limit - tail.length();
                    ans[j] = message.substr(k, m) + tail;
                    k += m;
                }
            }
            return ans;
        }
    }
};
```

```go [sol1-Go]
func splitMessage(message string, limit int) []string {
	for i, cap, tailLen := 1, 0, 0; ; i++ {
		if i < 10 {
			tailLen = 5 // 结尾的长度
		} else if i < 100 {
			if i == 10 { cap -= 9 } // 前面的结尾的长度都 +1，那么容量就要减小
			tailLen = 7
		} else if i < 1000 {
			if i == 100 { cap -= 99 }
			tailLen = 9
		} else {
			if i == 1000 { cap -= 999 }
			tailLen = 11
		}
		if tailLen >= limit { return nil } // cap 无法增大，寄
		cap += limit - tailLen
		if cap < len(message) { continue } // 容量没有达到，继续枚举

		ans := make([]string, i)
		for j := range ans {
			tail := fmt.Sprintf("<%d/%d>", j+1, i)
			if j == i-1 {
				ans[j] = message + tail
			} else {
				m := limit - len(tail)
				ans[j] = message[:m] + tail
				message = message[m:]
			}
		}
		return ans
	}
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{message}$ 的长度。构造答案需要 $O(n\log n)$ 的时间。
- 空间复杂度：$O(1)$，返回值的空间不计入。
