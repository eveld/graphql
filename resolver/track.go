package resolver

import (
	"context"

	"github.com/eveld/graphql/models"
)

func (r *queryResolver) Track(ctx context.Context, trackID string) (*models.Track, error) {
	return r.TrackService.GetTrack(trackID)
}

func (r *queryResolver) Tracks(ctx context.Context) ([]models.Track, error) {
	return r.TrackService.FindTracks()
}

func (r *mutationResolver) CreateTrack(ctx context.Context, input models.NewTrack) (*models.Track, error) {
	return r.TrackService.CreateTrack(input)
}
