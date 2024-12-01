package matrix

import "math"

type numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Matrix[T numeric] struct {
	Rows    int
	Columns int
	Data    []T
}

func New[T numeric](rows int, columns int, data ...T) *Matrix[T] {
	m := &Matrix[T]{
		Rows:    rows,
		Columns: columns,
		Data:    make([]T, rows*columns),
	}

	n := rows * columns
	if n > len(data) {
		n = len(data)
	}
	for i := 0; i < n; i++ {
		m.Data[i] = data[i]
	}

	return m
}

func Identity[T numeric](size int) *Matrix[T] {
	m := New[T](size, size)

	for i := 0; i < size; i++ {
		m.Data[i*size+i] = T(1)
	}

	return m
}

func (m *Matrix[T]) At(row, column int) T {
	if row < 0 || column < 0 || row >= m.Rows || column >= m.Columns {
		panic("matrix indexes out of bounds")
	}
	return m.Data[row*m.Columns+column]
}

func (m *Matrix[T]) Set(row, column int, value T) {
	if row < 0 || column < 0 || row >= m.Rows || column >= m.Columns {
		panic("matrix indexes out of bounds")
	}
	m.Data[row*m.Columns+column] = value
}

func (m *Matrix[T]) Scale(s T) *Matrix[T] {
	m2 := &Matrix[T]{
		Rows:    m.Rows,
		Columns: m.Columns,
		Data:    make([]T, len(m.Data)),
	}

	for i, v := range m.Data {
		m2.Data[i] = v * s
	}

	return m2
}

func (m *Matrix[T]) Len() T {
	if m.Columns != 1 {
		panic("matrix size for len must have 1 column")
	}

	s := T(0)
	for _, v := range m.Data {
		s += v * v
	}
	return T(math.Sqrt(float64(s)))
}

func (m *Matrix[T]) DotProduct(m2 *Matrix[T]) T {
	if m.Columns != 1 || m2.Columns != 1 {
		panic("matrix sizes for dot product must have 1 column")
	}

	s := T(0)
	for i, v := range m.Data {
		s += v * m2.Data[i]
	}

	return s
}

func (m *Matrix[T]) MustAdd(m2 *Matrix[T]) *Matrix[T] {
	if m.Rows != m2.Rows || m.Columns != m2.Columns {
		panic("matrix sizes are incompatible for addition")
	}

	ma := New[T](m.Rows, m2.Columns)

	for i, v := range m.Data {
		ma.Data[i] = v + m2.Data[i]
	}

	return ma
}

func (m *Matrix[T]) MustMul(m2 *Matrix[T]) *Matrix[T] {
	if m.Columns != m2.Rows {
		panic("matrix sizes are incompatible for multiplication")
	}

	ma := New[T](m.Rows, m2.Columns)

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m2.Columns; j++ {
			s := T(0)
			for n := 0; n < m.Columns; n++ {
				s += m.At(i, n) * m2.At(n, j)
			}
			ma.Set(i, j, s)
		}
	}

	return ma
}
