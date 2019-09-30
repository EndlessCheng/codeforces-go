package copypasta

import (
	"bufio"
	. "fmt"
	"os"
)

func simpleIO() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	Fprintln(out, n)
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
