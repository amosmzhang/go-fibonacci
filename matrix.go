package fibonacci

import (
	"gonum.org/v1/gonum/mat"
)

func Matrix(n int) int {
	matrix := mat.NewDense(2, 2, []float64{1, 1, 1, 0})
	matrix.Pow(matrix, n)
	return int(matrix.At(0, 1))
}
