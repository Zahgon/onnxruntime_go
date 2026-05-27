package onnxruntime_go

// This file contains code and types that we maintain for compatibility
// purposes, but is not expected to be regularly maintained or udpated.

import (
	"fmt"
)

// #include "onnxruntime_wrapper.h"
import "C"

// DEPRECATED: This type was written with a type parameter despite the fact
// that a type parameter is not necessary for any of its underlying
// implementation. It is preserved only for compatibility with older code, and
// new users should use AdvancedSession instead. Despite the name,
// AdvancedSession is equally simple to use and far more flexible.
type Session[T TensorData] struct {
	// We now delegate all of the implementation to an AdvancedSession here.
	s *AdvancedSession
}

// DEPRECATED: See the notes on Session[T]. Use DynamicAdvancedSession instead.
type DynamicSession[In TensorData, Out TensorData] struct {
	s *DynamicAdvancedSession
}

// DEPRECATED: See the notes on Session[T]. Use NewAdvancedSessionWithONNXData
// instead.
func NewSessionWithONNXData[T TensorData](onnxData []byte, inputNames,
	outputNames []string, inputs, outputs []*Tensor[T]) (*Session[T], error) {
	_ = "STUB: not implemented"
	// Unfortunately, a slice of pointers that satisfy an interface don't count
	// as a slice of interfaces (at least, as I write this), so we'll make the
	// conversion here.
	return nil, nil
}

// DEPRECATED: See the notes on Session[T]. Use
// NewDynamicAdvancedSessionWithONNXData instead.
func NewDynamicSessionWithONNXData[in TensorData, out TensorData](onnxData []byte,
	inputNames, outputNames []string) (*DynamicSession[in, out], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DEPRECATED: See the notes on Session[T]. Use NewAdvancedSession instead.
func NewSession[T TensorData](onnxFilePath string, inputNames,
	outputNames []string, inputs, outputs []*Tensor[T]) (*Session[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DEPRECATED: See the notes on Session[T]. Use NewDynamicAdvancedSession
// instead.
func NewDynamicSession[in TensorData, out TensorData](onnxFilePath string,
	inputNames, outputNames []string) (*DynamicSession[in, out], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Session[_]) Destroy() error { _ = "STUB: not implemented"; return nil }

func (s *DynamicSession[_, _]) Destroy() error { _ = "STUB: not implemented"; return nil }

func (s *Session[T]) Run() error { _ = "STUB: not implemented"; return nil }

func (s *DynamicSession[in, out]) Run(inputs []*Tensor[in],
	outputs []*Tensor[out]) error {
	_ = "STUB: not implemented"
	return nil
}

// This type alias is included to avoid breaking older code, where the inputs
// and outputs to session.Run() were ArbitraryTensors rather than Values.
type ArbitraryTensor = Value

// As with the ArbitraryTensor type, this type alias only exists to facilitate
// renaming an old type without breaking existing code.
type TensorInternalData = ValueInternalData

var TrainingAPIRemovedError error = fmt.Errorf("Support for the training " +
	"API has been removed from onnxruntime_go following its deprecation in " +
	"onnxruntime versions 1.19.2 and later. The last revision of " +
	"onnxruntime_go supporting the training API is version v1.12.1")

// Support for TrainingSessions has been removed from onnxruntime_go following
// the deprecation of the training API in onnxruntime 1.20.0.
type TrainingSession struct{}

// Always returns TrainingAPIRemovedError.
func (s *TrainingSession) ExportModel(path string, outputNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Always returns TrainingAPIRemovedError.
func (s *TrainingSession) SaveCheckpoint(path string,
	saveOptimizerState bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Always returns TrainingAPIRemovedError.
func (s *TrainingSession) Destroy() error { _ = "STUB: not implemented"; return nil }

// Always returns TrainingAPIRemovedError.
func (s *TrainingSession) TrainStep() error { _ = "STUB: not implemented"; return nil }

// Always returns TrainingAPIRemovedError.
func (s *TrainingSession) OptimizerStep() error { _ = "STUB: not implemented"; return nil }

// Always returns TrainingAPIRemovedError.
func (s *TrainingSession) LazyResetGrad() error { _ = "STUB: not implemented"; return nil }

// Support for TrainingInputOutputNames has been removed from onnxruntime_go
// following the deprecation of the training API in onnxruntime 1.20.0.
type TrainingInputOutputNames struct {
	TrainingInputNames  []string
	EvalInputNames      []string
	TrainingOutputNames []string
	EvalOutputNames     []string
}

// Always returns (nil, TrainingAPIRemovedError).
func GetInputOutputNames(checkpointStatePath string, trainingModelPath string,
	evalModelPath string) (*TrainingInputOutputNames, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always returns false.
func IsTrainingSupported() bool {
	_ = "STUB: not implemented"

	// Always returns (nil, TrainingAPIRemovedError).
	return false
}

func NewTrainingSessionWithOnnxData(checkpointData, trainingData, evalData,
	optimizerData []byte, inputs, outputs []Value,
	options *SessionOptions) (*TrainingSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always returns (nil, TrainingAPIRemovedError).
func NewTrainingSession(checkpointStatePath, trainingModelPath, evalModelPath,
	optimizerModelPath string, inputs, outputs []Value,
	options *SessionOptions) (*TrainingSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
