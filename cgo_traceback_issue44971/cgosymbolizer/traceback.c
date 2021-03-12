// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cgo
// +build linux

#include <stddef.h>
#include <stdint.h>

struct cgoTracebackArg {
	uintptr_t  context;
	uintptr_t  sigContext;
	uintptr_t* buf;
	uintptr_t  max;
};

// Gather addresses from the call stack.
void cgoTraceback(void* parg) {
	struct cgoTracebackArg* arg = (struct cgoTracebackArg*)(parg);

	// Always provide a bogus bad address.
	//
	// Note that this should be the address of an int 3 at the end of a Go
	// function (which contains inlined functions) in order to trigger a
	// crash.
	arg->buf[0] = 0x402f6e;
}
