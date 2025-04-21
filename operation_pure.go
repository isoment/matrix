package matrix

/*
Perform an addition operation on two matrixes creating a new matrix with the results
*/
func (m Matrix[T]) Add(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(&m, a) {
		return nil, ErrMustBeSameDimensions
	}

	rows := uint(len(m.data))
	columns := uint(len(m.data[0]))
	result := createEmptyMatrix[T](rows, columns)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			result.data[i][j] = m.data[i][j] + a.data[i][j]
		}
	}

	return &result, nil
}

/*
Performs scalar multiplication on a matrix returning a new result matrix
*/
func (m Matrix[T]) ScalarMultiply(c T) *Matrix[T] {
	result := createEmptyMatrix[T](m.rows, m.columns)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			result.data[i][j] = c * m.data[i][j]
		}
	}

	return &result
}

/*
Search the given matrix for an element. Returns a list of Location with
position and values and boolean noting if it was found or not.
*/
func (m Matrix[T]) Search(element T) ([]Location[T], bool) {
	var found []Location[T]

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			if m.data[i][j] == element {
				el := Location[T]{
					position: [2]uint{i, j},
					value:    m.data[i][j],
				}
				found = append(found, el)
			}
		}
	}

	if len(found) == 0 {
		return found, false
	}

	return found, true
}

func (m Matrix[T]) Transpose() *Matrix[T] {
	new := createEmptyMatrix[T](m.columns, m.rows)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			new.data[j][i] = m.data[i][j]
		}
	}

	return &new
}
