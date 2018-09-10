package fatorveiculo

import (
	"math"
)

const (
	valorFatorCaminhaoCacamba = 1.05
)

type fatorCacamba struct {
}

func NewFatorCaminhaoCacamba() *fatorCacamba {
	return &fatorCacamba{}
}

//Calcular
func (*fatorCacamba) Calcular(valor float64) float64 {
	total := valorFatorCaminhaoCacamba * valor
	return math.Floor((total * 100) / 100)
}
