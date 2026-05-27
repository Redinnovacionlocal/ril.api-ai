package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
	"ril.api-ia/internal/application/usecase"
	"ril.api-ia/internal/domain/entity"
)

type RunHandler struct {
	ctx            context.Context
	runners        map[string]*runner.Runner
	sessionUseCase usecase.SessionUseCase
	appName        string
}

func NewRunHandler(ctx context.Context, runners map[string]*runner.Runner, sessionUseCase usecase.SessionUseCase) *RunHandler {
	return &RunHandler{
		ctx:            ctx,
		runners:        runners,
		sessionUseCase: sessionUseCase,
		appName:        os.Getenv("APP_NAME"),
	}
}

type RunSSERequest struct {
	Parts     []*genai.Part `json:"parts" binding:"required"`
	SessionId *string       `json:"session_id"`
}

func (rh *RunHandler) RunSSE(c *gin.Context) {
	var runSseRequest RunSSERequest
	user := c.MustGet("user").(*entity.User)
	err := c.ShouldBindJSON(&runSseRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s, err := rh.GetSession(runSseRequest.SessionId, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("SessionId", s.ID())
	c.Writer.Flush()

	agentKey := c.GetHeader("x-agent")
	r, ok := rh.runners[agentKey]
	if !ok {
		r = rh.runners["orchestrator"]
	}

	resp := r.Run(rh.ctx,
		s.UserID(),
		s.ID(),
		genai.NewContentFromParts(runSseRequest.Parts, genai.RoleUser),
		agent.RunConfig{StreamingMode: agent.StreamingModeSSE})
	for event, err := range resp {
		if err != nil {
			fmt.Fprintf(c.Writer, "data: %s\n\n", err.Error())
			c.Writer.Flush()
			return
		}
		jsonEvent, _ := json.Marshal(entity.FromSessionEvent(*event))
		fmt.Fprintf(c.Writer, "data: %s\n\n", jsonEvent)
		c.Writer.Flush()
	}
}

func (rh *RunHandler) GetSession(sessionId *string, user *entity.User) (session.Session, error) {
	if sessionId == nil {
		err, s := rh.sessionUseCase.StoreSession(user, rh.appName)
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	err, s := rh.sessionUseCase.GetSession(user, rh.appName, *sessionId)
	if err != nil {
		return nil, err
	}
	return s, nil
}
