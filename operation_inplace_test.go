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
		matrix, _ := NewMatrix(3, 4, [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		})

		want, _ := NewMatrix(3, 4, [][]int{
			{3, 6, 9, 12},
			{15, 18, 21, 24},
			{27, 30, 33, 36},
		})

		result := matrix.ScalarMultiplyInPlace(3)

		if !pointersAreSame(result, matrix) {
			t.Error("Expected original matrix and return value to be pointed to the same struct.")
		}

		matrixesAreEqual(t, matrix, want)
	})
}
