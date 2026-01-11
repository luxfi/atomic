// Copyright (c) 2020-2025 Lux Industries, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// CacheLineSize is the CPU cache line size in bytes.
// 64 bytes is the most common cache line size for modern CPUs.
const CacheLineSize = 64

// PaddedUint64 is like atomic.Uint64 but padded to prevent false sharing.
// Use this for high-contention counters accessed by multiple goroutines.
type PaddedUint64 struct {
	_ [CacheLineSize - unsafe.Sizeof(atomic.Uint64{})%CacheLineSize]byte
	atomic.Uint64
	_ [CacheLineSize - unsafe.Sizeof(atomic.Uint64{})%CacheLineSize]byte
}

// PaddedInt64 is like atomic.Int64 but padded to prevent false sharing.
type PaddedInt64 struct {
	_ [CacheLineSize - unsafe.Sizeof(atomic.Int64{})%CacheLineSize]byte
	atomic.Int64
	_ [CacheLineSize - unsafe.Sizeof(atomic.Int64{})%CacheLineSize]byte
}

// PaddedUint32 is like atomic.Uint32 but padded to prevent false sharing.
type PaddedUint32 struct {
	_ [CacheLineSize - unsafe.Sizeof(atomic.Uint32{})%CacheLineSize]byte
	atomic.Uint32
	_ [CacheLineSize - unsafe.Sizeof(atomic.Uint32{})%CacheLineSize]byte
}

// PaddedInt32 is like atomic.Int32 but padded to prevent false sharing.
type PaddedInt32 struct {
	_ [CacheLineSize - unsafe.Sizeof(atomic.Int32{})%CacheLineSize]byte
	atomic.Int32
	_ [CacheLineSize - unsafe.Sizeof(atomic.Int32{})%CacheLineSize]byte
}

// PaddedBool is like atomic.Bool but padded to prevent false sharing.
type PaddedBool struct {
	_ [CacheLineSize - unsafe.Sizeof(atomic.Bool{})%CacheLineSize]byte
	atomic.Bool
	_ [CacheLineSize - unsafe.Sizeof(atomic.Bool{})%CacheLineSize]byte
}
