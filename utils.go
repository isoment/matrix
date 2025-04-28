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

	aRows, aColumns := a.reader.Shape()
	bRows, bColumns := a.reader.Shape()

	if aRows != bRows {
		return false
	}

	if aColumns != bColumns {
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
			m.writer.Write(i, j, v)
		}
	}
	return m
}

func (m Matrix[T]) Clone() (*Matrix[T], error) {
	new, err := NewEmptyMatrix[T](m.rows, m.columns)
	if err != nil {
		return nil, err
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			new.writer.Write(i, j, m.reader.Read(i, j))
		}
	}

	return new, nil
}

/*
Verify that each row has the same number of columns ensuring that the structure
is a valid matrix
*/
func verifyColumnCount[T Element](data [][]T) error {
	columnsCount := uint(len(data[0]))

	for i, row := range data {
		if uint(len(row)) != columnsCount {
			return ErrColumnCountMismatch(i)
		}
	}

	return nil
}
