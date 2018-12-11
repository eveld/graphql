package resolver

import (
	"context"

	"github.com/eveld/graphql/models"
)

func (r *queryResolver) Challenge(ctx context.Context, challengeID string) (*models.Challenge, error) {
	return nil, nil
}

func (r *queryResolver) Challenges(ctx context.Context, trackID string) ([]models.Challenge, error) {
	return r.ChallengeService.FindChallenges(trackID)
}

func (r *mutationResolver) CreateChallenge(ctx context.Context, input models.NewChallenge) (*models.Challenge, error) {
	return r.ChallengeService.CreateChallenge(input)
}
