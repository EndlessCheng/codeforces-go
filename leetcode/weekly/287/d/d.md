#### 提示 1

注意到 $\textit{values}$ 中有相同的字符串，因此解密出的字符串可能不是唯一的。

#### 提示 2

对于解密操作，不妨逆向思考，即加密所有 $\textit{dictionary}$ 中的字符串。用哈希表记录每个加密后的字符串的出现次数。这样每次调用 $\text{decrypt}$ 时，返回哈希表中 $\textit{word}_2$ 的出现次数即可。

```Python [sol1-Python3]
class Encrypter:
    def __init__(self, keys: List[str], values: List[str], dictionary: List[str]):
        self.map = dict(zip(keys, values))
        self.cnt = Counter(self.encrypt(s) for s in dictionary)

    def encrypt(self, word1: str) -> str:
        res = ""
        for ch in word1:
            if ch not in self.map:
                return ""
            res += self.map[ch]
        return res

    def decrypt(self, word2: str) -> int:
        return self.cnt[word2]
```

```go [sol1-Go]
type Encrypter struct {
	mp  [26]string
	cnt map[string]int
}

func Constructor(keys []byte, values, dictionary []string) Encrypter {
	mp := [26]string{}
	for i, key := range keys {
		mp[key-'a'] = values[i]
	}
	e := Encrypter{mp, map[string]int{}}
	for _, s := range dictionary {
		e.cnt[e.Encrypt(s)]++
	}
	return e
}

func (e *Encrypter) Encrypt(word1 string) string {
	res := make([]byte, 0, len(word1)*2)
	for _, ch := range word1 {
		s := e.mp[ch-'a']
		if s == "" { return "" }
		res = append(res, s...)
	}
	return string(res)
}

func (e *Encrypter) Decrypt(word2 string) int { return e.cnt[word2] }
```

```C++ [sol1-C++]
class Encrypter {
    array<string, 26> mp;
    unordered_map<string, int> cnt;
public:
    Encrypter(vector<char> &keys, vector<string> &values, vector<string> &dictionary) {
        for (int i = 0; i < keys.size(); ++i)
            mp[keys[i] - 'a'] = values[i];
        for (auto &s : dictionary)
            ++cnt[encrypt(s)];
    }

    string encrypt(string word1) {
        string res;
        for (char ch : word1) {
            auto &s = mp[ch - 'a'];
            if (s == "") return "";
            res += s;
        }
        return res;
    }

    int decrypt(string word2) { return cnt.count(word2) ? cnt[word2] : 0; } // 防止把不在 cnt 中的字符串加进去
};
```

```java [sol1-Java]
class Encrypter {
    String[] map = new String[26];
    Map<String, Integer> cnt = new HashMap<>();

    public Encrypter(char[] keys, String[] values, String[] dictionary) {
        map = new String[26];
        for (var i = 0; i < keys.length; i++)
            map[keys[i] - 'a'] = values[i];
        for (var s : dictionary) {
            var e = encrypt(s);
            cnt.put(e, cnt.getOrDefault(e, 0) + 1);
        }
    }

    public String encrypt(String word1) {
        var res = new StringBuilder();
        for (var i = 0; i < word1.length(); i++) {
            var s = map[word1.charAt(i) - 'a'];
            if (s == null) return "";
            res.append(s);
        }
        return res.toString();
    }

    public int decrypt(String word2) {
        return cnt.getOrDefault(word2, 0);
    }
}
```

时间复杂度：均为线性（与输入成正比）。
