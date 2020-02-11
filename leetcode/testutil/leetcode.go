package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func parseRawArg(tp reflect.Type, rawArg string) (v reflect.Value, err error) {
	invalidErr := fmt.Errorf("invalid test data: %s", rawArg)
	switch tp.Kind() {
	case reflect.String:
		if len(rawArg) < 2 || rawArg[0] != '"' || rawArg[len(rawArg)-1] != '"' {
			err = invalidErr
			return
		}
		// remove " at leftmost and rightmost
		v = reflect.ValueOf(rawArg[1 : len(rawArg)-1])
	case reflect.Uint8: // byte
		// sth like "a"
		if len(rawArg) != 3 || rawArg[0] != '"' || rawArg[2] != '"' {
			err = invalidErr
			return
		}
		v = reflect.ValueOf(rawArg[1])
	case reflect.Int:
		i, er := strconv.Atoi(rawArg)
		if er != nil {
			err = invalidErr
			return
		}
		v = reflect.ValueOf(i)
	case reflect.Uint:
		i, er := strconv.Atoi(rawArg)
		if er != nil {
			err = invalidErr
			return
		}
		v = reflect.ValueOf(uint(i))
	case reflect.Float64:
		f, er := strconv.ParseFloat(rawArg, 64)
		if er != nil {
			err = invalidErr
			return
		}
		v = reflect.ValueOf(f)
	case reflect.Bool:
		if rawArg != "true" && rawArg != "false" {
			err = invalidErr
			return
		}
		v = reflect.ValueOf(rawArg == "true")
	case reflect.Slice:
		// check [] at leftmost and rightmost
		if len(rawArg) <= 1 || rawArg[0] != '[' || rawArg[len(rawArg)-1] != ']' {
			err = invalidErr
			return
		}
		// ignore [] at leftmost and rightmost
		rawArg = rawArg[1 : len(rawArg)-1]

		v = reflect.New(tp).Elem()
		isStringSlice := strings.Contains(rawArg, `"`)
		depth := 0
		quotCnt := 0
		for start := 0; start < len(rawArg); {
			end := start
		outer:
			for ; end < len(rawArg); end++ {
				switch rawArg[end] {
				case '[':
					depth++
				case ']':
					depth--
				case '"':
					quotCnt++
				case ',':
					if depth == 0 {
						if isStringSlice && quotCnt%2 == 1 {
							continue
						}
						break outer
					}
				}
			}
			_v, er := parseRawArg(tp.Elem(), rawArg[start:end])
			if er != nil {
				err = er
				return
			}
			v = reflect.Append(v, _v)
			start = end + 1 // skip ,
		}
		if depth != 0 || quotCnt%2 != 0 {
			err = invalidErr
			return
		}
	default:
		err = fmt.Errorf("unknown type %s", tp.Name())
		return
	}
	return
}

func simpleValueString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Slice:
		res := "["
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				res += ","
			}
			res += simpleValueString(v.Index(i))
		}
		res += "]"
		return res
	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.Interface())
	case reflect.Uint8: // byte
		return fmt.Sprintf(`"%c"`, v.Interface())
	default: // int uint float64 bool
		return fmt.Sprintf(`%v`, v.Interface())
	}
}

func RunLeetCodeFuncWithCase(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string, targetCaseNum int) (err error) {
	tp := reflect.TypeOf(f)
	if tp.Kind() != reflect.Func {
		return fmt.Errorf("f must be a function")
	}

	allOk := true
	vFunc := reflect.ValueOf(f)
	for testCase, rawIn := range rawInputs {
		if targetCaseNum > 0 && testCase+1 != targetCaseNum {
			continue
		}

		if len(rawIn) != tp.NumIn() {
			return fmt.Errorf("len(rawIn) is not %d", tp.NumIn())
		}
		rawOut := rawOutputs[testCase]
		if len(rawOut) != tp.NumOut() {
			return fmt.Errorf("len(rawOut) is not %d", tp.NumOut())
		}

		in := make([]reflect.Value, len(rawIn))
		for i, rawArg := range rawIn {
			rawArg = trimSpaceAndNewLine(rawArg)
			in[i], err = parseRawArg(tp.In(i), rawArg)
			if err != nil {
				// rawArg 不合法
				return
			}
		}
		// 额外检测 rawOutputs 是否合法
		for i, rawArg := range rawOut {
			rawArg = trimSpaceAndNewLine(rawArg)
			if _, err = parseRawArg(tp.Out(i), rawArg); err != nil {
				// rawArg 不合法
				return
			}
		}

		actualOuts := vFunc.Call(in)
		for i, expectedRes := range rawOut {
			if !assert.Equal(t, expectedRes, simpleValueString(actualOuts[i]), "please check case %d", testCase+1) {
				allOk = false
			}
		}
	}

	if targetCaseNum > 0 && allOk {
		t.Logf("case %d is ok", targetCaseNum)
		return RunLeetCodeFuncWithCase(t, f, rawInputs, rawOutputs, 0)
	}

	return nil
}

func RunLeetCodeFunc(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string) error {
	return RunLeetCodeFuncWithCase(t, f, rawInputs, rawOutputs, 0)
}

func splitClassArgs(rawData string) (rawArgs []string, err error) {
	invalidErr := fmt.Errorf("invalid test data: %s", rawData)
	// check [] at leftmost and rightmost
	if len(rawData) <= 1 || rawData[0] != '[' || rawData[len(rawData)-1] != ']' {
		return nil, invalidErr
	}
	// ignore [] at leftmost and rightmost
	rawData = rawData[1 : len(rawData)-1]

	hasStringArg := strings.Contains(rawData, `"`)
	depth := 0
	quotCnt := 0
	for start := 0; start < len(rawData); {
		end := start
	outer:
		for ; end < len(rawData); end++ {
			switch rawData[end] {
			case '[':
				depth++
			case ']':
				depth--
			case '"':
				quotCnt++
			case ',':
				if depth == 0 {
					if hasStringArg && quotCnt%2 == 1 {
						continue
					}
					break outer
				}
			}
		}
		rawArgs = append(rawArgs, rawData[start:end])
		start = end + 1 // skip ,
	}
	if depth != 0 || quotCnt%2 != 0 {
		return nil, invalidErr
	}
	return
}

func RunLeetCodeClassWithCase(t *testing.T, constructor interface{}, rawInputs, rawOutputs []string, targetCaseNum int) (err error) {
	constructorType := reflect.TypeOf(constructor)
	if constructorType.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a function")
	}
	allOk := true
	constructorFunc := reflect.ValueOf(constructor)

	for testCase, rawIn := range rawInputs {
		if targetCaseNum > 0 && testCase+1 != targetCaseNum {
			continue
		}

		invalidErr := fmt.Errorf("invalid test data: %s", rawIn)

		// parse inputs
		splits := strings.Split(strings.TrimSpace(rawIn), "\n")
		if len(splits) != 2 {
			return invalidErr
		}
		methodNames := []string{}
		for _, name := range strings.Split(splits[0][1:len(splits[0])-1], ",") {
			methodNames = append(methodNames, strings.Title(name[1:len(name)-1]))
		}
		rawArgsList, er := splitClassArgs(splits[1])
		if er != nil {
			return er
		}
		if len(rawArgsList) != len(methodNames) {
			return fmt.Errorf("invalid test data: mismatch names and input args (%d != %d)", len(methodNames), len(rawArgsList))
		}

		// parse constructor input
		constructorArgs, er := splitClassArgs(rawArgsList[0])
		if er != nil {
			return er
		}
		constructorIn := make([]reflect.Value, len(constructorArgs))
		for i, arg := range constructorArgs {
			constructorIn[i], err = parseRawArg(constructorType.In(i), arg)
			if err != nil {
				return
			}
		}

		// call constructor
		objs := constructorFunc.Call(constructorIn)

		// we need a obj pointer cause all methods are declared with a pointer receiver
		obj := objs[0]
		pObj := reflect.New(obj.Type())
		pObj.Elem().Set(obj)

		actualOutStr := "[null"
		for callID := 1; callID < len(rawArgsList); callID++ {
			method := pObj.MethodByName(methodNames[callID])
			methodType := method.Type()

			// parse method input
			methodArgs, er := splitClassArgs(rawArgsList[callID])
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
				actualOutStr += "," + simpleValueString(actualOuts[0])
			} else {
				actualOutStr += ",null"
			}
		}
		actualOutStr += "]"

		if !assert.Equal(t, rawOutputs[testCase], actualOutStr, "please check case %d", testCase+1) {
			allOk = false
		}
	}

	if targetCaseNum > 0 && allOk {
		t.Logf("case %d is ok", targetCaseNum)
		return RunLeetCodeClassWithCase(t, constructor, rawInputs, rawOutputs, 0)
	}

	return nil
}

func RunLeetCodeClass(t *testing.T, constructor interface{}, rawInputs, rawOutputs []string) error {
	return RunLeetCodeClassWithCase(t, constructor, rawInputs, rawOutputs, 0)
}
