// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"runtime/pprof"

	_ "github.com/ianlancetaylor/cgosymbolizer"
)

// static inline void function() {
//   for (volatile int i = 0; i < 100000; i++) {
//   }
// }
import "C"

func main() {
	f, err := os.Create("/tmp/cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 10000; i++ {
		C.function()
	}
}
