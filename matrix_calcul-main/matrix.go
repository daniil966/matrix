package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	data [][]float64
	size int
}

func NewMatrix(size int) *Matrix {
	data := make([][]float64, size)
	for i := range data {
		data[i] = make([]float64, size)
	}
	return &Matrix{data: data, size: size}
}

func (m *Matrix) GetSize() int {
	return m.size
}

func (m *Matrix) GetData() [][]float64 {
	return m.data
}

func (m *Matrix) SetData(data [][]float64) error {
	if len(data) != m.size {
		return errors.New("неверный размер данных")
	}
	for _, row := range data {
		if len(row) != m.size {
			return errors.New("неверный размер строки")
		}
	}
	m.data = data
	return nil
}

func (m *Matrix) String() string {
	var s string
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			s += fmt.Sprintf("%8.2f ", m.data[i][j])
		}
		s += "\n"
	}
	return s
}
