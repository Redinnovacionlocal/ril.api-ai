package tree_agent

import (
	"context"
	"encoding/json"
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

const (
	redisKeyQuestions  = "tree:questions"
	redisKeyDimensions = "tree:dimensions"
	redisKeyTags       = "tree:tags"
)

type TreeCacheManager struct {
	mu         sync.Mutex
	gcsClient  *storage.Client
	bucketName string
	repo       *QuestionTreeRepository
	rdb        *redis.Client
	ttl        time.Duration
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

func (m *TreeCacheManager) GetData(ctx context.Context) ([]entity.QuestionTree, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	val, err := m.rdb.Get(ctx, redisKeyQuestions).Result()

	if err == nil {
		var preguntas []entity.QuestionTree
		if err := json.Unmarshal([]byte(val), &preguntas); err == nil {
			return preguntas, nil
		}
	}
	log.Printf("Cache miss para preguntas, refrescando cache: %v", err)
	preguntas, _, _, err := m.refreshCache(ctx)
	return preguntas, err
}

func (m *TreeCacheManager) GetDimensions(ctx context.Context) ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	val, err := m.rdb.Get(ctx, redisKeyDimensions).Result()
	if err == nil {
		var dims []string
		if json.Unmarshal([]byte(val), &dims) == nil {
			return dims, nil
		}
	}

	_, dims, _, err := m.refreshCache(ctx)
	return dims, err
}

func (m *TreeCacheManager) GetTags(ctx context.Context) ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	val, err := m.rdb.Get(ctx, redisKeyTags).Result()
	if err == nil {
		var tags []string
		if json.Unmarshal([]byte(val), &tags) == nil {
			return tags, nil
		}
	}

	_, _, tags, err := m.refreshCache(ctx)
	return tags, err
}

func (m *TreeCacheManager) refreshCache(ctx context.Context) ([]entity.QuestionTree, []string, []string, error) {
	objectName, err := m.repo.GetExcelGCSPath()
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

	data, _ := json.Marshal(preguntas)
	m.rdb.Set(ctx, redisKeyQuestions, data, m.ttl)

	dims, _ := json.Marshal(dimensions)
	m.rdb.Set(ctx, redisKeyDimensions, dims, m.ttl)

	t, _ := json.Marshal(tags)
	m.rdb.Set(ctx, redisKeyTags, t, m.ttl)

	fmt.Printf("Caché actualizada: %d preguntas, %d dimensiones, %d tags\n", len(preguntas), len(dimensions), len(tags))
	return preguntas, dimensions, tags, nil
}

func (m *TreeCacheManager) Lookup(ctx context.Context, id, dimension, tag, query string) ([]entity.QuestionTree, error) {
	data, err := m.GetData(ctx)
	if err != nil {
		return nil, err
	}

	var results []entity.QuestionTree
	q := strings.ToLower(query)

	for _, p := range data {
		if id != "" && p.ID != id {
			continue
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
		if q != "" && !strings.Contains(strings.ToLower(p.Pregunta), q) && !strings.Contains(strings.ToLower(p.Bajo.Descripcion), q) {
			continue
		}
		results = append(results, p)
	}
	return results, nil
}
