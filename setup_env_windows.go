//go:build windows

package onnxruntime_go

// This file includes the Windows-specific code for loading the onnxruntime
// library and setting up the environment.

import (
	"syscall"
)

// #include "onnxruntime_wrapper.h"
import "C"

// This will contain the handle to the onnxruntime dll if it has been loaded
// successfully.
var libraryHandle syscall.Handle

func platformCleanup() error { _ = "STUB: not implemented"; return nil }

func platformInitializeEnvironment() error { _ = "STUB: not implemented"; return nil }

// Converts the given string to a UTF-16 string, pointed to by a raw
// *C.char. Note that we actually keep ORTCHAR_T defined to char even
// on Windows, so do _not_ index into this string from Cgo code and expect to
// get correct characters! Instead, this should only be used to obtain pointers
// that are passed to onnxruntime windows DLL functions expecting ORTCHAR_T*
// args. This is required because we undefine _WIN32 for cgo compatibility when
// including onnxruntime_c_api.h, but still interact with a DLL that was
// compiled assuming _WIN32 was defined.
//
// The pointer returned by this function must still be freed using C.free when
// no longer needed. This will return an error if the given string contains
// non-UTF8 characters.
func createOrtCharString(str string) (*C.char, error) {
	_ = "STUB: not implemented"

	// Assumed common case: the utf16 buffer contains one uint16 per utf8 byte
	// plus one more for the required null terminator in the C buffer.
	return nil, nil
}

// Convert UTF-8 to UTF-16 by reading each subsequent rune from src and
// appending it as UTF-16 to dst.

// Make sure dst contains the null terminator. Additionally this will cause
// us to return an empty string if the original string was empty.

// Finally, we need to copy dst into a C array for compatibility with
// C.CString.
