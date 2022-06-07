package fmk

import "testing"

func Test_eq(t *testing.T) {
	tests := []struct {
		name       string
		v          any
		shouldFail bool
	}{
		{name: "string required success", v: "a", shouldFail: false},
		{name: "string required failure", v: "", shouldFail: true},
		{name: "int required success", v: 1, shouldFail: false},
		{name: "int required failure", v: 0, shouldFail: true},
		{name: "int8 required success", v: int8(1), shouldFail: false},
		{name: "int8 required failure", v: int8(0), shouldFail: true},
		{name: "int16 required success", v: int16(1), shouldFail: false},
		{name: "int16 required failure", v: int16(0), shouldFail: true},
		{name: "int32 required success", v: int32(1), shouldFail: false},
		{name: "int32 required failure", v: int32(0), shouldFail: true},
		{name: "float64 required success", v: float64(1), shouldFail: false},
		{name: "float64 required failure", v: float64(0), shouldFail: true},
		{name: "float32 required success", v: float32(1), shouldFail: false},
		{name: "float32 required failure", v: float32(0), shouldFail: true},
		{name: "float64 required success", v: float64(1), shouldFail: false},
		{name: "float64 required failure", v: float64(0), shouldFail: true},
		//{name: "bool required"},
		//{name: "complex32 required"},
		//{name: "complex64 required"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eq(tt.v); got != tt.shouldFail {
				t.Errorf("eq() = %v, shouldFail %v", got, tt.shouldFail)
			}
		})
	}
}

//
//func Test_gt(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := gt(tt.args.v); got != tt.shouldFail {
//				t.Errorf("gt() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_gte(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := gte(tt.args.v); got != tt.shouldFail {
//				t.Errorf("gte() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_lt(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := lt(tt.args.v); got != tt.shouldFail {
//				t.Errorf("lt() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_lte(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := lte(tt.args.v); got != tt.shouldFail {
//				t.Errorf("lte() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_max(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := max(tt.args.v); got != tt.shouldFail {
//				t.Errorf("max() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_min(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := min(tt.args.v); got != tt.shouldFail {
//				t.Errorf("min() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_neq(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := neq(tt.args.v); got != tt.shouldFail {
//				t.Errorf("neq() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
//
//func Test_required(t *testing.T) {
//	tests := []struct {
//		name string
//		v any
//		shouldFail bool
//	}{
//		{name: "string required"},
//		{name: "int required"},
//		{name: "int8 required"},
//		{name: "int16 required"},
//		{name: "int32 required"},
//		{name: "int64 required"},
//		{name: "bool required"},
//		{name: "float32 required"},
//		{name: "float64 required"},
//		{name: "complex32 required"},
//		{name: "complex64 required"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := required(tt.args.v); got != tt.shouldFail {
//				t.Errorf("required() = %v, shouldFail %v", got, tt.shouldFail)
//			}
//		})
//	}
//}
