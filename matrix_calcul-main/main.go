package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Калькулятор матриц Тощев ИС-323 💾 ===")
	fmt.Println()

	for {
		fmt.Println("Выберите операцию:")
		fmt.Println("1. Сложение матриц")
		fmt.Println("2. Умножение матрицы на число")
		fmt.Println("3. Умножение двух матриц")
		fmt.Println("4. Выход")

		var choice string
		fmt.Print("Ваш выбор: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			handleMatrixAddition()
		case "2":
			handleScalarMultiplication()
		case "3":
			handleMatrixMultiplication()
		case "4":
			fmt.Println("До свидания!")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
		fmt.Println()
	}
}

func handleMatrixAddition() {
	size := getMatrixSize()
	a := inputMatrix(size, "Первая матрица")
	b := inputMatrix(size, "Вторая матрица")

	result, err := AddMatrices(a, b)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}

	fmt.Println("\nРезультат сложения:")
	printMatrix(result)
}

func handleScalarMultiplication() {
	size := getMatrixSize()
	matrix := inputMatrix(size, "Матрица")
	var scalar float64
	fmt.Print("Введите число для умножения: ")
	fmt.Scanln(&scalar)

	result := MultiplyMatrixByScalar(matrix, scalar)
	fmt.Println("\nРезультат умножения на число:")
	printMatrix(result)
}

func handleMatrixMultiplication() {
	fmt.Println("Умножение матриц: A * B")
	fmt.Print("Размер первой матрицы (2 или 3): ")
	var sizeA int
	fmt.Scanln(&sizeA)

	fmt.Print("Размер второй матрицы (2 или 3): ")
	var sizeB int
	fmt.Scanln(&sizeB)

	if sizeA != sizeB {
		fmt.Println("Ошибка: размеры матриц должны совпадать для умножения 2x2 или 3x3.")
		return
	}

	a := inputMatrix(sizeA, "Первая матрица")
	b := inputMatrix(sizeB, "Вторая матрица")

	result, err := MultiplyMatrices(a, b)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}

	fmt.Println("\nРезультат умножения:")
	printMatrix(result)
}

func getMatrixSize() int {
	var size int
	for {
		fmt.Print("Введите размер матрицы (2 или 3): ")
		fmt.Scanln(&size)
		if size == 2 || size == 3 {
			break
		}
		fmt.Println("Неверный размер. Введите 2 или 3.")
	}
	return size
}

func inputMatrix(size int, name string) [][]float64 {
	fmt.Printf("\n%s (%dx%d):\n", name, size, size)
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			fmt.Printf("Элемент [%d][%d]: ", i+1, j+1)
			var val float64
			fmt.Scanln(&val)
			matrix[i][j] = val
		}
	}
	return matrix
}

func printMatrix(matrix [][]float64) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%8.2f ", matrix[i][j])
		}
		fmt.Println()
	}
}
