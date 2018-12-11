package resolver

import (
	"context"

	"github.com/eveld/graphql/models"
)

func (r *queryResolver) Video(ctx context.Context, videoID string) (*models.Video, error) {
	return nil, nil
}

func (r *queryResolver) Videos(ctx context.Context, trackID string) ([]models.Video, error) {
	return []models.Video{}, nil
}
