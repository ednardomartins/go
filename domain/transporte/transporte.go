package transporte

import (
	"math"

	"github.com/ednardomartins/gerenciador-financeiro-transporte/domain"
)

type service struct {
}

const (
	fatorDistanciaPavimentada    = 0.54
	fatorDistanciaNaoPavimentada = 0.62
	totalDeCargaSemCustoExtra    = 5
	fatorCargaExtraPorTonelada   = 0.02
)

func NewService() *service {
	return &service{}
}

func (*service) CalcularCustoTransporte(detalhe *domain.Transporte) (float64, error) {
	fatorVeiculo, err := CriarFatorVeiculo(detalhe.CodigoVeiculo)
	if err != nil {
		return 0.0, err
	}
	total := calcularValorDistancias(detalhe)
	total = fatorVeiculo.Calcular(total)
	return calcularExcessoDeCarga(detalhe, total), nil
}

func (*service) ValidarDados(detalhe *domain.Transporte) bool {
	if detalhe.DistanciaNaoPavimentada == 0 && detalhe.DistanciaPavimentada == 0 {
		return false
	}
	if detalhe.CargaTransportada == 0 {
		return false
	}
	return true
}

func calcularValorDistancias(detalhe *domain.Transporte) float64 {
	valorPavimentada := calcularValorDistanciaPavimenta(detalhe.DistanciaPavimentada)
	valorNaoPavimentada := calcularValorDistanciaNaoPavimenta(detalhe.DistanciaNaoPavimentada)
	return valorPavimentada + valorNaoPavimentada
}

func calcularValorDistanciaPavimenta(distancia int) float64 {
	total := fatorDistanciaPavimentada * float64(distancia)
	return math.Floor((total * 100) / 100)
}

func calcularValorDistanciaNaoPavimenta(distancia int) float64 {
	total := fatorDistanciaNaoPavimentada * float64(distancia)
	return math.Floor((total * 100) / 100)
}

func calcularExcessoDeCarga(detalhe *domain.Transporte, valorAtual float64) float64 {
	totalCarga := detalhe.CargaTransportada
	if totalCarga <= totalDeCargaSemCustoExtra {
		return valorAtual
	}
	distanciaTotal := getDistanciaTotalPercorrida(detalhe)
	valorPorDistancia := getFatorExcessoDeCargaPorDistanciaTotal(totalCarga, distanciaTotal)
	valorTotal := valorAtual * valorPorDistancia
	return math.Floor((valorTotal * 100) / 100)
}

func getFatorExcessoDeCargaPorDistanciaTotal(totalCarga int, distanciaTotal int) float64 {
	excesso := totalCarga - totalDeCargaSemCustoExtra
	fator := fatorCargaExtraPorTonelada * float64(excesso)
	return float64(distanciaTotal) * fator
}

func getDistanciaTotalPercorrida(detalhe *domain.Transporte) int {
	return detalhe.DistanciaPavimentada + detalhe.DistanciaNaoPavimentada
}

func (s *service) GetVeiculos() string {
	jsonVeiculos := "[{'codigo':1,'nome':'Caminhão baú','fator':1.00},{'codigo':2,'nome':'Caminhão baú','fator':1.00},{'codigo':1,'nome':'Carreta','fator':1.12}]"
	return jsonVeiculos
}
