package tree_agent

import "github.com/jmoiron/sqlx"

type TreeQuestionRepository struct {
	db *sqlx.DB
}

func NewTreeQuestionRepository(db *sqlx.DB) *TreeQuestionRepository {
	return &TreeQuestionRepository{db: db}
}

func (r *TreeQuestionRepository) GetDimensions() ([]string, error) {
	var dimensions []string
	err := r.db.Select(&dimensions, `
        SELECT DISTINCT dimension FROM tree_questions WHERE dimension IS NOT NULL
    `)
	return dimensions, err
}

func (r *TreeQuestionRepository) GetTags() ([]string, error) {
	var tags []string
	err := r.db.Select(&tags, `
        SELECT DISTINCT unnest(tags) AS tag FROM tree_questions
    `)
	return tags, err
}
