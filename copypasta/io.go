package copypasta

import (
	"bufio"
	. "fmt"
	"os"
)

func simpleIO() {
	// NOTE: just a bufio.NewReader is enough, there is no difference between this and ioutil.ReadAll
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	Fprintln(out, n)
	// NOTE: to print int as char, use Fprintf(out, "%c", 'a'+1)
}

func lineIO() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(nil, 1e6+5) // default maxTokenSize is 65536
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for in.Scan() {
		line := in.Text()

		Fprintln(out, line)
	}
}
