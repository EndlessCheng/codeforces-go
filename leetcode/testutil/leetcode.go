package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func parseRawArg(tp reflect.Type, rawArg string) (v reflect.Value) {
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
		v = reflect.New(tp).Elem()
		isStringSlice := strings.Contains(rawArg, `"`)
		quotCnt := 0
		// ignore [] at leftmost and rightmost
		for start, depth := 1, 0; start < len(rawArg)-1; {
			end := start
		outer:
			for ; end < len(rawArg)-1; end++ {
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
			v = reflect.Append(v, parseRawArg(tp.Elem(), rawArg[start:end]))
			start = end + 1 // skip ,
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

func RunLeetCodeFunc(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string) error {
	tp := reflect.TypeOf(f)
	if tp.Kind() != reflect.Func {
		return fmt.Errorf("f must be a function")
	}

	vFunc := reflect.ValueOf(f)
	for testCase, rawIn := range rawInputs {
		if len(rawIn) != tp.NumIn() {
			return fmt.Errorf("len(rawIn) is not %d", tp.NumIn())
		}
		rawOut := rawOutputs[testCase]
		if len(rawOut) != tp.NumOut() {
			return fmt.Errorf("len(rawOut) is not %d", tp.NumOut())
		}

		in := make([]reflect.Value, len(rawIn))
		for i, rawArg := range rawIn {
			in[i] = parseRawArg(tp.In(i), rawArg)
		}
		actualOut := vFunc.Call(in)
		for i, expectedRes := range rawOut {
			assert.Equal(t, expectedRes, simpleValueString(actualOut[i]), "failed at example %d", testCase+1)
		}
	}
	return nil
}
