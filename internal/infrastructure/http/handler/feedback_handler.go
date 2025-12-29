package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ril.api-ia/internal/application/usecase"
	"ril.api-ia/internal/domain/entity"
)

type FeedbackHandler struct {
	ctx                  context.Context
	eventFeedbackUseCase usecase.EventFeedbackUseCase
}

type FeedbackRequest struct {
	IsPositive bool    `json:"is_positive" binding:"required"`
	Comments   *string `json:"comments"`
	ErrorType  *string `json:"error_type"`
}

func NewFeedbackHandler(ctx context.Context, eventFeedbackUseCase usecase.EventFeedbackUseCase) *FeedbackHandler {
	return &FeedbackHandler{
		ctx:                  ctx,
		eventFeedbackUseCase: eventFeedbackUseCase,
	}
}

func (handler *FeedbackHandler) SaveFeedback(c *gin.Context) {
	invocationId := c.Param("invocation_id")
	user := c.MustGet("user").(*entity.User)
	var feedbackRequest FeedbackRequest
	if err := c.ShouldBindJSON(&feedbackRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := handler.eventFeedbackUseCase.SaveFeedback(invocationId, user, feedbackRequest.IsPositive, feedbackRequest.Comments, feedbackRequest.ErrorType)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Feedback saved successfully"})
}
