package middleware

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"google.golang.org/genai"
	"ril.api-ia/internal/domain/entity"
)

type GroundingState struct {
	Metadata *genai.GroundingMetadata
}

type contextKey string

const GroundingStateKey contextKey = "grounding_state"

func GroundingMetadataLogger(dbAgent *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("start grounding middleware")
		c.Next()

		sessionIDRaw, exists := c.Get("session_id")
		if !exists {
			return
		}

		sessionID, ok := sessionIDRaw.(string)
		if !ok {
			log.Println(">> ERROR: session_id is not a string")
			return
		}
		
		log.Print("end grounding middleware")
		defer entity.GroundingCache.Delete(sessionID)
		
		if val, ok := entity.GroundingCache.Load(sessionID); ok {
			metadata, ok := val.(*genai.GroundingMetadata)
			if !ok {
				log.Println(">> ERROR: cached grounding metadata is of invalid type")
				return
			}
			metadataJSON, _ := json.Marshal(metadata)

			query := `
				UPDATE "events" 
				SET grounding_metadata = $1 
				WHERE session_id = $2 
				AND id = (
					SELECT id FROM "events" 
					WHERE session_id = $2 
					ORDER BY timestamp DESC LIMIT 1
				)
			`
			_, err := dbAgent.Exec(query, string(metadataJSON), sessionID)
			if err != nil {
				log.Printf("Error actualizando grounding_metadata: %v\n", err)
			} else {
				log.Printf("✅ Grounding metadata guardada por SessionID.")
			}
		}
	}
}
