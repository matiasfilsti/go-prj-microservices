package controllers

import (
	"aioperator/src/api/controllers/ctrserrors"
	"aioperator/src/api/domain/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ask(c *gin.Context) {
	fmt.Println("Ask endpoint hit")

	var req models.QuestionRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pregunta no válida"})
		return
	}

	docs, err := s.Core.Cache.Search(c.Request.Context(), req.Question)
	if err != nil {
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error al buscar en el cache")
		return
	}
	fmt.Println("cache response", docs)
	llmResponse, err := s.Core.Llm.GenerateResponse(c.Request.Context(), docs, req.Question)
	if err != nil {
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error al generar respuesta")
		return
	}
	fmt.Println("llm response", llmResponse)
	c.String(http.StatusOK, llmResponse)
}
