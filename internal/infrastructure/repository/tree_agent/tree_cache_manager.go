package tree_agent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/redis/go-redis/v9"
	"ril.api-ia/internal/domain/entity"
)

type TreeCacheManager struct {
	mu         sync.Mutex
	gcsClient  *storage.Client
	bucketName string
	repo       *QuestionTreeRepository
	rdb        *redis.Client
	ttl        time.Duration
}

var ErrTreeNotConfigured = errors.New("Agent tree not configured (configuration pending)")

func (m *TreeCacheManager) isConfigured() bool {
	return m.repo != nil
}

func NewTreeCacheManager(client *storage.Client, bucket string, repo *QuestionTreeRepository, rdb *redis.Client, ttl time.Duration) *TreeCacheManager {
	return &TreeCacheManager{
		gcsClient:  client,
		bucketName: bucket,
		repo:       repo,
		rdb:        rdb,
		ttl:        ttl,
	}
}

func (m *TreeCacheManager) GetData(ctx context.Context, agentPrefix string) ([]entity.QuestionTree, error) {
	if !m.isConfigured() {
		return nil, ErrTreeNotConfigured
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	keyQuestions := fmt.Sprintf("%s:tree:questions", agentPrefix)
	val, err := m.rdb.Get(ctx, keyQuestions).Result()

	if err == nil {
		var preguntas []entity.QuestionTree
		if err := json.Unmarshal([]byte(val), &preguntas); err == nil {
			return preguntas, nil
		}
	}
	log.Printf("Cache miss para preguntas, refrescando cache: %v", err)
	preguntas, _, _, err := m.refreshCache(ctx, agentPrefix)
	return preguntas, err
}

func (m *TreeCacheManager) GetDimensions(ctx context.Context, agentPrefix string) ([]string, error) {
	if !m.isConfigured() {
		return []string{}, nil
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	keyDimensions := fmt.Sprintf("%s:tree:dimensions", agentPrefix)
	val, err := m.rdb.Get(ctx, keyDimensions).Result()
	if err == nil {
		var dims []string
		if json.Unmarshal([]byte(val), &dims) == nil {
			return dims, nil
		}
	}

	_, dims, _, err := m.refreshCache(ctx, agentPrefix)
	return dims, err
}

func (m *TreeCacheManager) GetTags(ctx context.Context, agentPrefix string) ([]string, error) {
	if !m.isConfigured() {
		return []string{}, nil
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	keyTags := fmt.Sprintf("%s:tree:tags", agentPrefix)
	val, err := m.rdb.Get(ctx, keyTags).Result()
	if err == nil {
		var tags []string
		if json.Unmarshal([]byte(val), &tags) == nil {
			return tags, nil
		}
	}

	_, _, tags, err := m.refreshCache(ctx, agentPrefix)
	return tags, err
}

func (m *TreeCacheManager) refreshCache(ctx context.Context, agentPrefix string) ([]entity.QuestionTree, []string, []string, error) {
	objectName, err := m.repo.GetExcelGCSPath(agentPrefix)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error obteniendo excel_gcs_path: %w", err)
	}

	rc, err := m.gcsClient.Bucket(m.bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error leyendo de GCS: %w", err)
	}
	defer rc.Close()

	fileBytes, err := io.ReadAll(rc)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error leyendo bytes: %w", err)
	}

	preguntas, dimensions, tags, err := parseExcelContent(fileBytes)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error parseando excel: %w", err)
	}

	keyQuestions := fmt.Sprintf("%s:tree:questions", agentPrefix)
	keyDimensions := fmt.Sprintf("%s:tree:dimensions", agentPrefix)
	keyTags := fmt.Sprintf("%s:tree:tags", agentPrefix)

	data, _ := json.Marshal(preguntas)
	m.rdb.Set(ctx, keyQuestions, data, m.ttl)

	dims, _ := json.Marshal(dimensions)
	m.rdb.Set(ctx, keyDimensions, dims, m.ttl)

	t, _ := json.Marshal(tags)
	m.rdb.Set(ctx, keyTags, t, m.ttl)

	fmt.Printf("Caché actualizada: %d preguntas, %d dimensiones, %d tags\n", len(preguntas), len(dimensions), len(tags))
	return preguntas, dimensions, tags, nil
}

func (m *TreeCacheManager) Lookup(ctx context.Context, id, dimension, tag, query string, agentPrefix string) ([]entity.QuestionTree, error) {
	data, err := m.GetData(ctx, agentPrefix)
	if err != nil {
		return nil, err
	}

	var results []entity.QuestionTree
	q := strings.ToLower(query)

	for _, p := range data {
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
			for _, t := range p.TagsRAG {
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
			for _, t := range p.TagsRAG {
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
