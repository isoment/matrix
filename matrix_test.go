package matrix

import (
	"reflect"
	"testing"
)

func TestNewEmptyMatrix(t *testing.T) {
	t.Run("it creates a new empty matrix", func(t *testing.T) {
		got := NewEmptyMatrix[int](3, 3)

		want, _ := NewMatrix(3, 3, [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		})

		matrixesAreEqual(t, got, want)
	})
}

func pointersAreSame[T any](a, b *T) bool {
	if reflect.ValueOf(a).Pointer() == reflect.ValueOf(b).Pointer() {
		return true
	}
	return false
}

func matrixesAreEqual[T Element](t *testing.T, a, b *Matrix[T]) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v does not equal %+v", a, b)
	}
}
