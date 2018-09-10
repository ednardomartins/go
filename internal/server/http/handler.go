package http

import (
	"net/http"

	"github.com/ednardomartins/gerenciador-financeiro-transporte/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	transporteService domain.TransporteService
}

//Criar um handler do trasnporte service....
func NewHandler(transporteService domain.TransporteService) http.Handler {
	handler := &handler{
		transporteService: transporteService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(handler.recovery())
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	v1 := router.Group("/v1")

	v1.GET("/calcularTransporte", handler.PostCalcularTransporte)
	v1.GET("/veiculos", handler.GetVeiculos)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
