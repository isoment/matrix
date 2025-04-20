package matrix

import "testing"

func TestAdd(t *testing.T) {
	t.Run("it sums a matrix", func(t *testing.T) {
		m1, _ := NewMatrix(3, 3, [][]float32{
			{1, 2, 3},
			{10, 25, 50},
			{99, 5, 32},
		})

		m2 := m1.Clone()

		want, _ := NewMatrix(3, 3, [][]float32{
			{2, 4, 6},
			{20, 50, 100},
			{198, 10, 64},
		})

		got, _ := m1.Add(m2)

		matrixesAreEqual(t, want, got)
	})

	t.Run("it errors if the matrixes have different dimensions", func(t *testing.T) {
		m1, _ := NewMatrix(3, 3, [][]float32{
			{1, 2, 3},
			{10, 25, 50},
			{99, 5, 32},
		})

		m2, _ := NewMatrix(2, 2, [][]float32{
			{1, 2},
			{1, 2},
		})

		_, err := m1.Add(m2)

		if err == nil {
			t.Error("expected error but got none")
		}
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
