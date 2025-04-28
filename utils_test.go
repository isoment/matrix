package matrix

import "testing"

func TestZero(t *testing.T) {
	t.Run("it zeros the given matrix", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{1, 11, 1, 41},
		}
		matrix, _ := NewMatrixFromSlice(input)

		input = [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		want, _ := NewMatrixFromSlice(input)

		matrix.Zero()

		matrixesAreEqual(t, matrix, want)
	})
}

func TestFill(t *testing.T) {
	t.Run("it fills the given matrix with a value", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{1, 11, 1, 41},
		}
		matrix, _ := NewMatrixFromSlice(input)

		input = [][]int{
			{11, 11, 11, 11},
			{11, 11, 11, 11},
			{11, 11, 11, 11},
		}
		want, _ := NewMatrixFromSlice(input)

		matrix.Fill(11)

		matrixesAreEqual(t, matrix, want)
	})
}

func TestClone(t *testing.T) {
	t.Run("it clones the given matrix", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{1, 1, 1, 1},
		}
		matrix, _ := NewMatrixFromSlice(input)

		new, _ := matrix.Clone()

		matrixesAreEqual(t, matrix, new)

		// We want to compare the address of the structs and ensure they are different
		if pointersAreSame(&matrix, &new) {
			t.Error("new matrix is the same as the original")
		}
	})
}

func TestAreSameDimensions(t *testing.T) {
	t.Run("it returns true if the matrixes are the same dimensions", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(1)

		m2, _ := NewEmptyMatrix[int](3, 3)
		m2.Fill(2)

		same := AreSameDimensions(m1, m2)

		if !same {
			t.Error("expected true, got false")
		}
	})

	t.Run("it returns false if matrixes have different dimensions", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(1)

		m2, _ := NewEmptyMatrix[int](6, 6)
		m2.Fill(2)

		same := AreSameDimensions(m1, m2)

		if same {
			t.Error("expected false but got true")
		}
	})
}
