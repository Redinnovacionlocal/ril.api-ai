package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

type MemoryRecord struct {
	RecordType   string         `json:"record_type" jsonSchema:"enum:respuesta_AD,nivel_madurez,odm_en_curso,contexto_municipio"`
	AdQuestionId *string        `json:"ad_question_id,omitempty"`
	Tema         *string        `json:"tema,omitempty"`
	Payload      map[string]any `json:"payload"`
}

type SaveUserMemoryToolInput struct {
	Records []MemoryRecord `json:"records" jsonSchema:"minItems:1"`
}

func NewSaveMemoryTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "save_user_memory",
		Description: "LLAMADA OBLIGATORIA PREVIA: Si el usuario aporta cualquier dato nuevo sobre su municipio (ej. responde a tus preguntas), DEBÉS llamar a esta herramienta ANTES de escribir una sola palabra en tu respuesta. Guarda en la memoria todo lo que el municipio aporta durante la conversación.",
	}, func(ctx tool.Context, input SaveUserMemoryToolInput) (map[string]any, error) {
		log.Printf("Saving %d memory records for user %s", len(input.Records), ctx.UserID())
		db, err := sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
		if err != nil {
			return nil, fmt.Errorf("failed to open db: %w", err)
		}
		defer db.Close()

		teamId, _ := ctx.State().Get("team_id")
		if teamId == nil {
			teamId = "default_team"
		}

		now := time.Now().UTC()
		saved := 0
		upserted := 0

		for _, record := range input.Records {
			if record.RecordType == "context_municipio" {
				record.RecordType = "contexto_municipio"
			}
			payloadJSON, err := json.Marshal(record.Payload)
			if err != nil {
				return nil, fmt.Errorf("failed to serialize payload for record_type %s: %w", record.RecordType, err)
			}

			if (record.RecordType == "odm_en_curso" || record.RecordType == "nivel_madurez") && record.Tema != nil {
				result, err := db.ExecContext(ctx,
					`UPDATE public.user_security_memory
						SET payload = $1,
							updated_at = $2
					  WHERE user_id     = $3
						AND record_type = $4
						AND payload->>'tema' = $5`,
					payloadJSON,
					now,
					ctx.UserID(),
					record.RecordType,
					*record.Tema,
				)
				if err != nil {
					return nil, fmt.Errorf("failed to update record_type %s tema %s: %w", record.RecordType, *record.Tema, err)
				}

				rows, _ := result.RowsAffected()
				if rows > 0 {
					upserted++
					continue
				}
			}

			if record.RecordType == "contexto_municipio" {
				result, err := db.ExecContext(ctx,
					`UPDATE public.user_security_memory
						SET payload = $1,
							updated_at = $2
					WHERE user_id     = $3
						AND record_type = $4
						AND payload->>'key' = $5`,
					payloadJSON,
					now,
					ctx.UserID(),
					record.RecordType,
					record.Payload["key"],
				)
				if err == nil {
					rows, _ := result.RowsAffected()
					if rows > 0 {
						upserted++
						continue
					}
				}
			}

			_, err = db.ExecContext(ctx,
				`INSERT INTO public.user_security_memory
				   (id, user_id, quality_status, team_id, session_id, record_type, ad_question_id, payload, source_agent, created_at, updated_at)
				 VALUES
				   ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $10)`,
				uuid.New(),
				ctx.UserID(),
				"validated",
				teamId,
				ctx.SessionID(),
				record.RecordType,
				record.AdQuestionId,
				payloadJSON,
				ctx.AgentName(),
				now,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to insert record_type %s: %w", record.RecordType, err)
			}
			saved++
		}

		return map[string]any{
			"status":   "success",
			"inserted": saved,
			"updated":  upserted,
		}, nil
	})
}
