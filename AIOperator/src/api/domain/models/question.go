package models

type QuestionRequest struct {
	Question string `form:"question" json:"question" binding:"required"`
}
