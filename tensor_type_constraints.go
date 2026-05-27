package onnxruntime_go

// This file contains definitions for the generic tensor data types we support.

// #include "onnxruntime_wrapper.h"
import "C"

type FloatData interface {
	~float32 | ~float64
}

type IntData interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}

// This is used as a type constraint for the generic Tensor type.
type TensorData interface {
	FloatData | IntData | ~bool
}

// Returns the ONNX enum value used to indicate TensorData type T.
func GetTensorElementDataType[T TensorData]() C.ONNXTensorElementDataType {
	_ = "STUB: not implemented"
	// Sadly, we can't do type assertions to get underlying types, so we need
	// to use reflect here instead.
	return *new(C.ONNXTensorElementDataType)
}
