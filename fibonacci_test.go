package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectness(t *testing.T) {
	var tests = []struct {
		index,
		value int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		assert.Equal(t, test.value, Recursive(test.index), "value should match")

		InitDynamic()
		assert.Equal(t, test.value, Dynamic(test.index), "value should match")

		assert.Equal(t, test.value, Matrix(test.index), "value should match")

		assert.Equal(t, test.value, FastDouble(test.index), "value should match")
	}
}

func BenchmarkRecursive25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Recursive(25)
	}
}

func BenchmarkDynamic25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitDynamic()
		Dynamic(25)
	}
}

func BenchmarkMatrix25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Matrix(25)
	}
}

func BenchmarkFastDouble25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FastDouble(25)
	}
}

func BenchmarkDynamic100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitDynamic()
		Dynamic(100)
	}
}

func BenchmarkMatrix100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Matrix(100)
	}
}

func BenchmarkFastDouble100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FastDouble(100)
	}
}

func BenchmarkMatrix500(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Matrix(500)
	}
}

func BenchmarkFastDouble500(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FastDouble(500)
	}
}
