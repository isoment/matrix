package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	t.Run("it creates a new matrix", func(t *testing.T) {
		input := [][]int{
			{1, 2},
			{4, 5},
		}

		reader := &DefaultDataStore[int]{data: input}
		r, err := NewMatrix(reader)

		cases := []struct {
			position []uint
			expected int
		}{
			{position: []uint{0, 0}, expected: 1},
			{position: []uint{0, 1}, expected: 2},
			{position: []uint{1, 0}, expected: 4},
			{position: []uint{1, 1}, expected: 5},
		}

		for _, v := range cases {
			actualValue := r.reader.Read(v.position[0], v.position[1])
			if actualValue != v.expected {
				t.Errorf("expected %v at matrix position %v but got %v", v.expected, v.position, actualValue)
			}
		}

		if err != nil {
			t.Error("got an error but expected none")
		}
	})

	t.Run("it returns an error if the matrix dimensions are 0", func(t *testing.T) {
		t.Run("no rows", func(t *testing.T) {
			var input [][]int
			reader := &DefaultDataStore[int]{data: input}

			_, err := NewMatrix(reader)
			if err == nil {
				t.Error("expected error but got none")
			}
		})

		t.Run("no columns", func(t *testing.T) {
			input := make([][]int, 1)
			input[0] = []int{}
			reader := &DefaultDataStore[int]{data: input}

			_, err := NewMatrix(reader)
			if err == nil {
				t.Error("expected error but got none")
			}
		})
	})

	t.Run("it returns an error for column mismatch", func(t *testing.T) {
		input := [][]int{
			{0, 0, 0},
			{0, 0},
			{0},
		}
		reader := &DefaultDataStore[int]{data: input}

		_, err := NewMatrix(reader)
		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestNewMatrixFromSlice(t *testing.T) {
	t.Run("it creates a new matrix", func(t *testing.T) {
		input := [][]float32{
			{2.4, 9.5},
			{1.2, 3.45},
		}
		r, err := NewMatrixFromSlice(input)

		cases := []struct {
			position []uint
			expected float32
		}{
			{position: []uint{0, 0}, expected: 2.4},
			{position: []uint{0, 1}, expected: 9.5},
			{position: []uint{1, 0}, expected: 1.2},
			{position: []uint{1, 1}, expected: 3.45},
		}

		for _, v := range cases {
			actualValue := r.reader.Read(v.position[0], v.position[1])
			if actualValue != v.expected {
				t.Errorf("expected %v at matrix position %v but got %v", v.expected, v.position, actualValue)
			}
		}

		if err != nil {
			t.Error("got an error but expected none")
		}
	})

	t.Run("it returns an error if the matrix dimensions are 0", func(t *testing.T) {
		t.Run("no rows", func(t *testing.T) {
			var input [][]int

			_, err := NewMatrixFromSlice(input)
			if err == nil {
				t.Error("expected error but got none")
			}
		})

		t.Run("no columns", func(t *testing.T) {
			input := make([][]int, 1)
			input[0] = []int{}

			_, err := NewMatrixFromSlice(input)
			if err == nil {
				t.Error("expected error but got none")
			}
		})
	})

	t.Run("it returns an error for column mismatch", func(t *testing.T) {
		input := [][]int{
			{0, 0, 0},
			{0, 0},
			{0},
		}
		_, err := NewMatrixFromSlice(input)
		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestNewEmptyMatrix(t *testing.T) {
	t.Run("it creates a new empty matrix", func(t *testing.T) {
		got, _ := NewEmptyMatrix[int](3, 3)

		input := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		want, _ := NewMatrixFromSlice(input)

		matrixesAreEqual(t, got, want)
	})

	t.Run("it does not allow a zero element matrix", func(t *testing.T) {
		_, err := NewEmptyMatrix[int](0, 0)
		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestHasIndex(t *testing.T) {
	t.Run("it returns true if there is an index", func(t *testing.T) {
		m, _ := NewEmptyMatrix[int](3, 3)
		m.index = make(map[int][][2]uint)
		hasIndex := m.HasIndex()
		if !hasIndex {
			t.Error("expected HasIndex to return true, got false")
		}
	})

	t.Run("it returns false if there is no index", func(t *testing.T) {
		m, _ := NewEmptyMatrix[int](3, 3)
		hasIndex := m.HasIndex()
		if hasIndex {
			t.Error("expected HasIndex to return false, got true")
		}
	})
}

func TestIndex(t *testing.T) {
	t.Run("it creates an index for the matrix", func(t *testing.T) {
		input := [][]int{
			{23, 2, 7},
			{2, 9, 2},
			{9, 7, 23},
		}
		m, _ := NewMatrixFromSlice(input)

		err := m.Index()
		if err != nil {
			t.Error("got error but expected none")
		}

		want := map[int][][2]uint{
			2: {
				{0, 1},
				{1, 0},
				{1, 2},
			},
			7: {
				{0, 2},
				{2, 1},
			},
			9: {
				{1, 1},
				{2, 0},
			},
			23: {
				{0, 0},
				{2, 2},
			},
		}

		if !reflect.DeepEqual(want, m.index) {
			t.Errorf("want %v, got %v", want, m.index)
		}
	})

	t.Run("it returns an error if the matrix has an index", func(t *testing.T) {
		m, _ := NewEmptyMatrix[int](3, 3)
		m.index = make(map[int][][2]uint)
		err := m.Index()
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
