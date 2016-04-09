// Copyright (c) 2016 Uber Technologies, Inc.
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

package zap_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/uber-common/zap"
)

func withBenchedStandard(b *testing.B, f func(zap.StandardLogger)) {
	logger := zap.NewJSON(zap.All, zap.Output(ioutil.Discard))
	std, err := zap.Standardize(logger, zap.Info)
	if err != nil {
		panic(fmt.Sprint("Unexpected error creating a StandardLogger: ", err.Error()))
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f(std)
		}
	})
}

func BenchmarkStandardPrint(b *testing.B) {
	withBenchedStandard(b, func(std zap.StandardLogger) {
		std.Print("foo ", "bar ", 42, -42.0, true)
	})
}

func BenchmarkStandardPrintf(b *testing.B) {
	withBenchedStandard(b, func(std zap.StandardLogger) {
		std.Printf("%s the %s with %v, %v, and %v.", "foo ", "bar ", 42, -42.0, true)
	})
}
