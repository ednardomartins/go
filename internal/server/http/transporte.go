package http

import (
	"net/http"

	"github.com/ednardomartins/gerenciador-financeiro-transporte/domain"
	"github.com/gin-gonic/gin"
)

// Calcular Custo Transporte
func (h *handler) PostCalcularTransporte(c *gin.Context) {
	transporte := &domain.Transporte{}
	if err := c.BindJSON(&transporte); err != nil {
		return
	}

	if !h.transporteService.ValidarDados(transporte) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	total, err := h.transporteService.CalcularCustoTransporte(transporte)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, total)
}

// Calcular Custo Transporte
func (h *handler) GetVeiculos(c *gin.Context) {

	user := h.transporteService.GetVeiculos()
	//if err != nil {
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}
	c.JSON(http.StatusOK, user)
}
