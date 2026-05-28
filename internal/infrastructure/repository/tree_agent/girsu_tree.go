package tree_agent

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed girsu_tree.json
var girsuTreeData []byte

type GirsuQuestion struct {
	ID                  string   `json:"id"`
	Dimension           string   `json:"dimension"`
	Question            string   `json:"question"`
	Options             string   `json:"options"`
	SupportingDocument  string   `json:"supporting_document"`
	MinimumCertCriteria string   `json:"minimum_cert_criteria"`
	ValidFormats        string   `json:"valid_formats"`
	GoodPractices       string   `json:"good_practices"`
	AlertSignals        string   `json:"alert_signals"`
	AgentHelp           string   `json:"agent_help"`
	Tags                []string `json:"-"`
	RawTags             string   `json:"tags"`
	ODMs                string   `json:"odms"`
}

type GirsuTreeRepository struct {
	questions []GirsuQuestion
}

func NewGirsuTreeRepository() (*GirsuTreeRepository, error) {
	var questions []GirsuQuestion
	if err := json.Unmarshal(girsuTreeData, &questions); err != nil {
		return nil, err
	}
	for i := range questions {
		questions[i].Tags = strings.Split(questions[i].RawTags, "\n")
	}
	return &GirsuTreeRepository{questions: questions}, nil
}

func (r *GirsuTreeRepository) Lookup(id, dimension, tag, query string) ([]GirsuQuestion, error) {
	var results []GirsuQuestion
	q := strings.ToLower(query)

	for _, p := range r.questions {
		if id != "" {
			ids := strings.Split(id, ",")
			matched := false
			for _, i := range ids {
				if p.ID == strings.TrimSpace(i) {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}
		if dimension != "" && !strings.Contains(strings.ToLower(p.Dimension), strings.ToLower(dimension)) {
			continue
		}
		if tag != "" {
			found := false
			for _, t := range p.Tags {
				if strings.EqualFold(t, tag) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		if q != "" {
			matched := false
			for _, t := range p.Tags {
				if strings.EqualFold(t, q) {
					matched = true
					break
				}
			}
			if !matched && strings.Contains(strings.ToLower(p.Dimension), strings.ToLower(q)) {
				matched = true
			}
			if !matched {
				continue
			}
		}
		results = append(results, p)
	}
	return results, nil
}
