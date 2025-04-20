package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	t.Run("it creates a new matrix", func(t *testing.T) {
		r, err := NewMatrix(2, 2, [][]float32{
			{2.4, 9.5},
			{1.2, 3.45},
		})

		cases := []struct {
			position []int
			expected float32
		}{
			{position: []int{0, 0}, expected: 2.4},
			{position: []int{0, 1}, expected: 9.5},
			{position: []int{1, 0}, expected: 1.2},
			{position: []int{1, 1}, expected: 3.45},
		}

		for _, v := range cases {
			actualValue := r.data[v.position[0]][v.position[1]]
			if actualValue != v.expected {
				t.Errorf("expected %v at matrix position %v but got %v", v.expected, v.position, actualValue)
			}
		}

		if err != nil {
			t.Error("got an error but expected none")
		}

	})

	t.Run("it returns an error if the matrix dimensions are 0", func(t *testing.T) {
		cases := []struct {
			rows    uint
			columns uint
		}{
			{0, 0},
			{0, 1},
			{1, 0},
		}

		for _, c := range cases {
			_, err := NewMatrix(c.rows, c.columns, [][]int{})
			if err == nil {
				t.Error("expected error but got none")
			}
		}
	})

	t.Run("it returns an error if the row count does not match the matrix rows", func(t *testing.T) {
		_, err := NewMatrix(2, 2, [][]int{
			{1, 1},
			{1, 1},
			{1, 1},
		})

		if err == nil {
			t.Error("expected error but got none")
		}
	})

	t.Run("it returns an error if the column count does not match the matrix columns", func(t *testing.T) {
		_, err := NewMatrix(2, 2, [][]int{
			{1, 1},
			{1, 1, 2},
		})

		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestNewEmptyMatrix(t *testing.T) {
	t.Run("it creates a new empty matrix", func(t *testing.T) {
		got, _ := NewEmptyMatrix[int](3, 3)

		want, _ := NewMatrix(3, 3, [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		})

		matrixesAreEqual(t, got, want)
	})

	t.Run("it does not allow a zero element matrix", func(t *testing.T) {
		_, err := NewEmptyMatrix[int](0, 0)
		if err == nil {
			t.Error("expected error but got none")
		}
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
