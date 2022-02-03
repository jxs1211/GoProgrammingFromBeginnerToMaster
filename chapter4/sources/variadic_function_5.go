package main

import (
	"fmt"
	"strings"
)

func concat(sep string, args ...interface{}) string {
	var result string
	for i, v := range args {
		if i != 0 {
			result += sep
		}
		switch v.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			result += fmt.Sprintf("%d", v)
		case string:
			result += fmt.Sprintf("%s", v)
		case []int:
			ints := v.([]int)
			for i, v := range ints {
				if i != 0 {
					result += sep
				}
				result += fmt.Sprintf("%d", v)
			}
		case []string:
			strs := v.([]string)
			result += strings.Join(strs, sep)
		default:
			fmt.Printf("the argument type [%T] is not supported", v)
			return ""
		}
	}
	return result
}

func Show() {
	println(concat("-", 1, 2))
	println(concat("-", "hello", "gopher"))
	println(concat("-", "hello", 1, uint32(2),
		[]int{11, 12, 13}, 17,
		[]string{"robot", "ai", "ml"},
		"hacker", 33))
}

func handle(res *string, index int, sep string, a ...interface{}) {
	if index == len(a)-1 {
		*res += fmt.Sprintf("%v", a[index])
	} else {
		*res += fmt.Sprintf("%v%s", a[index], sep)
	}
}

func concat2(sep string, args ...interface{}) string {
	var res string
	for i, v := range args {
		switch v.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			res += fmt.Sprintf("%d%s", v, sep)
			// handle(&res, i, sep, args)
		case []int:
			ints := v.([]int)
			for _, v := range ints {
				res += fmt.Sprintf("%d%s", v, sep)
			}
			// concat2(sep, v)
		case string:
			s := v.(string)
			for _, v := range s {
				res += fmt.Sprintf("%s%s", string(v), sep)
			}
		case []string:
			strs := v.([]string)
			res += strings.Join(strs, sep)
			res += sep
		default:
			fmt.Printf("unknown type err: %v\n", v)
			return ""
		}
		if i == len(args)-1 && string(res[len(res)-1]) == sep {
			res = strings.TrimRight(res, sep)
		}
	}
	return res
}

func Show2() {
	println(concat2("中",
		1, 2, 3,
		[]int{1, 2, 2, 3},
		[]string{"shen", "xian", "jie"},
		"shen",
	))
}

func main() {
	Show2()
}
