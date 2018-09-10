package fatorveiculo

import (
	"math"
)

const (
	valorFatorCarreta = 1.05
)

type fatorCarreta struct {
}

func NewFatorCarreta() *fatorCarreta {
	return &fatorCarreta{}
}

//Calcular
func (*fatorCarreta) Calcular(valor float64) float64 {
	total := valorFatorCaminhaoBau * valor
	return math.Floor((total * 100) / 100)
}
