package main

import (
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func lastSubstring(s string) string {
	sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").UnsafeAddr()))
	return s[sa[len(sa)-1]:]
}
