package testutil

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func parseRawArray(rawArray string) (splits []string, err error) {
	invalidErr := fmt.Errorf("invalid test data: %s", rawArray)

	// check [] at leftmost and rightmost
	if len(rawArray) <= 1 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, invalidErr
	}

	// ignore [] at leftmost and rightmost
	rawArray = rawArray[1 : len(rawArray)-1]
	if rawArray == "" {
		return
	}

	isPoint := rawArray[0] == '('

	const sep = ','
	var depth, quotCnt, bracketCnt int
	for start := 0; start < len(rawArray); {
		end := start
	outer:
		for ; end < len(rawArray); end++ {
			switch rawArray[end] {
			case '[':
				depth++
			case ']':
				depth--
			case '"':
				quotCnt++
			case '(':
				bracketCnt++
			case ')':
				bracketCnt--
			case sep:
				if depth == 0 {
					if !isPoint {
						if quotCnt%2 == 0 {
							break outer
						}
					} else {
						if bracketCnt%2 == 0 {
							break outer
						}
					}
				}
			}
		}
		splits = append(splits, strings.TrimSpace(rawArray[start:end]))
		start = end + 1 // skip sep
	}
	if depth != 0 || quotCnt%2 != 0 {
		return nil, invalidErr
	}
	return
}

func parseRawArg(tp reflect.Type, rawData string) (v reflect.Value, err error) {
	rawData = strings.TrimSpace(rawData)
	invalidErr := fmt.Errorf("invalid test data: %s", rawData)
	switch tp.Kind() {
	case reflect.String:
		if len(rawData) <= 1 || rawData[0] != '"' && rawData[0] != '\'' || rawData[len(rawData)-1] != rawData[0] {
			return reflect.Value{}, invalidErr
		}
		// remove " (or ') at leftmost and rightmost
		v = reflect.ValueOf(rawData[1 : len(rawData)-1])
	case reflect.Uint8: // byte
		// rawData like "a" or 'a'
		if len(rawData) != 3 || rawData[0] != '"' && rawData[0] != '\'' || rawData[2] != rawData[0] {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(rawData[1])
	case reflect.Int:
		i, er := strconv.Atoi(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(i)
	case reflect.Uint:
		i, er := strconv.Atoi(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(uint(i))
	case reflect.Int64:
		i, er := strconv.ParseInt(rawData, 10, 64)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(i)
	case reflect.Uint64:
		i, er := strconv.ParseUint(rawData, 10, 64)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(i)
	case reflect.Float64:
		f, er := strconv.ParseFloat(rawData, 64)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(f)
	case reflect.Bool:
		b, er := strconv.ParseBool(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(b)
	case reflect.Slice:
		splits, er := parseRawArray(rawData)
		if er != nil {
			return reflect.Value{}, er
		}
		v = reflect.New(tp).Elem()
		for _, s := range splits {
			_v, er := parseRawArg(tp.Elem(), s)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.Append(v, _v)
		}
	case reflect.Ptr: // *TreeNode, *ListNode, *Point, *Interval
		switch tpName := tp.Elem().Name(); tpName {
		case "TreeNode":
			root, er := buildTreeNode(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(root)
		case "ListNode":
			head, er := buildListNode(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(head)
		case "Point": // nowcoder
			p, er := buildPoint(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(p)
		case "Interval": // nowcoder
			p, er := buildInterval(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(p)
		default:
			return reflect.Value{}, fmt.Errorf("unknown type %s", tpName)
		}
	default:
		return reflect.Value{}, fmt.Errorf("unknown type %s", tp.Name())
	}
	return
}

func toRawString(v reflect.Value) (s string, err error) {
	switch v.Kind() {
	case reflect.Slice:
		sb := &strings.Builder{}
		sb.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				sb.WriteByte(',')
			}
			_s, er := toRawString(v.Index(i))
			if er != nil {
				return "", er
			}
			sb.WriteString(_s)
		}
		sb.WriteByte(']')
		s = sb.String()
	case reflect.Ptr: // *TreeNode, *ListNode, *Point, *Interval
		switch tpName := v.Type().Elem().Name(); tpName {
		case "TreeNode":
			s = v.Interface().(*TreeNode).toRawString()
		case "ListNode":
			s = v.Interface().(*ListNode).toRawString()
		case "Point":
			s = v.Interface().(*Point).toRawString()
		case "Interval":
			s = v.Interface().(*Interval).toRawString()
		default:
			return "", fmt.Errorf("unknown type %s", tpName)
		}
	case reflect.String:
		s = fmt.Sprintf(`"%s"`, v.Interface())
	case reflect.Uint8: // byte
		s = fmt.Sprintf(`"%c"`, v.Interface())
	case reflect.Float64:
		s = fmt.Sprintf(`%.5f`, v.Interface())
	default: // int uint int64 uint64 bool
		s = fmt.Sprintf(`%v`, v.Interface())
	}
	return
}

func isTLE(f func()) bool {
	if DebugTLE == 0 || IsDebugging() {
		f()
		return false
	}

	done := make(chan struct{})
	timer := time.NewTimer(DebugTLE)
	defer timer.Stop()
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}

// rawExamples[i] = 输入+输出
// 若反射出来的函数或 rawExamples 数据不合法，则会返回一个非空的 error，否则返回 nil
func RunLeetCodeFuncWithExamples(t *testing.T, f interface{}, rawExamples [][]string, targetCaseNum int) (err error) {
	if len(rawExamples) == 0 {
		return fmt.Errorf("test cases is empty")
	}

	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return fmt.Errorf("f must be a function")
	}

	fNumIn := fType.NumIn()
	fNumOut := fType.NumOut()

	// 例如，-1 表示最后一个测试用例
	if targetCaseNum < 0 {
		targetCaseNum += len(rawExamples) + 1
	}

	allCasesOk := true
	fValue := reflect.ValueOf(f)
	for curCaseNum, example := range rawExamples {
		if targetCaseNum > 0 && curCaseNum+1 != targetCaseNum {
			continue
		}

		if len(example) != fNumIn+fNumOut {
			return fmt.Errorf("len(example) = %d, but we need %d+%d", len(example), fNumIn, fNumOut)
		}

		rawIn := example[:fNumIn]
		ins := make([]reflect.Value, len(rawIn))
		for i, rawArg := range rawIn {
			rawArg = trimSpace(rawArg)
			ins[i], err = parseRawArg(fType.In(i), rawArg)
			if err != nil {
				return
			}
		}
		// just check rawExpectedOuts is valid or not
		rawExpectedOuts := example[fNumIn:]
		for i := range rawExpectedOuts {
			rawExpectedOuts[i] = trimSpace(rawExpectedOuts[i])
			if _, err = parseRawArg(fType.Out(i), rawExpectedOuts[i]); err != nil {
				return
			}
		}

		t.Run(fmt.Sprintf("Case %d", curCaseNum+1), func(t *testing.T) {
			const maxInputSize = 150
			inputInfo := strings.Join(rawIn, "\n")
			if len(inputInfo) > maxInputSize { // 截断过长的输入
				inputInfo = inputInfo[:maxInputSize] + "..."
			}

			var outs []reflect.Value
			_f := func() { outs = fValue.Call(ins) }
			if targetCaseNum == 0 && isTLE(_f) {
				allCasesOk = false
				t.Errorf("【超时 %d】\nInput:\n%s", curCaseNum+1, inputInfo)
				return
			}
			if targetCaseNum != 0 {
				_f()
			}

			for i, out := range outs {
				rawActualOut, er := toRawString(out)
				if er != nil {
					t.Fatal(er)
				}
				if AssertOutput && !assert.Equal(t, rawExpectedOuts[i], rawActualOut, "【答案错误 %d】\nInput:\n%s", curCaseNum+1, inputInfo) {
					allCasesOk = false
				}
			}
		})
	}

	// 若有测试用例未通过，则前面必然会打印一些信息，这里直接返回
	if !allCasesOk {
		return nil
	}

	// 若测试的是单个用例，则接着测试所有用例
	if targetCaseNum > 0 {
		t.Logf("case %d is passed", targetCaseNum)
		return RunLeetCodeFuncWithExamples(t, f, rawExamples, 0)
	}

	return nil
}

func RunLeetCodeFuncWithCase(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string, targetCaseNum int) (err error) {
	examples := [][]string{}
	for i, input := range rawInputs {
		examples = append(examples, append(append([]string{}, input...), rawOutputs[i]...))
	}
	return RunLeetCodeFuncWithExamples(t, f, examples, targetCaseNum)
}

func RunLeetCodeFunc(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string) error {
	return RunLeetCodeFuncWithCase(t, f, rawInputs, rawOutputs, 0)
}

func RunLeetCodeFuncWithFile(t *testing.T, f interface{}, filePath string, targetCaseNum int) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := trimSpaceAndEmptyLine(string(data))
	n := len(lines)
	if n == 0 {
		return fmt.Errorf("输入数据为空，请检查文件路径和文件内容是否正确")
	}

	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return fmt.Errorf("f 必须是函数")
	}

	// 每 fNumIn+fNumOut 行一组数据
	fNumIn := fType.NumIn()
	fNumOut := fType.NumOut()
	tcSize := fNumIn + fNumOut
	if n%tcSize != 0 {
		return fmt.Errorf("有效行数 %d，应该是 %d 的倍数", n, tcSize)
	}

	examples := make([][]string, 0, n/tcSize)
	for i := 0; i < n; i += tcSize {
		examples = append(examples, lines[i:i+tcSize])
	}
	return RunLeetCodeFuncWithExamples(t, f, examples, targetCaseNum)
}

//

// 若反射出来的函数或 rawExamples 数据不合法，则会返回一个非空的 error，否则返回 nil
func RunLeetCodeClassWithExamples(t *testing.T, constructor interface{}, rawExamples [][3]string, targetCaseNum int) (err error) {
	if len(rawExamples) == 0 {
		return fmt.Errorf("test cases is empty")
	}

	cType := reflect.TypeOf(constructor)
	if cType.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a function")
	}
	if cType.NumOut() != 1 {
		return fmt.Errorf("constructor must have one and only one return value")
	}
	allCasesOk := true
	cFunc := reflect.ValueOf(constructor)

	// 例如，-1 表示最后一个测试用例
	if targetCaseNum < 0 {
		targetCaseNum += len(rawExamples) + 1
	}

	for curCaseNum, example := range rawExamples {
		if targetCaseNum > 0 && curCaseNum+1 != targetCaseNum {
			continue
		}

		names := strings.TrimSpace(example[0])
		inputArgs := strings.TrimSpace(example[1])
		rawExpectedOut := strings.TrimSpace(example[2])

		// parse called names
		// 调用 parseRawArray 确保数据是合法的
		methodNames, er := parseRawArray(names)
		if er != nil {
			return er
		}
		for i, name := range methodNames {
			name = name[1 : len(name)-1] // 移除引号
			name = strings.Title(name)   // 首字母大写以匹配模板方法名称
			methodNames[i] = name
		}

		// parse inputs
		rawArgsList, er := parseRawArray(inputArgs)
		if er != nil {
			return er
		}
		if len(rawArgsList) != len(methodNames) {
			return fmt.Errorf("invalid test data: mismatch names and input args (%d != %d)", len(methodNames), len(rawArgsList))
		}

		// parse constructor input
		constructorArgs, er := parseRawArray(rawArgsList[0])
		if er != nil {
			return er
		}
		constructorIns := make([]reflect.Value, len(constructorArgs))
		for i, arg := range constructorArgs {
			constructorIns[i], err = parseRawArg(cType.In(i), arg)
			if err != nil {
				return
			}
		}

		t.Run(fmt.Sprintf("Case %d", curCaseNum+1), func(t *testing.T) {
			// call constructor, get struct instance
			obj := cFunc.Call(constructorIns)[0]

			// use a pointer to call methods
			pObj := reflect.New(obj.Type())
			pObj.Elem().Set(obj)

			if DebugCallIndex < 0 {
				DebugCallIndex += len(rawArgsList)
			}
			rawActualOut := &strings.Builder{}
			rawActualOut.WriteString("[null")
			for callIndex := 1; callIndex < len(rawArgsList); callIndex++ {
				name := methodNames[callIndex]
				method := pObj.MethodByName(name)
				emptyValue := reflect.Value{}
				if method == emptyValue {
					t.Fatalf("invalid test data: %s", methodNames[callIndex])
				}
				methodType := method.Type()

				// parse method input
				methodArgs, er := parseRawArray(rawArgsList[callIndex])
				if er != nil {
					t.Fatal(er)
				}
				in := make([]reflect.Value, methodType.NumIn()) // 注意：若入参为空，methodArgs 可能是 [] 也可能是 [null]
				for i := range in {
					in[i], er = parseRawArg(methodType.In(i), methodArgs[i])
					if er != nil {
						t.Fatal(er)
					}
				}

				if callIndex == DebugCallIndex {
					print() // 在这里打断点
				}

				// call method
				var actualOuts []reflect.Value
				_f := func() { actualOuts = method.Call(in) }
				if targetCaseNum == 0 && isTLE(_f) {
					allCasesOk = false
					t.Errorf("【【超时 %d】】\nCall Index %d", curCaseNum+1, callIndex)
					return // 直接跑下一个测试用例
				} 
				if targetCaseNum != 0 {
					_f()
				}

				if len(actualOuts) > 0 {
					s, er := toRawString(actualOuts[0])
					if er != nil {
						t.Fatal(er)
					}
					rawActualOut.WriteByte(',')
					rawActualOut.WriteString(s)
				} else {
					rawActualOut.WriteString(",null")
				}
			}
			rawActualOut.WriteByte(']')

			// 比较前，去除 rawExpectedOut 中逗号后的空格
			rawExpectedOut = strings.ReplaceAll(rawExpectedOut, ", ", ",")

			// todo: 提示错在哪个 callIndex 上
			if AssertOutput && !assert.Equal(t, rawExpectedOut, rawActualOut.String(), "【答案错误 %d】", curCaseNum+1) {
				allCasesOk = false
			}
		})
	}

	if targetCaseNum > 0 && allCasesOk {
		t.Logf("case %d is ok", targetCaseNum)
		return RunLeetCodeClassWithExamples(t, constructor, rawExamples, 0)
	}

	if allCasesOk {
		t.Log("OK")
	}

	return nil
}

func RunLeetCodeClassWithCase(t *testing.T, constructor interface{}, rawInputs, rawOutputs []string, targetCaseNum int) (err error) {
	examples := [][3]string{}
	for i, input := range rawInputs {
		input = strings.TrimSpace(input)
		lines := strings.Split(input, "\n")
		examples = append(examples, [3]string{lines[0], lines[1], rawOutputs[i]})
	}
	return RunLeetCodeClassWithExamples(t, constructor, examples, targetCaseNum)
}

func RunLeetCodeClass(t *testing.T, constructor interface{}, rawInputs, rawOutputs []string) error {
	return RunLeetCodeClassWithCase(t, constructor, rawInputs, rawOutputs, 0)
}

func RunLeetCodeClassWithFile(t *testing.T, constructor interface{}, filePath string, targetCaseNum int) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := trimSpaceAndEmptyLine(string(data))
	n := len(lines)
	if n == 0 {
		return fmt.Errorf("输入数据为空，请检查文件路径和文件内容是否正确")
	}

	// 每三行一组数据
	if n%3 != 0 {
		return fmt.Errorf("有效行数 %d，应该是 3 的倍数", n)
	}

	examples := make([][3]string, 0, n/3)
	for i := 0; i < n; i += 3 {
		examples = append(examples, [3]string{lines[i], lines[i+1], lines[i+2]})
	}
	return RunLeetCodeClassWithExamples(t, constructor, examples, targetCaseNum)
}

// 无尽对拍模式
func CompareInf(t *testing.T, inputGenerator, runACFunc, runFunc interface{}) {
	ig := reflect.ValueOf(inputGenerator)
	if ig.Kind() != reflect.Func {
		t.Fatal("input generator must be a function")
	}
	runAC := reflect.ValueOf(runACFunc)
	run := reflect.ValueOf(runFunc)
	// just check numbers
	if !assert.Equal(t, run.Type().NumIn(), runAC.Type().NumIn()) ||
		!assert.Equal(t, run.Type().NumOut(), runAC.Type().NumOut()) {
		t.Fatal("different input/output")
	}

	for tc := 1; ; tc++ {
		inArgs := ig.Call(nil)

		// 先生成字符串，以免 inArgs 被修改
		inputInfo := []byte{}
		for i, arg := range inArgs {
			if i > 0 {
				inputInfo = append(inputInfo, '\n')
			}
			s, err := toRawString(arg)
			if err != nil {
				t.Fatal(err)
			}
			inputInfo = append(inputInfo, s...)
		}

		// todo deep copy slice
		expectedOut := runAC.Call(inArgs)
		var actualOut []reflect.Value
		if isTLE(func() { actualOut = run.Call(inArgs) }) {
			t.Errorf("【超时 %d】\nInput:\n%s", tc, inputInfo)
			continue
		}

		for i, eOut := range expectedOut {
			assert.Equal(t, eOut.Interface(), actualOut[i].Interface(), "【答案错误 %d】\nInput:\n%s", tc, inputInfo)
		}

		// 每到 2 的幂次就打印检测了多少个测试数据
		if tc&(tc-1) == 0 {
			t.Logf("%d cases checked.", tc)
		}

		if Once {
			break
		}
	}
}
