package transporte

import (
	"errors"

	"github.com/ednardomartins/gerenciador-financeiro-transporte/domain/fatorveiculo"
)

const (
	caminhaoBau     = 1
	caminhaoCacamba = 2
	carreta         = 3
)

//Erroa quando veiculo nao existe
var ErrVeiculoNaoExiste = errors.New("Veiculo não cadastrado.")

//Criar um fator do veiculo
func CriarFatorVeiculo(codigo int) (fatorveiculo.FatorVeiculo, error) {
	if caminhaoBau == codigo {
		return fatorveiculo.NewFatorCaminhaoBau(), nil
	}
	if caminhaoCacamba == codigo {
		return fatorveiculo.NewFatorCaminhaoCacamba(), nil
	}
	if carreta == codigo {
		return fatorveiculo.NewFatorCarreta(), nil
	}
	return nil, ErrVeiculoNaoExiste
}
