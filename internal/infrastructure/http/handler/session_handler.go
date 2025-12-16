package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"ril.api-ia/internal/application/usecase"
	"ril.api-ia/internal/domain/entity"
)

type SessionHandler struct {
	sessionUseCase *usecase.SessionUseCase
}

func NewSessionHandler(sessionUseCase *usecase.SessionUseCase) *SessionHandler {
	return &SessionHandler{
		sessionUseCase: sessionUseCase,
	}
}

func (sessionHandler *SessionHandler) CreateSession(c *gin.Context) {
	appName := os.Getenv("APP_NAME")
	user := c.MustGet("user").(*entity.User)
	err, session := sessionHandler.sessionUseCase.StoreSession(user, appName)
	if err != nil {
		log.Printf("Error while storing session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	s, _ := entity.FromSession(session)
	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
	return
}

func (sessionHandler *SessionHandler) ListSessions(c *gin.Context) {
	appName := os.Getenv("APP_NAME")
	user := c.MustGet("user").(*entity.User)
	sessionData := make([]entity.Session, 0)
	err, sessions := sessionHandler.sessionUseCase.GetAllSessions(user, appName)
	if err != nil {
		log.Printf("Error while getting all sessions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}
	for _, s := range sessions {
		session, _ := entity.FromSession(s)
		sessionData = append(sessionData, session)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": sessionData,
	})
}

func (sessionHandler *SessionHandler) GetSession(c *gin.Context) {
	appName := os.Getenv("APP_NAME")
	user := c.MustGet("user").(*entity.User)
	sessionId := c.Param("session_id")
	err, session := sessionHandler.sessionUseCase.GetSession(user, appName, sessionId)
	if err != nil {
		log.Printf("Error while getting session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Session not found",
		})
		return
	}
	s, _ := entity.FromSession(session)
	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
}

func (sessionHandler *SessionHandler) DeleteSession(c *gin.Context) {
	appName := os.Getenv("APP_NAME")
	user := c.MustGet("user").(*entity.User)
	sessionId := c.Param("session_id")
	err := sessionHandler.sessionUseCase.RemoveSession(user, appName, sessionId)
	if err != nil {
		log.Printf("Error while removing session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Removed session successfully",
	})
	return
}
