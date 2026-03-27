package tools

import (
	"database/sql"
	"os"

	"github.com/jmoiron/sqlx"
	"google.golang.org/adk/tool"
)

type GetUserMemoryToolArgs struct {
	Username string `json:"username,omitempty" jsonSchema:"The username of the user whose memory you want to retrieve. If not provided, it defaults to the current user."`
}

func GetUserMemoryToolFunc(ctx tool.Context, args GetUserMemoryToolArgs) (map[string]any, error) {
	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, record_type, ad_question_id, odm_id, payload, created_at FROM public.user_security_memory WHERE user_id = $1 ORDER BY created_at DESC", ctx.UserID())
	defer rows.Close()
	// return rows
	var memories []map[string]any
	for rows.Next() {
		var id, recordType string
		var adQuestionId, odmId sql.NullString
		var payload []byte
		var createdAt string
		if err := rows.Scan(&id, &recordType, &adQuestionId, &odmId, &payload, &createdAt); err != nil {
			return nil, err
		}
		memories = append(memories, map[string]any{
			"id":             id,
			"record_type":    recordType,
			"ad_question_id": adQuestionId.String,
			"odm_id":         odmId.String,
			"payload":        payload,
			"created_at":     createdAt,
		})
	}
	return map[string]any{"memories": memories}, nil
}
