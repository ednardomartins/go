package fatorveiculo

import (
	"math"
)

const (
	valorFatorCaminhaoBau = 1.05
)

type fatorBau struct {
}

func NewFatorCaminhaoBau() *fatorBau {
	return &fatorBau{}
}

//Calcular
func (*fatorBau) Calcular(valor float64) float64 {
	total := valorFatorCaminhaoBau * valor
	return math.Floor((total * 100) / 100)
}
