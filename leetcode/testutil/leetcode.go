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
	switch tp.Kind() {
	case reflect.String:
		// remove " at leftmost and rightmost
		v = reflect.ValueOf(rawArg[1 : len(rawArg)-1])
	case reflect.Uint8: // byte
		// sth like "a"
		v = reflect.ValueOf(rawArg[1])
	case reflect.Int:
		i, _ := strconv.Atoi(rawArg)
		v = reflect.ValueOf(i)
	case reflect.Uint:
		i, _ := strconv.Atoi(rawArg)
		v = reflect.ValueOf(uint(i))
	case reflect.Float64:
		f, _ := strconv.ParseFloat(rawArg, 64)
		v = reflect.ValueOf(f)
	case reflect.Bool:
		v = reflect.ValueOf(rawArg == "true")
	case reflect.Slice:
		// check [] at leftmost and rightmost
		if len(rawArg) <= 1 || rawArg[0] != '[' || rawArg[len(rawArg)-1] != ']' {
			err = fmt.Errorf("invalid test data: %s", rawArg)
			return
		}
		// ignore [] at leftmost and rightmost
		rawArg = rawArg[1:len(rawArg)-1]

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
		if depth != 0 {
			err = fmt.Errorf("invalid test data: %s", rawArg)
			return
		}
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

		actualOut := vFunc.Call(in)
		for i, expectedRes := range rawOut {
			if !assert.Equal(t, expectedRes, simpleValueString(actualOut[i]), "please check case %d", testCase+1) {
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
