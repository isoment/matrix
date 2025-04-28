package matrix

import (
	"testing"
)

func TestAddInPlace(t *testing.T) {
	t.Run("it adds matrix modifying the original", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(6)

		want, _ := NewEmptyMatrix[int](3, 3)
		want.Fill(12)

		result, _ := m1.AddInPlace(m1)

		if !pointersAreSame(result, m1) {
			t.Error("Expected original matrix and return value to be pointed to the same struct.")
		}

		matrixesAreEqual(t, m1, want)
	})

	t.Run("it errors if the matrixes have different dimensions", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(8)

		m2, _ := NewEmptyMatrix[int](2, 2)
		m2.Fill(1)

		_, err := m1.AddInPlace(m2)

		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestSubtractInPlace(t *testing.T) {
	t.Run("it subtracts matrixes modifying the original", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(6)

		m2, _ := NewEmptyMatrix[int](3, 3)
		m2.Fill(2)

		want, _ := NewEmptyMatrix[int](3, 3)
		want.Fill(4)

		result, _ := m1.SubtractInPlace(m2)

		if !pointersAreSame(result, m1) {
			t.Error("Expected original matrix and return value to be pointed to the same struct.")
		}

		matrixesAreEqual(t, m1, want)
	})

	t.Run("it errors if the matrixes have different dimensions", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(8)

		m2, _ := NewEmptyMatrix[int](2, 2)
		m2.Fill(1)

		_, err := m1.SubtractInPlace(m2)

		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestScalarMultiplyInPlace(t *testing.T) {
	t.Run("it scalar multiples the given matrix in place", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}
		matrix, _ := NewMatrixFromSlice(input)

		input = [][]int{
			{3, 6, 9, 12},
			{15, 18, 21, 24},
			{27, 30, 33, 36},
		}
		want, _ := NewMatrixFromSlice(input)

		result := matrix.ScalarMultiplyInPlace(3)

		if !pointersAreSame(result, matrix) {
			t.Error("Expected original matrix and return value to be pointed to the same struct.")
		}

		matrixesAreEqual(t, matrix, want)
	})
}

func TestSet(t *testing.T) {
	t.Run("it sets the element to the given value at the given position", func(t *testing.T) {
		matrix, _ := NewEmptyMatrix[int](8, 8)
		matrix.Set(0, 0, 12)
		matrix.Set(6, 4, 81)

		if matrix.reader.Read(0, 0) != 12 {
			t.Errorf("expected 12 but got %d", matrix.reader.Read(0, 0))
		}

		if matrix.reader.Read(6, 4) != 81 {
			t.Errorf("expected 81 but got %d", matrix.reader.Read(6, 4))
		}
	})

	t.Run("it returns an error when the provided position is out of matrix bounds", func(t *testing.T) {
		matrix, _ := NewEmptyMatrix[int](4, 4)
		_, err := matrix.Set(6, 4, 81)
		if err == nil {
			t.Error("expected error but got none")
		}
	})
}
