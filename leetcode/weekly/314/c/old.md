#### 提示 1

$t$ 是一个栈。

问题相当于从左到右遍历 $s$，在允许用一个辅助栈的前提下，计算能得到的字典序最小的字符串。

#### 提示 2

贪心地思考，为了让字典序最小，在遍历 $s$ 的过程中，如果栈顶字符 $\le$ 后续字符（未入栈）的最小值，那么应该出栈并加到答案末尾，否则应当继续遍历，取到比栈顶字符小的那个字符，这样才能保证字典序最小。

#### 提示 3

代码实现时，为了快速判断剩余字符的最小值，我们可以先统计 $s$ 每个字符的出现次数 $\textit{cnt}$，然后在遍历 $s$ 的过程中更新 $\textit{cnt}$，这样 $\textit{cnt}$ 中第一个正数对应的字符就是剩余字符中最小的。

```py [sol1-Python3]
class Solution:
    def robotWithString(self, s: str) -> str:
        ans = []
        cnt = Counter(s)
        min = 0  # 剩余最小字母
        st = []
        for c in s:
            cnt[c] -= 1
            while min < 25 and cnt[ascii_lowercase[min]] == 0:
                min += 1
            st.append(c)
            while st and st[-1] <= ascii_lowercase[min]:
                ans.append(st.pop())
        return ''.join(ans)
```

```java [sol1-Java]
class Solution {
    public String robotWithString(String S) {
        var ans = new StringBuilder();
        var s = S.toCharArray();
        var cnt = new int[26];
        for (var c : s) ++cnt[c - 'a'];
        var min = 0; // 剩余最小字母
        var st = new ArrayDeque<Character>();
        for (var c : s) {
            --cnt[c - 'a'];
            while (min < 25 && cnt[min] == 0) ++min;
            st.push(c);
            while (!st.isEmpty() && st.peek() - 'a' <= min)
                ans.append(st.poll());
        }
        return ans.toString();
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    string robotWithString(string s) {
        string ans;
        int cnt[26]{}, min = 0; // min 表示剩余最小字母
        for (char c : s) ++cnt[c - 'a'];
        stack<char> st;
        for (char c : s) {
            --cnt[c - 'a'];
            while (min < 25 && cnt[min] == 0) ++min;
            st.push(c);
            while (!st.empty() && st.top() - 'a' <= min) {
                ans += st.top();
                st.pop();
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func robotWithString(s string) string {
	ans := make([]byte, 0, len(s))
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	min := byte(0) // 剩余最小字母
	st := []byte{}
	for _, c := range s {
		cnt[c-'a']--
		for min < 25 && cnt[min] == 0 { // 找到第一个正数
			min++
		}
		st = append(st, byte(c))
		for len(st) > 0 && st[len(st)-1]-'a' <= min {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$O(n+|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。注意到每个字母只会入栈出栈各一次，且 $\textit{min}$ 只增不减，因此时间复杂度为 $O(n+|\Sigma|)$。
- 空间复杂度：$O(n+|\Sigma|)$。最坏情况下栈需要 $O(n)$ 的空间。

---

欢迎关注我的B站频道：[灵茶山艾府](https://space.bilibili.com/206214)，定期更新算法讲解视频哦~
