package cluster

import "gonum.org/v1/gonum/mat"

// This file wraps types from gonum in order to prevent vendoring issues such
// as type mismatch:
//    *"gonum.org/v1/gonum/mat".Dense
//    vs
//    *"github.com/e-XpertSolutions/go-cluster/vendor/gonum.org/v1/gonum/mat".Dense

// DenseVector wraps *gonum.org/v1/gonum/mat.VecDense type.
type DenseVector struct {
	*mat.VecDense
}

func NewDenseVector(n int, data []float64) *DenseVector {
	return &DenseVector{VecDense: mat.NewVecDense(n, data)}
}

type DenseMatrix struct {
	*mat.Dense
}

func NewDenseMatrix(r, c int, data []float64) *DenseMatrix {
	return &DenseMatrix{Dense: mat.NewDense(r, c, data)}
}
