package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ril.api-ia/internal/application/usecase"
)

type SpeechToTextHandler struct {
	ctx               context.Context
	transcribeUseCase *usecase.TranscribeUseCase
}

type GenerateTranscriptionRequest struct {
	AudioBase64 string `json:"audio_base64" binding:"required"`
}

func NewSpeechToTextHandler(ctx context.Context, transcribeUseCase *usecase.TranscribeUseCase) *SpeechToTextHandler {

	return &SpeechToTextHandler{
		ctx:               ctx,
		transcribeUseCase: transcribeUseCase,
	}
}

func (s *SpeechToTextHandler) GenerateTranscription(c *gin.Context) {
	var reqBody GenerateTranscriptionRequest
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	base64Audio := reqBody.AudioBase64
	if base64Audio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "audio_base64 query parameter is required"})
		return
	}
	resp, err := s.transcribeUseCase.SpeechToText(base64Audio)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "failed to recognize speech"})
		return
	}
	if len(resp.Results) == 0 || len(resp.Results[0].Alternatives) == 0 {
		c.JSON(500, gin.H{"error": "no transcription results"})
		return
	}
	c.JSON(200, gin.H{"data": resp.Results[0].Alternatives[0].Transcript})
	return
}
