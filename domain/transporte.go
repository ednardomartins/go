package domain

import (
	"fmt"
)

type TransporteService interface {
	ValidarDados(transporte *Transporte) bool
	CalcularCustoTransporte(transporte *Transporte) (float64, error)
	GetVeiculos() string
}

type Transporte struct {
	DistanciaPavimentada    int `json:"distanciaPavimentada,omitempty" bson:"distanciaPavimentada"`
	DistanciaNaoPavimentada int `json:"distanciaNaoPavimentada,omitempty" bson:"distanciaNaoPavimentada"`
	CodigoVeiculo           int `json:"codigoVeiculo,omitempty" bson:"codigoVeiculo"`
	CargaTransportada       int `json:"cargaTransportada,omitempty" bson:"cargaTransportada,omitempty"`
}

func NewTransporte(distanciaPavimentada int, distanciaNaoPavimentada int, codigoVeiculo int, cargaTransportada int) *Transporte {
	return &Transporte{DistanciaPavimentada: distanciaPavimentada, DistanciaNaoPavimentada: distanciaNaoPavimentada,
		CodigoVeiculo: codigoVeiculo, CargaTransportada: cargaTransportada}
}

func (u *Transporte) String() string {
	return fmt.Sprintf("cargaTransportada: %s - codigoVeiculo: %d", u.CargaTransportada, u.CodigoVeiculo)
}
