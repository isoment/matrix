package matrix

import "testing"

func TestAdd(t *testing.T) {
	t.Run("it sums 2x2 matrix", func(t *testing.T) {
		m1, _ := NewMatrix(2, 2, [][]int{
			{1, 2},
			{3, 4},
		})

		m2 := m1.Clone()

		want, _ := NewMatrix(2, 2, [][]int{
			{2, 4},
			{6, 8},
		})

		got := m1.Add(m2)

		matrixesAreEqual(t, want, got)
	})

	t.Run("it sums a 2x3 matrix", func(t *testing.T) {
		m1, _ := NewMatrix(2, 3, [][]float32{
			{1, 2, 3},
			{4, 5, 6},
		})

		m2 := m1.Clone()

		want, _ := NewMatrix(2, 3, [][]float32{
			{2, 4, 6},
			{8, 10, 12},
		})

		got := m1.Add(m2)

		matrixesAreEqual(t, want, got)
	})
}

func TestScalarMultiply(t *testing.T) {
	t.Run("it scalar multiplies the given matrix", func(t *testing.T) {
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

		result := matrix.ScalarMultiply(3)

		matrixesAreEqual(t, result, want)
	})
}
