package tree_agent

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type QuestionTreeRepository struct {
	db         *sqlx.DB
	subAgentID string
}

func NewQuestionTreeRepository(db *sqlx.DB) *QuestionTreeRepository {
	id := os.Getenv("SECURITY_TREE_SUB_AGENT_ID")
	if _, err := uuid.Parse(id); err != nil {
		id = ""
	}
	return &QuestionTreeRepository{db: db, subAgentID: id}
}

func (r *QuestionTreeRepository) GetExcelGCSPath() (string, error) {
	if r.subAgentID == "" {
		return "", fmt.Errorf("SECURITY_TREE_SUB_AGENT_ID no configurado")
	}
	var path string
	err := r.db.Get(&path, `
        SELECT excel_gcs_path 
        FROM tree_sub_agent 
        WHERE id_tree_sub_agent = $1
    `, r.subAgentID)
	return path, err
}
