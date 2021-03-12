// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdint.h>

// Initialize the backtrace state.
void cgoSymbolizerInit(char* filename) {
}

struct cgoSymbolizerArg {
	uintptr_t   pc;
	const char* file;
	uintptr_t   lineno;
	const char* func;
	uintptr_t   entry;
	uintptr_t   more;
	uintptr_t   data;
};

struct cgoSymbolizerMore {
	struct cgoSymbolizerMore *more;

	const char* file;
	uintptr_t   lineno;
	const char* func;
};

// For the details of how this is called see runtime.SetCgoTraceback.
void cgoSymbolizer(void* parg) {
}
