package matrix

import (
	"testing"
)

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

func TestIsIdentityMatrix(t *testing.T) {
	t.Run("it returns false if matrix is not square", func(t *testing.T) {
		m, _ := NewEmptyMatrix[int](5, 3)
		m.Fill(4)

		r := m.IsIdentityMatrix()

		if r {
			t.Error("expected false but got true")
		}
	})

	t.Run("it returns true if the matrix is an identity matrix", func(t *testing.T) {
		input := [][]int{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		}
		m, _ := NewMatrixFromSlice(input)

		r := m.IsIdentityMatrix()

		if !r {
			t.Error("expected true but got false")
		}
	})

	t.Run("it returns false if the matrix is not an identity matrix", func(t *testing.T) {
		cases := []struct {
			name  string
			input [][]int
		}{
			{
				name: "case1",
				input: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			{
				name: "case2",
				input: [][]int{
					{1, 3, 2},
					{5, 1, 2},
					{6, 3, 1},
				},
			},
			{
				name: "case3",
				input: [][]int{
					{1, 5, 5},
					{5, 1, 5},
					{5, 5, 5},
				},
			},
		}

		for _, test := range cases {
			t.Run(test.name, func(t *testing.T) {
				m, _ := NewMatrixFromSlice(test.input)
				r := m.IsIdentityMatrix()

				if r {
					t.Error("expected false but got true")
				}
			})
		}
	})
}
