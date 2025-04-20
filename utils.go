package matrix

/*
Check if the given matrixes have the same dimensions. Check...
- For row, column count equality
- If they have the same number of rows
- If they have the same number of columns
*/
func AreSameDimensions[T Element](a, b *Matrix[T]) bool {
	if a.rows != b.rows || a.columns != b.columns {
		return false
	}

	if len(a.data) != len(b.data) {
		return false
	}

	if len(a.data[0]) != len(b.data[0]) {
		return false
	}

	return true
}

func (m *Matrix[T]) Zero() *Matrix[T] {
	var zero T
	return m.Fill(zero)
}

/*
Fill a matrix with a given value element
*/
func (m *Matrix[T]) Fill(v T) *Matrix[T] {
	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			m.data[i][j] = v
		}
	}
	return m
}

func (m Matrix[T]) Clone() *Matrix[T] {
	new := createEmptyMatrix[T](m.rows, m.columns)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			new.data[i][j] = m.data[i][j]
		}
	}

	return &new
}

func createEmptyMatrix[T Element](rows, columns uint) Matrix[T] {
	matrix := make([][]T, rows)

	for h := uint(0); h < rows; h++ {
		matrix[h] = make([]T, columns)
	}

	return Matrix[T]{
		rows:    rows,
		columns: columns,
		data:    matrix,
	}
}
