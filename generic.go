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
	"encoding/json"
	"sync"
)

// Atomic is a generic thread-safe wrapper for any type.
//
// For primitive types (bool, int32, int64, uint32, uint64), prefer using
// the stdlib sync/atomic types directly (atomic.Bool, atomic.Int64, etc.)
// as they are more efficient and use hardware atomics.
//
// Use Atomic[T] for:
//   - Complex types (structs, slices, maps, interfaces)
//   - Types that need JSON serialization
//   - When you need a generic container
//
// Example:
//
//	var config Atomic[Config]
//	config.Store(Config{Name: "default"})
//	c := config.Load()
type Atomic[T any] struct {
	_ nocmp // disallow non-atomic comparison

	mu    sync.RWMutex
	value T
}

// NewAtomic creates a new Atomic with the given initial value.
func NewAtomic[T any](value T) *Atomic[T] {
	return &Atomic[T]{value: value}
}

// Load atomically returns the current value.
// This is the preferred method name (matches sync/atomic).
func (a *Atomic[T]) Load() T {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.value
}

// Store atomically sets the value.
// This is the preferred method name (matches sync/atomic).
func (a *Atomic[T]) Store(value T) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.value = value
}

// Swap atomically sets the value and returns the old value.
func (a *Atomic[T]) Swap(value T) T {
	a.mu.Lock()
	defer a.mu.Unlock()
	old := a.value
	a.value = value
	return old
}

// Get is an alias for Load for backward compatibility.
func (a *Atomic[T]) Get() T {
	return a.Load()
}

// Set is an alias for Store for backward compatibility.
func (a *Atomic[T]) Set(value T) {
	a.Store(value)
}

// MarshalJSON implements json.Marshaler.
func (a *Atomic[T]) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *Atomic[T]) UnmarshalJSON(b []byte) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return json.Unmarshal(b, &a.value)
}
