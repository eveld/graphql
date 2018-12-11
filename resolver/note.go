package resolver

import (
	"context"

	"github.com/eveld/graphql/models"
)

func (r *queryResolver) Note(ctx context.Context, noteID string) (*models.Note, error) {
	return nil, nil
}

func (r *queryResolver) Notes(ctx context.Context, trackID string) ([]models.Note, error) {
	return []models.Note{}, nil
}
