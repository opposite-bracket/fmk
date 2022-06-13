package fmk

import (
	"log"
	"reflect"
)

func required(v any) bool {
	log.Printf("--------> type %s", reflect.TypeOf(v))
	switch reflect.TypeOf(v) {
	case nil:
		return true
	}
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
