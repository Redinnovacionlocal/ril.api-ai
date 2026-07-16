package tree_agent

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

type QuestionTreeRepository struct {
	db *sqlx.DB
}

func NewQuestionTreeRepository(db *sqlx.DB) *QuestionTreeRepository {
	return &QuestionTreeRepository{db: db}
}

func (r *QuestionTreeRepository) GetExcelGCSPath(agentPrefix string) (string, error) {
	subAgentId := os.Getenv(fmt.Sprintf("%s_TREE_SUB_AGENT_ID", strings.ToUpper(agentPrefix)))
	if subAgentId == "" {
		return "", fmt.Errorf("AGENT_TREE_SUB_AGENT_ID no configurado")
	}
	var path string
	err := r.db.Get(&path, `
        SELECT excel_gcs_path 
        FROM tree_sub_agent 
        WHERE id_tree_sub_agent = $1
    `, subAgentId)
	return path, err
}
