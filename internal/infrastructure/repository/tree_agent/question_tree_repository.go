package tree_agent

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type QuestionTreeRepository struct {
	db         *sqlx.DB
	subAgentID string
}

func NewQuestionTreeRepository(db *sqlx.DB) *QuestionTreeRepository {
	return &QuestionTreeRepository{db: db, subAgentID: os.Getenv("SECURITY_TREE_SUB_AGENT_ID")}
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
