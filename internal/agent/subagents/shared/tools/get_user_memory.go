package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

type GetUserMemoryToolArgs struct{}

func NewGetMemoryTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "get_user_memory",
		Description: "Recupera la memoria acumulada del usuario sobre su municipio. Devuelve datos concretos aportados por el usuario, oportunidades de mejora identificadas y contexto relevante.",
	}, func(ctx tool.Context, _ GetUserMemoryToolArgs) (map[string]any, error) {
		db, err := getDB()
		if err != nil {
			return nil, fmt.Errorf("failed to get db: %w", err)
		}

		rows, err := db.QueryContext(ctx,
			`SELECT id, record_type, ad_question_id, payload, source_agent, created_at, updated_at
			   FROM public.user_security_memory
			  WHERE user_id = $1
			  ORDER BY updated_at DESC`,
			ctx.UserID(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to query memory: %w", err)
		}
		defer rows.Close()

		currentAgent := ctx.AgentName()

		currentAgentData := map[string][]map[string]any{
			"respuesta_AD":       {},
			"nivel_madurez":      {},
			"odm_en_curso":       {},
			"contexto_municipio": {},
		}

		otherAgentData := make(map[string]map[string][]map[string]any)

		for rows.Next() {
			var (
				id           string
				recordType   string
				adQuestionId sql.NullString
				payloadRaw   []byte
				sourceAgent  string
				createdAt    string
				updatedAt    string
			)

			if err := rows.Scan(&id, &recordType, &adQuestionId, &payloadRaw, &sourceAgent, &createdAt, &updatedAt); err != nil {
				return nil, fmt.Errorf("failed to scan row: %w", err)
			}

			var payload map[string]any
			if err := json.Unmarshal(payloadRaw, &payload); err != nil {
				return nil, fmt.Errorf("failed to parse payload for record %s: %w", id, err)
			}

			record := map[string]any{
				"id":         id,
				"payload":    payload,
				"created_at": createdAt,
				"updated_at": updatedAt,
			}
			if adQuestionId.Valid {
				record["ad_question_id"] = adQuestionId.String
			}

			if sourceAgent == currentAgent {
				currentAgentData[recordType] = append(currentAgentData[recordType], record)
			} else {
				if _, ok := otherAgentData[sourceAgent]; !ok {
					otherAgentData[sourceAgent] = map[string][]map[string]any{
						"respuesta_AD":       {},
						"nivel_madurez":      {},
						"odm_en_curso":       {},
						"contexto_municipio": {},
					}
				}
				otherAgentData[sourceAgent][recordType] = append(otherAgentData[sourceAgent][recordType], record)
			}
		}

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error iterating rows: %w", err)
		}

		return map[string]any{
			"tu_memoria":            currentAgentData,
			"memoria_otros_agentes": otherAgentData,
		}, nil
	})
}
