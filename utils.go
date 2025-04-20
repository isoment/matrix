package matrix

func (m *Matrix[T]) Zero() *Matrix[T] {
	var zero T
	return m.Fill(zero)
}

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
