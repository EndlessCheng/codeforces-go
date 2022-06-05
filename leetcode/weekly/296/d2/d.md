本题 [视频讲解](https://www.bilibili.com/video/BV1w34y1L7yu/) 已出炉，欢迎三连~

---

## 方法一：双向链表

#### 提示 1

注意添加、删除和移动操作均在光标附近完成，且 $\textit{text}$ 的长度和 $k$ 都很小。

#### 提示 2

用链表模拟所有操作，每个节点存储一个字符。

光标指向链表中的一个节点（该节点保存着光标左边的字符）。

#### 提示 3

用一个哨兵节点表示光标位于文本最左侧的情况，此时光标左侧无字符，即指向哨兵节点。

#### 复杂度分析

时空复杂度均与输入量成正比（线性）。

```Python [sol1-Python3]
# 手写双向链表
class Node:
    __slots__ = ('prev', 'next', 'ch')

    def __init__(self, ch=''):
        self.prev = None
        self.next = None
        self.ch = ch

    # 在 self 后插入 node，并返回该 node
    def insert(self, node: 'Node') -> 'Node':
        node.prev = self
        node.next = self.next
        node.prev.next = node
        node.next.prev = node
        return node

    # 从链表中移除 self
    def remove(self) -> None:
        self.prev.next = self.next
        self.next.prev = self.prev

class TextEditor:
    def __init__(self):
        self.root = self.cur = Node()  # 哨兵节点
        self.root.prev = self.root
        self.root.next = self.root  # 初始化双向链表，下面判断节点的 next 若为 self.root，则表示 next 为空

    def addText(self, text: str) -> None:
        for ch in text:
            self.cur = self.cur.insert(Node(ch))

    def deleteText(self, k: int) -> int:
        k0 = k
        while k and self.cur != self.root:
            self.cur = self.cur.prev
            self.cur.next.remove()
            k -= 1
        return k0 - k

    def text(self) -> str:
        s = []
        k, cur = 10, self.cur
        while k and cur != self.root:
            s.append(cur.ch)
            cur = cur.prev
            k -= 1
        return ''.join(reversed(s))

    def cursorLeft(self, k: int) -> str:
        while k and self.cur != self.root:
            self.cur = self.cur.prev
            k -= 1
        return self.text()

    def cursorRight(self, k: int) -> str:
        while k and self.cur.next != self.root:
            self.cur = self.cur.next
            k -= 1
        return self.text()
```

```java [sol1-Java]
class TextEditor {
    Node root, cur;

    public TextEditor() {
        root = cur = new Node(); // 哨兵节点
        root.prev = root;
        root.next = root; // 初始化双向链表，下面判断节点的 next 若为 root，则表示 next 为空
    }

    public void addText(String text) {
        for (var i = 0; i < text.length(); i++)
            cur = cur.insert(new Node(text.charAt(i)));
    }

    public int deleteText(int k) {
        var k0 = k;
        for (; k > 0 && cur != root; --k) {
            cur = cur.prev;
            cur.next.remove();
        }
        return k0 - k;
    }

    String text() {
        var s = new StringBuilder();
        var cur = this.cur;
        for (var k = 10; k > 0 && cur != root; --k) {
            s.append(cur.ch);
            cur = cur.prev;
        }
        return s.reverse().toString();
    }

    public String cursorLeft(int k) {
        for (; k > 0 && cur != root; --k)
            cur = cur.prev;
        return text();
    }

    public String cursorRight(int k) {
        for (; k > 0 && cur.next != root; --k)
            cur = cur.next;
        return text();
    }
}

// 手写双向链表
class Node {
    Node prev, next;
    char ch;

    Node() {}

    Node(char ch) {
        this.ch = ch;
    }

    // 在 this 后插入 node，并返回该 node
    Node insert(Node node) {
        node.prev = this;
        node.next = this.next;
        node.prev.next = node;
        node.next.prev = node;
        return node;
    }

    // 从链表中移除 this
    void remove() {
        this.prev.next = this.next;
        this.next.prev = this.prev;
    }
}
```

```C++ [sol1-C++]
class TextEditor {
    list<char> l;
    list<char>::iterator cur = l.begin();

public:
    TextEditor() {}

    void addText(string text) {
        for (char ch : text)
            l.insert(cur, ch);
    }

    int deleteText(int k) {
        int k0 = k;
        for (; k && cur != l.begin(); --k)
            cur = l.erase(prev(cur));
        return k0 - k;
    }

    string text() {
        string s;
        auto it = cur;
        for (int k = 10; k && it != l.begin(); --k) {
            it = prev(it);
            s += *it;
        }
        reverse(s.begin(), s.end());
        return s;
    }

    string cursorLeft(int k) {
        for (; k && cur != l.begin(); --k)
            cur = prev(cur);
        return text();
    }

    string cursorRight(int k) {
        for (; k && cur != l.end(); --k)
            cur = next(cur);
        return text();
    }
};
```

```go [sol1-Go]
type TextEditor struct {
	*list.List
	cur *list.Element
}

func Constructor() TextEditor {
	l := list.New()
	return TextEditor{l, l.PushBack(nil)} // 哨兵节点
}

func (l *TextEditor) AddText(text string) {
	for _, ch := range text {
		l.cur = l.InsertAfter(byte(ch), l.cur)
	}
}

func (l *TextEditor) DeleteText(k int) int {
	k0 := k
	for ; k > 0 && l.cur.Value != nil; k-- {
		pre := l.cur.Prev()
		l.Remove(l.cur)
		l.cur = pre
	}
	return k0 - k
}

func (l *TextEditor) text() string {
	s := []byte{}
	for k, cur := 10, l.cur; k > 0 && cur.Value != nil; k-- {
		s = append(s, cur.Value.(byte))
		cur = cur.Prev()
	}
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i] // reverse s
	}
	return string(s)
}

func (l *TextEditor) CursorLeft(k int) string {
	for ; k > 0 && l.cur.Value != nil; k-- {
		l.cur = l.cur.Prev()
	}
	return l.text()
}

func (l *TextEditor) CursorRight(k int) string {
	for ; k > 0 && l.cur.Next() != nil; k-- {
		l.cur = l.cur.Next()
	}
	return l.text()
}
```


## 方法二：对顶栈

用两个栈头对头，光标的左右移动就相当于两个栈来回倒；对于插入和删除操作，就相当于在左边那个栈的末尾入栈出栈。

#### 复杂度分析

时空复杂度均与输入量成正比（线性）。

#### 相似题目

- [HDU4699 Editor](http://acm.hdu.edu.cn/showproblem.php?pid=4699)

```Python [sol1-Python3]
class TextEditor:
    def __init__(self):
        self.left, self.right = [], []

    def addText(self, text: str) -> None:
        self.left.extend(list(text))

    def deleteText(self, k: int) -> int:
        k0 = k
        while k and self.left:
            self.left.pop()
            k -= 1
        return k0 - k

    def text(self) -> str:
        return ''.join(self.left[-10:])

    def cursorLeft(self, k: int) -> str:
        while k and self.left:
            self.right.append(self.left.pop())
            k -= 1
        return self.text()

    def cursorRight(self, k: int) -> str:
        while k and self.right:
            self.left.append(self.right.pop())
            k -= 1
        return self.text()
```

```java [sol1-Java]
class TextEditor {
    ArrayList<Character> left = new ArrayList<>(), right = new ArrayList<>();

    public TextEditor() {}

    public void addText(String text) {
        for (var i = 0; i < text.length(); i++)
            left.add(text.charAt(i));
    }

    public int deleteText(int k) {
        var k0 = k;
        for (; k > 0 && !left.isEmpty(); --k)
            left.remove(left.size() - 1);
        return k0 - k;
    }

    String text() {
        var s = new StringBuilder();
        for (var i = Math.max(left.size() - 10, 0); i < left.size(); ++i)
            s.append(left.get(i));
        return s.toString();
    }

    public String cursorLeft(int k) {
        for (; k > 0 && !left.isEmpty(); --k)
            right.add(left.remove(left.size() - 1));
        return text();
    }

    public String cursorRight(int k) {
        for (; k > 0 && !right.isEmpty(); --k)
            left.add(right.remove(right.size() - 1));
        return text();
    }
}
```

```C++ [sol1-C++]
class TextEditor {
    vector<char> left, right;

public:
    TextEditor() {}

    void addText(string text) {
        left.insert(left.end(), text.begin(), text.end());
    }

    int deleteText(int k) {
        int k0 = k;
        for (; k && !left.empty(); --k)
            left.pop_back();
        return k0 - k;
    }

    string text() {
        return string(next(left.begin(), max((int) left.size() - 10, 0)), left.end());
    }

    string cursorLeft(int k) {
        for (; k && !left.empty(); --k) {
            right.emplace_back(left.back());
            left.pop_back();
        }
        return text();
    }

    string cursorRight(int k) {
        for (; k && !right.empty(); --k) {
            left.emplace_back(right.back());
            right.pop_back();
        }
        return text();
    }
};
```

```go [sol1-Go]
type TextEditor struct{ l, r []byte }

func Constructor() TextEditor { return TextEditor{} }

func (t *TextEditor) AddText(text string) {
	t.l = append(t.l, text...)
}

func (t *TextEditor) DeleteText(k int) int {
	k0 := k
	for ; k > 0 && len(t.l) > 0; k-- {
		t.l = t.l[:len(t.l)-1]
	}
	return k0 - k
}

func (t *TextEditor) text() string {
	return string(t.l[max(len(t.l)-10, 0):])
}

func (t *TextEditor) CursorLeft(k int) string {
	for ; k > 0 && len(t.l) > 0; k-- {
		t.r = append(t.r, t.l[len(t.l)-1])
		t.l = t.l[:len(t.l)-1]
	}
	return t.text()
}

func (t *TextEditor) CursorRight(k int) string {
	for ; k > 0 && len(t.r) > 0; k-- {
		t.l = append(t.l, t.r[len(t.r)-1])
		t.r = t.r[:len(t.r)-1]
	}
	return t.text()
}

func max(a, b int) int { if b > a { return b }; return a }
```

## 方法三：Splay（超纲）

本题还可以用 [Splay](https://oi-wiki.org/ds/splay/) 来模拟文本的添加和删除，由于该算法超纲，感兴趣的同学可以查阅相关资料。具体做法在本题的 [视频讲解](https://www.bilibili.com/video/BV1w34y1L7yu/) 中有说明。完整的 Splay 模板见我的 [算法竞赛模板库](https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/splay.go)。

```go [sol3-Go]
type node struct {
	ch  [2]*node
	sz  int
	key byte
}

// 设置如下返回值是为了方便使用 node 中的 ch 数组
func (o *node) cmpKth(k int) int {
	d := k - o.ch[0].size() - 1
	switch {
	case d < 0:
		return 0 // 左儿子
	case d > 0:
		return 1 // 右儿子
	default:
		return -1
	}
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) maintain() {
	o.sz = 1 + o.ch[0].size() + o.ch[1].size()
}

// 构建一颗中序遍历为 [l,r] 的 splay 树
// 比如，给你一个序列和一些修改操作，每次取出一段子区间，cut 掉然后 append 到末尾，输出完成所有操作后的最终序列：
//     我们可以 buildSplay(1,n)，每次操作调用两次 split 来 cut 区间，得到三颗子树 a b c
//     append 之后应该是 a c b，那么我们可以 a.merge(c.merge(b)) 来完成这一操作
//     注意 merge 后可能就不满足搜索树的性质了，但是没有关系，中序遍历的结果仍然是正确的，我们只要保证这一点成立，就能正确得到完成所有操作后的最终序列
func buildSplay(s string) *node {
	if s == "" {
		return nil
	}
	m := len(s) / 2
	o := &node{key: s[m]}
	o.ch[0] = buildSplay(s[:m])
	o.ch[1] = buildSplay(s[m+1:])
	o.maintain()
	return o
}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	o.maintain()
	x.maintain()
	return x
}

// 将子树 o（中序遍历）的第 k 个节点伸展到 o，并返回该节点
// 1 <= k <= o.size()
func (o *node) splay(k int) (kth *node) {
	d := o.cmpKth(k)
	if d < 0 {
		return o
	}
	k -= d * (o.ch[0].size() + 1)
	c := o.ch[d]
	if d2 := c.cmpKth(k); d2 >= 0 {
		c.ch[d2] = c.ch[d2].splay(k - d2*(c.ch[0].size()+1))
		if d2 == d {
			o = o.rotate(d ^ 1)
		} else {
			o.ch[d] = c.rotate(d)
		}
	}
	return o.rotate(d ^ 1)
}

// 分裂子树 o，把 o（中序遍历）的前 k 个节点放在 lo 子树，其余放在 ro 子树
// 返回的 lo 节点为 o（中序遍历）的第 k 个节点
// 1 <= k <= o.size()
// 特别地，k = o.size() 时 ro 为 nil
func (o *node) split(k int) (lo, ro *node) {
	lo = o.splay(k)
	ro = lo.ch[1]
	lo.ch[1] = nil
	lo.maintain()
	return
}

// 把子树 ro 合并进子树 o，返回合并前 o（中序遍历）的最后一个节点
// 相当于把 ro 的中序遍历 append 到 o 的中序遍历之后
// ro 可以为 nil，但 o 不能为 nil
func (o *node) merge(ro *node) *node {
	// 把最大节点伸展上来，这样会空出一个右儿子用来合并 ro
	o = o.splay(o.size())
	o.ch[1] = ro
	o.maintain()
	return o
}

type TextEditor struct {
	root *node
	cur  int
}

func Constructor() TextEditor { return TextEditor{} }

func (t *TextEditor) AddText(text string) {
	if t.cur == 0 {
		t.root = buildSplay(text).merge(t.root)
	} else {
		lo, ro := t.root.split(t.cur)
		t.root = lo.merge(buildSplay(text)).merge(ro)
	}
	t.cur += len(text)
}

func (t *TextEditor) DeleteText(k int) int {
	if t.cur == 0 {
		return 0
	}
	if t.cur <= k { // 左边全部删除
		_, t.root = t.root.split(t.cur)
		ans := t.cur
		t.cur = 0
		return ans
	} else {
		lo, ro := t.root.split(t.cur)
		t.cur -= k
		lo, _ = lo.split(t.cur) // 删除中间 k 个
		t.root = lo.merge(ro)
		return k
	}
}

func (t *TextEditor) text() string {
	if t.cur == 0 {
		return ""
	}
	k := max(t.cur-10, 0)
	t.root = t.root.splay(k + 1)
	ans := make([]byte, 1, t.cur-k)
	ans[0] = t.root.key
	var inorder func(*node) bool
	inorder = func(o *node) bool {
		if o == nil {
			return false
		}
		if inorder(o.ch[0]) || len(ans) == cap(ans) {
			return true
		}
		ans = append(ans, o.key)
		return inorder(o.ch[1])
	}
	inorder(t.root.ch[1])
	return string(ans)
}

func (t *TextEditor) CursorLeft(k int) string {
	t.cur = max(t.cur-k, 0)
	return t.text()
}

func (t *TextEditor) CursorRight(k int) string {
	t.cur = min(t.cur+k, t.root.size())
	return t.text()
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
```
