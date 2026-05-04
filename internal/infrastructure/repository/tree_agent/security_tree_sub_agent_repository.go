package tree_agent

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type SecurityTreeSubAgentRepository struct {
	db         *sqlx.DB
	subAgentID string
}

func NewSecurityTreeSubAgentRepository(db *sqlx.DB) *SecurityTreeSubAgentRepository {
	return &SecurityTreeSubAgentRepository{db: db, subAgentID: os.Getenv("SECURITY_TREE_SUB_AGENT_ID")}
}

func (r *SecurityTreeSubAgentRepository) GetDimensions() ([]string, error) {
	var dimensions []string
	err := r.db.Select(&dimensions, `
        SELECT DISTINCT tq.dimension 
        FROM tree_questions tq
        INNER JOIN tree_sub_agent tsa ON tq.id_tree_sub_agent = tsa.id_tree_sub_agent
        WHERE tsa.id_tree_sub_agent = $1
        AND tq.dimension IS NOT NULL
    `, r.subAgentID)
	return dimensions, err
}

func (r *SecurityTreeSubAgentRepository) GetTags() ([]string, error) {
	var tags []string
	err := r.db.Select(&tags, `
        SELECT DISTINCT unnest(tq.tags) AS tag
        FROM tree_questions tq
        INNER JOIN tree_sub_agent tsa ON tq.id_tree_sub_agent = tsa.id_tree_sub_agent
        WHERE tsa.id_tree_sub_agent = $1
    `, r.subAgentID)
	return tags, err
}
