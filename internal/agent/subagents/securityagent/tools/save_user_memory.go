package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/adk/tool"
)

type SaveUserMemoryToolInput struct {
	RecordType   string         `json:"record_type"    jsonSchema:"Enabled record types: 'respuesta_AD', 'odm_detectada', 'contexto_municipio'"`
	AdQuestionId *string        `json:"ad_question_id,omitempty" jsonSchema:"Required if record_type is 'respuesta_AD'. The ID of the AD question associated with the response."`
	OdmId        *string        `json:"odm_id,omitempty"         jsonSchema:"Required if record_type is 'odm_detectada'. The ID of the detected ODM."`
	Payload      map[string]any `json:"payload"                 jsonSchema:"Object structure depends on record_type. For 'respuesta_AD': {value, raw_text, alert_triggered (bool), alert_detail?}. For 'odm_detectada': {description, dimension, origin_question_id?, suggested_actions (string[])}. For 'contexto_municipio': {key (one of: poblacion|tamanio_ciudad|provincia_pais|presupuesto_seguridad|restriccion_presupuestaria|prioridad_politica|nombre_responsable_area), value}."`
}

// SaveUserMemoryToolFunc saves user-related information (AD responses, detected ODMs, or municipal context) into the database for future reference and analysis.
// The specific structure of the payload depends on the type of record being saved.

func SaveUserMemoryToolFunc(ctx tool.Context, input SaveUserMemoryToolInput) (map[string]any, error) {
	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	payloadJSON, err := json.Marshal(input.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize payload: %w", err)
	}

	teamId, _ := ctx.State().Get("team_id")
	if teamId == nil {
		teamId = "default_team" // o maneja el error según tu lógica de negocio
	}

	_, err = db.ExecContext(ctx,
		`INSERT INTO public.user_security_memory
       (id, user_id, quality_status, team_id, session_id, record_type, ad_question_id, odm_id, payload, source_agent)
    VALUES
       ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		uuid.New(),
		ctx.UserID(),
		"validated",
		teamId,
		ctx.SessionID(),
		input.RecordType,
		input.AdQuestionId,
		input.OdmId,
		payloadJSON,
		ctx.AgentName(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save memory: %w", err)
	}

	return map[string]any{"status": "success"}, nil
}
