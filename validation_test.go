package fmk

import "testing"

func Test_required(t *testing.T) {
	tests := []struct {
		name       string
		v          any
		shouldFail bool
	}{
		{name: "string requirement is fulfilled", v: "a", shouldFail: false},
		{name: "string cannot be empty", v: "", shouldFail: true},
		{name: "value cannot be null", v: nil, shouldFail: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := required(tt.v); got != tt.shouldFail {
				t.Errorf("required() = %v, shouldFail %v", got, tt.shouldFail)
			}
		})
	}
}

func Test_email(t *testing.T) {
	tests := []struct {
		name       string
		v          string
		shouldFail bool
	}{
		{name: "email requirement fulfilled", v: "sample@sample.com", shouldFail: false},
		{name: "sample@@sample.com", v: "sample@@sample.com", shouldFail: true},
		{name: "sample@sam@ple.com", v: "sample@sam@ple.com", shouldFail: true},
		{name: "no @", v: "samplesample.com", shouldFail: true},

		{name: "prefix starts with .", v: ".sample@sample.com", shouldFail: true},
		{name: "prefix ends with .", v: "sample.@sample.com", shouldFail: true},
		{name: "prefix starts with -", v: "-sample@sample.com", shouldFail: true},
		{name: "prefix ends with -", v: "sample-@sample.com", shouldFail: true},
		{name: "prefix starts with _", v: "_sample@sample.com", shouldFail: true},
		{name: "prefix ends with _", v: "sample_@sample.com", shouldFail: true},
		{name: "prefix starts with #", v: "#sample@sample.com", shouldFail: true},
		{name: "prefix ends with #", v: "sample#@sample.com", shouldFail: true},
		{name: "prefix starts with 1", v: "1sample@sample.com", shouldFail: true},
		{name: "prefix ends with 1", v: "sample1@sample.com", shouldFail: true},

		{name: "domain starts with .", v: "sample@.sample.com", shouldFail: true},
		{name: "domain ends with .", v: "sample@sample.com.", shouldFail: true},
		{name: "domain starts with -", v: "sample@-sample.com", shouldFail: true},
		{name: "domain ends with -", v: "sample@sample.com-", shouldFail: true},
		{name: "domain starts with _", v: "sample@_sample.com_", shouldFail: true},
		{name: "domain ends with _", v: "sample@sample.com_", shouldFail: true},
		{name: "domain starts with #", v: "sample@#sample.com", shouldFail: true},
		{name: "domain ends with #", v: "sample@sample.com#", shouldFail: true},
		{name: "domain starts with 1", v: "sample@1sample.com", shouldFail: true},
		{name: "domain ends with 1", v: "sample@sample.com1", shouldFail: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := email(tt.v); got != tt.shouldFail {
				t.Errorf("email() = %v, shouldFail %v", got, tt.shouldFail)
			}
		})
	}
}
