package copypasta

import (
	"bufio"
	. "fmt"
	"io"
)

func simpleIO(reader io.Reader, writer io.Writer) {
	// NOTE: just a bufio.NewReader is enough, there is no difference between this and ioutil.ReadAll
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	Fprintln(out, n)
	// NOTE: to print int as char, use Fprintf(out, "%c", 'a'+1)
	// NOTE: to print []byte as string, use Fprintf(out, "%s", data)
}

// 数据量在 ~10^6 时使用，能明显加快运行速度！
func fastIO(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	readInt := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}
	readInt = func() (x int) {
		in.Scan()
		data := in.Bytes()
		sig := 1
		if data[0] == '-' {
			sig = -1
			data = data[1:]
		}
		for _, b := range data {
			x = x*10 + int(b-'0')
		}
		return sig * x
	}

	_ = []interface{}{readInt}
}

func lineIO(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Buffer(nil, 1e6+5) // default maxTokenSize is 65536
	out := bufio.NewWriter(writer)
	defer out.Flush()

	for in.Scan() {
		line := in.Text()

		Fprintln(out, line)
	}
}
