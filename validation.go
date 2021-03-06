package fmk

import (
	"reflect"
	"strings"
	"unicode"
)

func required(v any) bool {
	vType := reflect.TypeOf(v)
	switch {
	case vType == nil:
		return true
	case vType.Name() == "string" && v == "":
		return true
	}
	return false
}

func email(v string) bool {

	atSections := strings.Split(v, "@")
	if len(atSections) != 2 {
		return true
	}
	prefix := []rune(atSections[0])
	for i, r := range []rune(atSections[0]) {
		f := i == 0
		l := i == len(prefix)-1
		// first and last must be letters
		if (f && !unicode.IsLetter(r)) || (l && !unicode.IsLetter(r)) {
			return true
		}
	}

	domain := []rune(atSections[1])
	for i, r := range []rune(atSections[1]) {
		f := i == 0
		l := i == len(domain)-1
		// first and last must be letters
		if (f && !unicode.IsLetter(r)) || (l && !unicode.IsLetter(r)) {
			return true
		}
	}

	return false
}
