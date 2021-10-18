// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go_bench

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func fibonacci(n int) int {
	prev2 := 0
	prev1 := 1
	if n == 0 {
		return prev2
	}
	if n == 1 {
		return prev1
	}
	var next int
	for i := 2; i <= n; i++ {
		next = prev1 + prev2
		prev2 = prev1
		prev1 = next
	}

	return next
}

func TestFibonacci(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    1,
			want: 1,
		},
		{
			n:    2,
			want: 1,
		},
		{
			n:    3,
			want: 2,
		},
		{
			n:    14,
			want: 377,
		},
	} {
		t.Run(strconv.Itoa(tc.n), func(t *testing.T) {
			if got := fibonacci(tc.n); got != tc.want {
				t.Errorf("fibonacci(%d) = %d want %d", tc.n, got, tc.want)
			}
		})
	}
}

var a int = 64

func BenchmarkConsistent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(a)
	}
}

func DisabledBenchmarkMap(b *testing.B) {
	type V *int
	value := reflect.ValueOf((V)(nil))
	stringKeys := []string{}
	mapOfStrings := map[string]V{}
	uint64Keys := []uint64{}
	mapOfUint64s := map[uint64]V{}
	for i := 0; i < 100; i++ {
		stringKey := fmt.Sprintf("key%d", i)
		stringKeys = append(stringKeys, stringKey)
		mapOfStrings[stringKey] = nil

		uint64Key := uint64(i)
		uint64Keys = append(uint64Keys, uint64Key)
		mapOfUint64s[uint64Key] = nil
	}

	tests := []struct {
		label          string
		m, keys, value reflect.Value
	}{
		{"StringKeys", reflect.ValueOf(mapOfStrings), reflect.ValueOf(stringKeys), value},
		{"Uint64Keys", reflect.ValueOf(mapOfUint64s), reflect.ValueOf(uint64Keys), value},
	}

	for _, tt := range tests {
		b.Run(tt.label, func(b *testing.B) {
			b.Run("MapIndex", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					for j := tt.keys.Len() - 1; j >= 0; j-- {
						tt.m.MapIndex(tt.keys.Index(j))
					}
				}
			})
			b.Run("SetMapIndex", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					for j := tt.keys.Len() - 1; j >= 0; j-- {
						tt.m.SetMapIndex(tt.keys.Index(j), tt.value)
					}
				}
			})
		})
	}
}
