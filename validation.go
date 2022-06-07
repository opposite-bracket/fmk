package fmk

import (
	"log"
	"reflect"
)

func required(v any) bool {
	log.Print(reflect.TypeOf(v))
	return false
}

func max(v any) bool {
	return false
}

func min(v any) bool {
	return false
}

func eq(v any) bool {
	return false
}

func neq(v any) bool {
	return false
}

func gt(v any) bool {
	return false
}

func lt(v any) bool {
	return false
}

func gte(v any) bool {
	return false
}

func lte(v any) bool {
	return false
}
