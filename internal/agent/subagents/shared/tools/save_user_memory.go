package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

type MemoryRecord struct {
	RecordType   string         `json:"record_type" jsonschema:"Tipo de registro: respuesta_AD (respuesta a una pregunta que devolvió lookup_tree_questions), nivel_madurez, odm_en_curso o contexto_municipio (datos del gobierno local y del usuario, incluidas las respuestas a los datos mínimos)."`
	AdQuestionId *string        `json:"ad_question_id,omitempty" jsonschema:"Solo para respuesta_AD. El campo id de la pregunta que devolvió lookup_tree_questions en esta conversación (por ejemplo \"43\"). Nunca un tag ni un texto armado por vos. Si la pregunta no salió de lookup_tree_questions, no uses respuesta_AD."`
	Tema         *string        `json:"tema,omitempty" jsonschema:"Solo para nivel_madurez y odm_en_curso. El id de la pregunta del árbol (el campo id que devuelve lookup_tree_questions) a la que corresponde el nivel o la OdM, por ejemplo \"23\"."`
	Payload      map[string]any `json:"payload" jsonschema:"Contenido del registro, con la forma que indican las instrucciones para cada record_type."`
}

type SaveUserMemoryToolInput struct {
	Records []MemoryRecord `json:"records" jsonschema:"Registros a guardar. Al menos uno."`
}

var (
	dbInstance *sqlx.DB
	dbOnce     sync.Once
	dbErr      error
)

func getDB() (*sqlx.DB, error) {
	dbOnce.Do(func() {
		dbInstance, dbErr = sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
		if dbErr == nil {
			dbInstance.SetMaxOpenConns(25)
			dbInstance.SetMaxIdleConns(25)
			dbInstance.SetConnMaxLifetime(5 * time.Minute)
		}
	})
	return dbInstance, dbErr
}

func NewSaveMemoryTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "save_user_memory",
		Description: "LLAMADA OBLIGATORIA PREVIA: Si el usuario aporta cualquier dato nuevo sobre su gobierno local (ej. responde a tus preguntas), DEBÉS llamar a esta herramienta ANTES de escribir una sola palabra en tu respuesta. Guarda en la memoria todo lo que el gobierno local aporta durante la conversación.",
	}, func(ctx tool.Context, input SaveUserMemoryToolInput) (map[string]any, error) {
		log.Printf("Saving %d memory records for user %s", len(input.Records), ctx.UserID())
		db, err := getDB()
		if err != nil {
			return nil, fmt.Errorf("failed to get db: %w", err)
		}

		teamId, _ := ctx.State().Get("team_id")
		if teamId == nil {
			teamId = "default_team"
		}

		now := time.Now().UTC()
		saved := 0
		upserted := 0
		var rejected []map[string]any

		for _, record := range input.Records {
			if record.RecordType == "context_municipio" {
				record.RecordType = "contexto_municipio"
			}
			if reason := invalidRecordReason(record); reason != "" {
				rejected = append(rejected, rejectedRecord(record, reason))
				continue
			}
			updated, err := saveRecord(ctx, db, teamId, now, record)
			if err != nil {
				log.Printf("save_user_memory: %v", err)
				rejected = append(rejected, rejectedRecord(record, err.Error()))
				continue
			}
			if updated {
				upserted++
			} else {
				saved++
			}
		}

		result := map[string]any{
			"status":   "success",
			"inserted": saved,
			"updated":  upserted,
		}
		if len(rejected) > 0 {
			result["status"] = "partial"
			result["rejected"] = rejected
		}
		return result, nil
	})
}

// invalidRecordReason returns why a record cannot be saved, or "" if it can.
// The reason is returned to the model so it can retry with the right data.
func invalidRecordReason(record MemoryRecord) string {
	if len(record.Payload) == 0 {
		return "payload is required and cannot be empty"
	}
	switch record.RecordType {
	case "respuesta_AD":
		if record.AdQuestionId == nil || !isTreeID(*record.AdQuestionId) {
			return "ad_question_id must be the numeric id returned by lookup_tree_questions. " +
				"If the question did not come from lookup_tree_questions, save it as contexto_municipio or nivel_madurez"
		}
	case "nivel_madurez", "odm_en_curso":
		if record.Tema == nil || !isTreeID(*record.Tema) {
			return "tema must be the numeric id of the tree question returned by lookup_tree_questions, never a tag"
		}
	}
	return ""
}

func isTreeID(id string) bool {
	if id == "" {
		return false
	}
	for _, r := range id {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func rejectedRecord(record MemoryRecord, reason string) map[string]any {
	r := map[string]any{"record_type": record.RecordType, "reason": reason}
	if record.AdQuestionId != nil {
		r["ad_question_id"] = *record.AdQuestionId
	}
	return r
}

// saveRecord updates the existing record when there is one, otherwise inserts
// it. Returns true when it updated.
func saveRecord(ctx tool.Context, db *sqlx.DB, teamId any, now time.Time, record MemoryRecord) (bool, error) {
	payloadJSON, err := json.Marshal(record.Payload)
	if err != nil {
		return false, fmt.Errorf("failed to serialize payload for record_type %s: %w", record.RecordType, err)
	}

	if (record.RecordType == "odm_en_curso" || record.RecordType == "nivel_madurez") && record.Tema != nil {
		result, err := db.ExecContext(ctx,
			`UPDATE public.user_security_memory
				SET payload = $1,
					updated_at = $2,
					source_agent = $6
			  WHERE user_id     = $3
				AND record_type = $4
				AND payload->>'tema' = $5`,
			payloadJSON,
			now,
			ctx.UserID(),
			record.RecordType,
			*record.Tema,
			ctx.AgentName(),
		)
		if err != nil {
			return false, fmt.Errorf("failed to update record_type %s tema %s: %w", record.RecordType, *record.Tema, err)
		}
		if rows, _ := result.RowsAffected(); rows > 0 {
			return true, nil
		}
	}

	if record.RecordType == "contexto_municipio" {
		result, err := db.ExecContext(ctx,
			`UPDATE public.user_security_memory
				SET payload = $1,
					updated_at = $2,
					source_agent = $6
			WHERE user_id     = $3
				AND record_type = $4
				AND payload->>'key' = $5`,
			payloadJSON,
			now,
			ctx.UserID(),
			record.RecordType,
			record.Payload["key"],
			ctx.AgentName(),
		)
		if err != nil {
			return false, fmt.Errorf("failed to update contexto_municipio key %v: %w", record.Payload["key"], err)
		}
		if rows, _ := result.RowsAffected(); rows > 0 {
			return true, nil
		}
	}

	if record.RecordType == "respuesta_AD" {
		result, err := db.ExecContext(ctx,
			`UPDATE public.user_security_memory
				SET payload = $1,
					updated_at = $2
			  WHERE user_id        = $3
				AND record_type    = 'respuesta_AD'
				AND source_agent   = $4
				AND ad_question_id = $5`,
			payloadJSON,
			now,
			ctx.UserID(),
			ctx.AgentName(),
			*record.AdQuestionId,
		)
		if err != nil {
			return false, fmt.Errorf("failed to update respuesta_AD question %s: %w", *record.AdQuestionId, err)
		}
		if rows, _ := result.RowsAffected(); rows > 0 {
			return true, nil
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
		return false, fmt.Errorf("failed to insert record_type %s: %w", record.RecordType, err)
	}
	return false, nil
}
