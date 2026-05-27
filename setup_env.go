//go:build !windows

package onnxruntime_go

import (
	"unsafe"
)

/*
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
#include "onnxruntime_wrapper.h"

typedef OrtApiBase* (*GetOrtApiBaseFunction)(void);

// Since Go can't call C function pointers directly, we just use this helper
// when calling GetApiBase
OrtApiBase *CallGetAPIBaseFunction(void *fn) {
	OrtApiBase *to_return = ((GetOrtApiBaseFunction) fn)();
	return to_return;
}
*/
import "C"

// This file includes the code for loading the onnxruntime and setting up the
// environment on non-Windows systems. For now, it has been tested on Linux and
// arm64 OSX.

// This will contain the handle to the onnxruntime shared library if it has
// been loaded successfully.
var libraryHandle unsafe.Pointer

func platformCleanup() error { _ = "STUB: not implemented"; return nil }

// Should only be called on Apple systems; looks up the CoreML provider
// function which should only be exported on apple onnxruntime dylib files.
func setAppendCoreMLFunctionPointer(libraryHandle unsafe.Pointer) error {
	_ = "STUB: not implemented"
	// This function name must match the name in coreml_provider_factory.h,
	// which is provided in the onnxruntime release's include/ directory on for
	// Apple platforms.
	return nil
}

func platformInitializeEnvironment() error { _ = "STUB: not implemented"; return nil }

// We'll silently ignore potential errors returned by
// setAppendCoreMLFunctionPointer (for now at least). Even though we're
// on Apple hardware, it's possible that the user will have compiled
// the onnxruntime library from source without CoreML support.
// A failure here will only leave the coreml function pointer as NULL
// in our C code, which will be detected and result in an error at
// runtime.

// Converts the given path to an ORTCHAR_T string, pointed to by a *C.char. The
// returned string must be freed using C.free when no longer needed. This
// wrapper is used for source compatibility with onnxruntime API functions
// requiring paths, which must be UTF-16 on Windows but UTF-8 elsewhere.
func createOrtCharString(str string) (*C.char, error) { _ = "STUB: not implemented"; return nil, nil }
