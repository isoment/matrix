package matrix

import (
	"reflect"
	"testing"
)

// Create a spy implementation satisfying the DataReader interface
type SpyReader[T Element] struct {
	realReader   DataReader[T]
	readDetected bool
}

// DataReader interface requires a Read method. We can use readDetected to tell if it was called.
func (s *SpyReader[T]) Read(i, j uint) T {
	s.readDetected = true
	return s.realReader.Read(i, j)
}

func (s *SpyReader[T]) Shape() (rows, columns uint) {
	return s.realReader.Shape()
}

func (s *SpyReader[T]) Validate() error {
	return s.realReader.Validate()
}

func TestAdd(t *testing.T) {
	t.Run("it sums a matrix", func(t *testing.T) {
		input := [][]float32{
			{1, 2, 3},
			{10, 25, 50},
			{99, 5, 32},
		}
		m1, _ := NewMatrixFromSlice(input)

		m2, _ := m1.Clone()

		input = [][]float32{
			{2, 4, 6},
			{20, 50, 100},
			{198, 10, 64},
		}
		want, _ := NewMatrixFromSlice(input)

		got, err := m1.Add(m2)
		if err != nil {
			t.Error(err)
		}

		matrixesAreEqual(t, want, got)
	})

	t.Run("it errors if the matrixes have different dimensions", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[float32](3, 3)
		m1.Fill(3)

		m2, _ := NewEmptyMatrix[float32](2, 2)
		m2.Fill(2)

		_, err := m1.Add(m2)

		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestScalarMultiply(t *testing.T) {
	t.Run("it scalar multiplies the given matrix", func(t *testing.T) {
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

		result, err := matrix.ScalarMultiply(3)
		if err != nil {
			t.Error(err)
		}

		matrixesAreEqual(t, result, want)
	})
}

func TestSearch(t *testing.T) {
	t.Run("it searches and returns a single found element", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}
		matrix, _ := NewMatrixFromSlice(input)
		search := 10

		s, found := matrix.Search(search)

		if !found {
			t.Error("expected found to be true, got false")
		}

		if len(s) == 0 {
			t.Error("expected a found element, got none")
		}

		e := s[0]

		if e.position != [2]uint{2, 1} || e.value != search {
			t.Errorf("got %+v which is not correct", s)
		}
	})

	t.Run("it searches and returns multiple found elements", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{2, 4, 2},
		}
		matrix, _ := NewMatrixFromSlice(input)
		search := 2

		s, found := matrix.Search(search)

		if !found {
			t.Error("expected found to be true, got false")
		}

		if len(s) != 3 {
			t.Error("expected 3 found element, got none")
		}

		for _, v := range s {
			matrixElement := matrix.reader.Read(v.position[0], v.position[1])
			if matrixElement != v.value {
				t.Errorf("expected %v at position %v but got %v", v.value, v.position, matrixElement)
			}
		}
	})

	t.Run("it returns false if element not found", func(t *testing.T) {
		matrix, _ := NewEmptyMatrix[int8](3, 3)
		matrix.Fill(31)
		search := int8(55)

		_, found := matrix.Search(search)

		if found {
			t.Error("expected found to be false, got true")
		}
	})

	t.Run("it searches using index without accessing data field", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{2, 4, 2},
		}
		store := &DefaultDataStore[int]{data: input}
		spy := &SpyReader[int]{realReader: store}

		matrix, _ := NewMatrix(spy)

		matrix.index = map[int][][2]uint{
			2: {{0, 1}, {2, 0}, {2, 2}},
		}

		matrix.reader = spy

		search := 2
		result, found := matrix.Search(search)

		if !found {
			t.Error("expected found to be true, got false")
		}

		if len(result) != 3 {
			t.Errorf("expected 3 found elements, got %d", len(result))
		}

		if spy.readDetected {
			t.Error("expected no read from data, but data was accessed")
		}
	})
}

func TestTranspose(t *testing.T) {
	t.Run("it transposes the matrix", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}
		matrix, _ := NewMatrixFromSlice(input)

		input = [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		}
		want, _ := NewMatrixFromSlice(input)

		got, err := matrix.Transpose()
		if err != nil {
			t.Error(err)
		}

		matrixesAreEqual(t, got, want)
	})
}

func TestSubtract(t *testing.T) {
	t.Run("it subtracts matrixes", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(4)

		m2, _ := m1.Clone()
		m2.Fill(2)

		got, _ := m1.Subtract(m2)

		matrixesAreEqual(t, got, m2)
	})

	t.Run("it returns an error if matrixes have different dimensions", func(t *testing.T) {
		m1, _ := NewEmptyMatrix[int](3, 3)
		m1.Fill(4)

		m2, _ := NewEmptyMatrix[int](2, 2)
		m2.Fill(4)

		_, err := m1.Subtract(m2)

		if err == nil {
			t.Error("expected error but got none")
		}
	})
}

func TestFlatten(t *testing.T) {
	t.Run("it flattens a matrix into a slice", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{44, 5, 6},
			{13, 4, 98},
		}
		matrix, _ := NewMatrixFromSlice(input)

		want := []int{1, 2, 3, 44, 5, 6, 13, 4, 98}

		got := matrix.Flatten()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}

func TestExpand(t *testing.T) {
	t.Run("it expands a slice into a new matrix", func(t *testing.T) {
		s := []int{1, 2, 3, 44, 5, 6, 13, 4, 98}

		input := [][]int{
			{1, 2, 3},
			{44, 5, 6},
			{13, 4, 98},
		}
		want, _ := NewMatrixFromSlice(input)

		got, _ := ExpandSliceToMatrix(s, 3, 3)

		matrixesAreEqual(t, got, want)
	})

	t.Run("it returns an error if the slice cannot fit in the matrix dimensions", func(t *testing.T) {
		s := []int{1, 2, 3, 44, 5, 6, 13, 4, 98}

		_, err := ExpandSliceToMatrix(s, 2, 2)

		if err == nil {
			t.Error("expected error but got none")
		}
	})

	t.Run("it zero fills if the slice is smaller than the new matrix", func(t *testing.T) {
		s := []int{1, 2, 3, 44, 5, 6, 8}

		input := [][]int{
			{1, 2, 3},
			{44, 5, 6},
			{8, 0, 0},
		}
		want, _ := NewMatrixFromSlice(input)

		got, _ := ExpandSliceToMatrix(s, 3, 3)

		matrixesAreEqual(t, got, want)
	})
}
