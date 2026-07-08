package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/models"
)

type ArticleService struct {
	db    *sql.DB
	graph neo4j.DriverWithContext
}

type ArticleInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func NewArticleService(db *sql.DB, graph neo4j.DriverWithContext) *ArticleService {
	return &ArticleService{db: db, graph: graph}
}

func (s *ArticleService) Create(ctx context.Context, authorID string, in ArticleInput) (models.Article, error) {
	if err := validateArticleInput(in); err != nil {
		return models.Article{}, err
	}
	if in.Status == "" {
		in.Status = "draft"
	}

	var article models.Article
	err := s.db.QueryRowContext(ctx, `
		INSERT INTO articles (author_id, title, content, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, author_id, title, content, status, synced_to_graph, created_at, updated_at
	`, authorID, strings.TrimSpace(in.Title), strings.TrimSpace(in.Content), in.Status).Scan(
		&article.ID, &article.AuthorID, &article.Title, &article.Content, &article.Status, &article.SyncedToGraph, &article.CreatedAt, &article.UpdatedAt,
	)
	if err != nil {
		return models.Article{}, err
	}

	if article.Status == "published" {
		if err := s.syncToGraph(ctx, article); err == nil {
			article.SyncedToGraph = true
			_, _ = s.db.ExecContext(ctx, `UPDATE articles SET synced_to_graph = TRUE WHERE id = $1`, article.ID)
		}
	}

	return article, nil
}

func (s *ArticleService) GetByID(ctx context.Context, id string) (models.Article, error) {
	var article models.Article
	err := s.db.QueryRowContext(ctx, `
		SELECT id, author_id, title, content, status, synced_to_graph, created_at, updated_at
		FROM articles WHERE id = $1
	`, id).Scan(
		&article.ID, &article.AuthorID, &article.Title, &article.Content, &article.Status, &article.SyncedToGraph, &article.CreatedAt, &article.UpdatedAt,
	)
	return article, err
}

func (s *ArticleService) Update(ctx context.Context, id string, in ArticleInput) (models.Article, error) {
	if err := validateArticleInput(in); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	err := s.db.QueryRowContext(ctx, `
		UPDATE articles
		SET title = $2, content = $3, status = $4
		WHERE id = $1
		RETURNING id, author_id, title, content, status, synced_to_graph, created_at, updated_at
	`, id, strings.TrimSpace(in.Title), strings.TrimSpace(in.Content), in.Status).Scan(
		&article.ID, &article.AuthorID, &article.Title, &article.Content, &article.Status, &article.SyncedToGraph, &article.CreatedAt, &article.UpdatedAt,
	)
	if err != nil {
		return models.Article{}, err
	}

	if article.Status == "published" {
		if err := s.syncToGraph(ctx, article); err == nil {
			article.SyncedToGraph = true
			_, _ = s.db.ExecContext(ctx, `UPDATE articles SET synced_to_graph = TRUE WHERE id = $1`, article.ID)
		}
	}

	return article, nil
}

func (s *ArticleService) syncToGraph(ctx context.Context, article models.Article) error {
	session := s.graph.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, `
		MERGE (u:User {id: $authorId})
		MERGE (a:Article {id: $id})
		SET a.title = $title, a.content = $content, a.status = $status
		MERGE (u)-[:AUTHORED]->(a)
	`, map[string]any{
		"id":       article.ID,
		"authorId": article.AuthorID,
		"title":    article.Title,
		"content":  article.Content,
		"status":   article.Status,
	})
	return err
}

func validateArticleInput(in ArticleInput) error {
	if strings.TrimSpace(in.Title) == "" || strings.TrimSpace(in.Content) == "" {
		return errors.New("title and content are required")
	}
	switch in.Status {
	case "", "draft", "published", "retracted":
		return nil
	default:
		return errors.New("invalid article status")
	}
}
