package util

import "fmt"

func Assert(condition bool, format string, args ...any) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+format, args...))
	}
}
