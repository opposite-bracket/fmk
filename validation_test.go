package fmk

import "testing"

func Test_required(t *testing.T) {
	tests := []struct {
		name       string
		v          any
		shouldFail bool
	}{
		{name: "string required success", v: "a", shouldFail: false},
		{name: "string required failure", v: "", shouldFail: true},
		//{name: "int required success", v: 1, shouldFail: false},
		//{name: "int required failure", v: 0, shouldFail: false},
		//{name: "int8 required success", v: int8(1), shouldFail: false},
		//{name: "int8 required failure", v: int8(0), shouldFail: true},
		//{name: "int16 required success", v: int16(1), shouldFail: false},
		//{name: "int16 required failure", v: int16(0), shouldFail: true},
		//{name: "int32 required success", v: int32(1), shouldFail: false},
		//{name: "int32 required failure", v: int32(0), shouldFail: true},
		//{name: "float64 required success", v: float64(1), shouldFail: false},
		//{name: "float64 required failure", v: float64(0), shouldFail: true},
		//{name: "float32 required success", v: float32(1), shouldFail: false},
		//{name: "float32 required failure", v: float32(0), shouldFail: true},
		//{name: "float64 required success", v: float64(1), shouldFail: false},
		//{name: "float64 required failure", v: float64(0), shouldFail: true},
		//{name: "bool required"},
		//{name: "complex32 required"},
		//{name: "complex64 required"},
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
		{name: "success", v: "sample@sample.com", shouldFail: false},
		{name: "success", v: "sample@@sample.com", shouldFail: true},
		{name: "success", v: "sample@sam@ple.com", shouldFail: true},
		{name: "success", v: "samplesample.com", shouldFail: true},
		{name: "success", v: "sample@-sample.com", shouldFail: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := email(tt.v); got != tt.shouldFail {
				t.Errorf("email() = %v, shouldFail %v", got, tt.shouldFail)
			}
		})
	}
}
