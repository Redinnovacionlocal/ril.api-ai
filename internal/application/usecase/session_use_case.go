package usecase

import (
	"context"
	"log"
	"strconv"

	"google.golang.org/adk/session"
	"google.golang.org/genai"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/domain/repository"
)

type SessionUseCase struct {
	ctx            context.Context
	SessionService session.Service
	UserRepository repository.UserRepository
}

func NewSessionUseCase(ctx context.Context, sessionService session.Service, userRepository repository.UserRepository) *SessionUseCase {
	return &SessionUseCase{
		ctx:            ctx,
		SessionService: sessionService,
		UserRepository: userRepository,
	}
}

func (s *SessionUseCase) StoreSession(user *entity.User, appName string) (error, session.Session) {
	userProfile, err := s.UserRepository.GetUserProfile(user)
	if err != nil {
		return err, nil
	}
	request := &session.CreateRequest{
		AppName: appName,
		UserID:  strconv.FormatInt(user.Id, 10),
		State: map[string]any{
			"user:first_name": user.FirstName,
			"user:last_name":  user.LastName,
			"user:country":    userProfile.Country,
			"user:charge":     userProfile.Charge,
			"user:sector":     userProfile.Sector,
			"user:area":       userProfile.Sector,
			"user:job_title":  userProfile.JobTitle,
		},
	}
	response, err := s.SessionService.Create(s.ctx, request)
	if err != nil {
		return err, nil
	}
	return nil, response.Session
}

func (s *SessionUseCase) GetSession(user *entity.User, appName string, sessionId string) (error, session.Session) {
	request := &session.GetRequest{
		AppName:   appName,
		UserID:    strconv.FormatInt(user.Id, 10),
		SessionID: sessionId,
	}
	response, err := s.SessionService.Get(s.ctx, request)
	if err != nil {
		return err, nil
	}
	return nil, response.Session
}

func (s *SessionUseCase) RemoveSession(user *entity.User, appName string, sessionId string) error {
	request := &session.DeleteRequest{
		AppName:   appName,
		UserID:    strconv.FormatInt(user.Id, 10),
		SessionID: sessionId,
	}
	err := s.SessionService.Delete(s.ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionUseCase) GetAllSessions(user *entity.User, appName string) (error, []session.Session) {
	request := &session.ListRequest{
		AppName: appName,
		UserID:  strconv.FormatInt(user.Id, 10),
	}
	response, err := s.SessionService.List(s.ctx, request)
	if err != nil {
		return err, nil
	}
	return nil, response.Sessions
}

func (s *SessionUseCase) GenerateSessionTitle(ctx context.Context, userPrompt string, modelResponse string) (string, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendVertexAI,
	})
	if err != nil {
		log.Fatal()
	}

	temperature := float32(0.5)
	m := "gemini-2.5-flash-lite"
	prompt := "Genera un título conciso y descriptivo (máximo 5 palabras) para una sesión de chat basada en el siguiente mensaje del usuario y la respuesta del agente.\n\nMensaje del usuario: \"" + userPrompt + "\"\n\nRespuesta del agente: \"" + modelResponse + "\"\n\nTítulo:"

	result, err := client.Models.GenerateContent(ctx, m, genai.Text(prompt), &genai.GenerateContentConfig{
		Temperature:     &temperature,
		MaxOutputTokens: 100,
	})
	if err != nil {
		log.Fatal("Error generating session title", err)
	}
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		text := result.Candidates[0].Content.Parts[0].Text
		return text, nil
	}
	return "", nil
}
