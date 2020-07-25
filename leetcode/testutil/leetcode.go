package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
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
		if len(rawData) <= 1 || rawData[0] != '"' || rawData[len(rawData)-1] != '"' {
			return reflect.Value{}, invalidErr
		}
		// remove " at leftmost and rightmost
		v = reflect.ValueOf(rawData[1 : len(rawData)-1])
	case reflect.Uint8: // byte
		// rawData like "a" or 'a'
		if len(rawData) != 3 || rawData[0] != '"' && rawData[0] != '\'' || rawData[2] != '"' && rawData[2] != '\'' {
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
		i, er := strconv.Atoi(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(int64(i))
	case reflect.Float64:
		f, er := strconv.ParseFloat(rawData, 64)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(f)
	case reflect.Bool:
		if rawData != "true" && rawData != "false" {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(rawData == "true")
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
	case reflect.Ptr: // *TreeNode, *ListNode, *Point
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
		s = "["
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				s += ","
			}
			_s, er := toRawString(v.Index(i))
			if er != nil {
				return "", er
			}
			s += _s
		}
		s += "]"
	case reflect.Ptr: // *TreeNode, *ListNode, *Point
		switch tpName := v.Type().Elem().Name(); tpName {
		case "TreeNode":
			s = v.Interface().(*TreeNode).toRawString()
		case "ListNode":
			s = v.Interface().(*ListNode).toRawString()
		case "Point":
			s = v.Interface().(*Point).toRawString()
		default:
			return "", fmt.Errorf("unknown type %s", tpName)
		}
	case reflect.String:
		s = fmt.Sprintf(`"%s"`, v.Interface())
	case reflect.Uint8: // byte
		s = fmt.Sprintf(`"%c"`, v.Interface())
	case reflect.Float64:
		s = fmt.Sprintf(`%.5f`, v.Interface())
	default: // int uint int64 bool
		s = fmt.Sprintf(`%v`, v.Interface())
	}
	return
}

// rawExamples[i] = 输入+输出
func RunLeetCodeFuncWithExamples(t *testing.T, f interface{}, rawExamples [][]string, targetCaseNum int) (err error) {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return fmt.Errorf("f must be a function")
	}

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

		if len(example) != fType.NumIn()+fType.NumOut() {
			return fmt.Errorf("len(example) is not %d+%d", fType.NumIn(), fType.NumOut())
		}

		rawIn := example[:fType.NumIn()]
		ins := make([]reflect.Value, len(rawIn))
		for i, rawArg := range rawIn {
			rawArg = trimSpaceAndNewLine(rawArg)
			ins[i], err = parseRawArg(fType.In(i), rawArg)
			if err != nil {
				return
			}
		}
		// just check rawExpectedOuts is valid or not
		rawExpectedOuts := example[fType.NumIn():]
		for i := range rawExpectedOuts {
			rawExpectedOuts[i] = trimSpaceAndNewLine(rawExpectedOuts[i])
			if _, err = parseRawArg(fType.Out(i), rawExpectedOuts[i]); err != nil {
				return
			}
		}

		const maxInputSize = 150
		inputInfo := strings.Join(rawIn, "\n")
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		outs := fValue.Call(ins)
		for i, out := range outs {
			rawActualOut, er := toRawString(out)
			if er != nil {
				return er
			}
			if !assert.Equal(t, rawExpectedOuts[i], rawActualOut, "WA %d\nInput:\n%s", curCaseNum+1, inputInfo) {
				allCasesOk = false
			}
		}
	}

	if targetCaseNum > 0 && allCasesOk {
		t.Logf("case %d is ok", targetCaseNum)
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

func RunLeetCodeClassWithExamples(t *testing.T, constructor interface{}, rawExamples [][3]string, targetCaseNum int) (err error) {
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

	for curCase, example := range rawExamples {
		if targetCaseNum > 0 && curCase+1 != targetCaseNum {
			continue
		}

		names := strings.TrimSpace(example[0])
		inputArgs := strings.TrimSpace(example[1])
		rawExpectedOut := strings.TrimSpace(example[2])

		// parse inputs
		methodNames := []string{}
		for _, name := range strings.Split(names[1:len(names)-1], ",") {
			methodNames = append(methodNames, strings.Title(name[1:len(name)-1]))
		}
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

		// call constructor
		obj := cFunc.Call(constructorIns)[0]

		// we need a pointer to obj cause all methods are declared with a pointer receiver
		pObj := reflect.New(obj.Type())
		pObj.Elem().Set(obj)

		rawActualOut := "[null"
		for callID := 1; callID < len(rawArgsList); callID++ {
			method := pObj.MethodByName(methodNames[callID])
			emptyValue := reflect.Value{}
			if method == emptyValue {
				return fmt.Errorf("invalid test data: %s", methodNames[callID])
			}
			methodType := method.Type()

			// parse method input
			methodArgs, er := parseRawArray(rawArgsList[callID])
			if er != nil {
				return er
			}
			in := make([]reflect.Value, len(methodArgs))
			for i, arg := range methodArgs {
				in[i], err = parseRawArg(methodType.In(i), arg)
				if err != nil {
					return
				}
			}

			// call method
			if actualOuts := method.Call(in); len(actualOuts) > 0 {
				s, er := toRawString(actualOuts[0])
				if er != nil {
					return er
				}
				rawActualOut += "," + s
			} else {
				rawActualOut += ",null"
			}
		}
		rawActualOut += "]"

		if !assert.Equal(t, rawExpectedOut, rawActualOut, "WA %d", curCase+1) {
			allCasesOk = false
		}
	}

	if targetCaseNum > 0 && allCasesOk {
		t.Logf("case %d is ok", targetCaseNum)
		return RunLeetCodeClassWithExamples(t, constructor, rawExamples, 0)
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

func CompareInf(t *testing.T, inputGenerator, runACFunc, runFunc interface{}) {
	ig := reflect.ValueOf(inputGenerator)
	if ig.Kind() != reflect.Func {
		t.Fatal("input generator must be a function")
	}
	runAC := reflect.ValueOf(runACFunc)
	run := reflect.ValueOf(runFunc)
	// just check numbers
	if runAC.Type().NumIn() != run.Type().NumIn() || runAC.Type().NumOut() != run.Type().NumOut() {
		t.Fatal("different input/output")
	}

	for tc := 1; ; tc++ {
		inArgs := ig.Call(nil)
		expectedOut := runAC.Call(inArgs)
		actualOut := run.Call(inArgs)

		//ins := make([]interface{}, len(inArgs))
		//for i, in := range inArgs {
		//	ins[i] = in.Interface()
		//}
		insStr := []byte{}
		for i, arg := range inArgs {
			if i > 0 {
				insStr = append(insStr, '\n')
			}
			s, err := toRawString(arg)
			if err != nil {
				t.Fatal(err)
			}
			insStr = append(insStr, s...)
		}
		for i, eOut := range expectedOut {
			assert.Equal(t, eOut.Interface(), actualOut[i].Interface(), "WA %d\nInput:\n%s", tc, insStr)
		}

		if tc%1e5 == 0 {
			t.Logf("%d cases passed.", tc)
		}
	}
}
