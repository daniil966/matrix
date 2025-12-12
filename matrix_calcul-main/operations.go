package main

import "errors"

func AddMatrices(a, b [][]float64) ([][]float64, error) {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil, errors.New("матрицы должны быть одинакового размера")
	}

	size := len(a)
	result := make([][]float64, size)
	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result, nil
}

func MultiplyMatrixByScalar(matrix [][]float64, scalar float64) [][]float64 {
	size := len(matrix)
	result := make([][]float64, size)
	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			result[i][j] = matrix[i][j] * scalar
		}
	}
	return result
}

func MultiplyMatrices(a, b [][]float64) ([][]float64, error) {
	size := len(a)
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil, errors.New("матрицы должны быть одинакового размера")
	}

	result := make([][]float64, size)
	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			for k := 0; k < size; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result, nil
}
