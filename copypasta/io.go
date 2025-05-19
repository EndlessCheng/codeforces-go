package copypasta

import (
	"bufio"
	"bytes"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// 带有 IO 缓冲区的输入输出，适用于绝大多数题目
//
// bufio.NewReader(os.Stdin) 相比 fmt.Scan，每读入 1e5 个 int 可以加速约 1300ms（Codeforces/AtCoder）
// 对比：（3e4 个 int）
// 623ms https://codeforces.com/problemset/submission/981/124239306
// 233ms https://codeforces.com/problemset/submission/981/124237530
//
// bufio.NewWriter(os.Stdout) 相比 fmt.Println，每减少 1e5 次直接换行输出，可以加速约 200ms（Codeforces/AtCoder）
// 换句话说，对于 <~1e4 的输出量，加不加 buffer 都一样，直接用 fmt.Println 输出即可（大多数 CF 题目都是 T<~2e4 的）
// 对比：（1e5 个 int）
// 405ms https://codeforces.com/contest/1603/submission/135520593
// 187ms https://codeforces.com/contest/1603/submission/134450945
// NOTE: 调用 Fprintln 打印 int(0)   1e6 次的耗时为  77ms https://codeforces.com/contest/1603/submission/169796327
// NOTE: 调用 Fprintln 打印 int(1e9) 1e6 次的耗时为 155ms https://codeforces.com/contest/1603/submission/169796385
func bufferIO() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n) // 如果行数未知，可以根据 Fscan 的第一个返回值是否为正来决定

	Fprintln(out, n)
}

// 快读，适用于输入量巨大的题目
// 相比上面的 bufferIO，每读入 1e6 个 int 可以加速约 400~450ms（Codeforces/AtCoder）   从字符串的角度来说是 1e7
// 更快的写法见更下面的 fasterIO
func fastIO() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// 读一个整数
	r := func() int {
		in.Scan()
		x, _ := strconv.Atoi(string(in.Bytes()))
		return x
	}
	// 更快的写法（非负数）
	r = func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b&15)
		}
		return
	}
	// 支持负数的写法
	r = func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data[1:] {
				x = x*10 + int(b&15)
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b&15)
		}
		return
	}

	// 读一个浮点数
	rf := func() float64 {
		in.Scan()
		f, _ := strconv.ParseFloat(string(in.Bytes()), 64)
		return f
	}
	// 更快的写法
	rf = func() float64 {
		in.Scan()
		s := in.Bytes()
		neg := false
		if s[0] == '-' {
			neg = true
			s = s[1:]
		}
		dotPos := len(s) - 1
		f := 0
		for i, b := range s {
			if b == '.' {
				dotPos = i
			} else {
				f = f*10 + int(b&15)
			}
		}
		if neg {
			f = -f
		}
		return float64(f) / math.Pow10(len(s)-1-dotPos) // 放心，math.Pow10 会直接查表，非常快
	}

	// NOTE: bufio.Scanner 在读长字符串的情况下可能会有奇怪的 bug，所以还是用下面的 fasterIO 吧！（CF827A WA5）
	in.Buffer(nil, 1e9)
	rs := func() []byte { in.Scan(); return in.Bytes() }

	_ = []interface{}{r, rf, rs}
}

// 超快读
// 以 CF1276B（1e6 数据量）为例，测试结果如下：
// bufferIO  670 ms
// fastIO    296 ms
// fasterIO  202 ms
// fasterIO  202 ms (use syscall.Read(syscall.Stdin, buf))
// NOTE: fasterIO 下的纯读入耗时为 61ms https://codeforces.com/contest/1276/submission/142793894
// 选择 4KB 作为缓存块大小的原因 https://stackoverflow.com/questions/6578394/whats-so-special-about-4kb-for-a-buffer-length
// NOTE: 如果只有数字的话，只需要判断字符与 '0' 的关系就行了；有小写字母的话，与 'z' 的大小判断可以省去（对运行耗时无影响）
// NOTE: 额外的好处是，这种避开 Fscan 的写法可以节省一部分内存（1e6 下有 10M 左右）
// C++ 选手可以参考 https://codeforces.com/contest/1826/submission/204581714
func fasterIO() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const eof = 0
	buf := make([]byte, 4096) // 4KB
	_i, _n := 0, 0

	// 读一个字符
	rc := func() byte {
		if _i == _n {
			_n, _ = os.Stdin.Read(buf)
			// EOF 一定要判断！不判断会 RE：https://codeforces.com/problemset/submission/323/250522741
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}

	// 读一个非负整数
	rd := func() (x int) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	// 读一个整数，支持负数
	rd = func() (x int) {
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	// 读一个数字或字母
	r1 := func() byte {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		return b
	}

	// 读一个仅包含小写字母的字符串
	rs := func() (s []byte) {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		}
		for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}

	// 读一个长度为 n 的仅包含小写字母的字符串
	rsn := func(n int) []byte {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		}
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}

	// 如果只有/还剩下一个长度未知的字符串（仅包含小写字母）
	readStringUntilEOF := func() (s []byte) {
		// 若之前 Read 过……
		for _i < len(buf) && buf[_i] < 'a' { // 'A'
			_i++
		}
		s = append(s, buf[_i:]...)

		// 核心是这一段
		for {
			n, _ := os.Stdin.Read(buf)
			if n == 0 {
				break
			}
			s = append(s, buf[:n]...)
		}

		// 注意末尾有 \r \n 的情况
		for ; s[len(s)-1] < 'a'; s = s[:len(s)-1] { // 'A'
		}
		return
	}

	// 手写输出，可能会加快几十 ms
	// 使用前 202ms https://codeforces.com/contest/1208/submission/176961129 （新版）https://codeforces.com/contest/1208/submission/269050123 218ms
	// 使用后 139ms https://codeforces.com/contest/1208/submission/176963572 （新版）https://codeforces.com/contest/1208/submission/269052669 171ms
	// 注：也可以全部初始化成空格/换行，这样可以直接倒着写入（需要 OJ 支持特判输出有多个空格/换行），不过实测没啥区别，处理负号还要多写一些逻辑
	// 注：也可以创建一个全局数组 _o，然后 outS := _o[:0]，不过效率几乎一样

	const outputN int = 1e6                       // 输出的 int 个数的最大值
	const intWidth = 20                           // 输出的 int 绝对值的十进制长度的最大值
	outS := make([]byte, 0, outputN*(intWidth+2)) // 如果没有负数，+2 改成 +1
	tmpS := [intWidth]byte{}                      // 临时保存输出的内容（因为遍历数位是从右往左）
	wInt := func(x int) {
		if x == 0 { // 如果保证 x != 0 则去掉
			outS = append(outS, '0')
			return
		}
		if x < 0 { // 如果保证 x >= 0 则去掉
			x = -x
			outS = append(outS, '-')
		}
		p := len(tmpS)
		for ; x > 0; x /= 10 {
			p--
			tmpS[p] = '0' | byte(x%10)
		}
		outS = append(outS, tmpS[p:]...)
		//outS = append(outS, '\n') // 空格/换行需要手动添加
	}

	// 最后，直接用 os.Stdout 输出（最上面的 out 是不需要创建的）
	os.Stdout.Write(outS)

	_ = []interface{}{rd, r1, rs, rsn, readStringUntilEOF, wInt}
}

// 如果输入按照行来读入更方便的话……
// 数据个数未知 https://www.luogu.com.cn/problem/P2762
// 仅加速用 https://codeforces.com/problemset/problem/375/B
// 注意由于 buffer 的缘故，bufio.Scanner 不要和 bufio.Reader 混用
// 如果每行只有几个数，可以用 fmt.Fscanln 读入
func lineIO() {
	in := bufio.NewScanner(os.Stdin) // 默认 4KB 初始 buffer
	in.Buffer(nil, math.MaxInt)      // 若单个 token 大小超过 65536 则加上这行，否则会报错
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for in.Scan() {
		line := in.Bytes()
		sp := bytes.Split(line, []byte{' '})
		// ...

		_ = sp
	}
}
