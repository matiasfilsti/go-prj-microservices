package controllers

import (
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
	c.String(http.StatusOK, "Recibí tu pregunta: "+req.Question)
}
