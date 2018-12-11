package service

import (
	"database/sql"
	"log"

	"github.com/eveld/graphql/models"
	"github.com/jmoiron/sqlx"
)

// ChallengeService handles interaction with challenges.
type ChallengeService struct {
	db *sqlx.DB
}

// NewChallengeService creates a new ChallengeService.
func NewChallengeService(db *sqlx.DB) *ChallengeService {
	return &ChallengeService{
		db: db,
	}
}

// CreateChallenge creates a new challenge.
func (s *ChallengeService) CreateChallenge(input models.NewChallenge) (*models.Challenge, error) {
	query, err := s.db.PrepareNamed(
		`INSERT INTO challenges(
			slug,
			title,
			teaser,
			assignment,
			difficulty,
			timelimit
		) 
		VALUES(
			:slug,
			:title,
			:teaser,
			:assignment,
			:difficulty,
			:timelimit
		)
		RETURNING *`)
	if err != nil {
		log.Fatal(err)
	}

	var challenge models.Challenge
	err = query.Get(&challenge, input)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Println("Failed to create challenge", err)
		return nil, err
	}

	return &challenge, nil
}

// GetChallenge returns a specific challenge by ID.
func (s *ChallengeService) GetChallenge(challengeID string) (*models.Challenge, error) {
	var challenge models.Challenge

	params := map[string]interface{}{
		"id": challengeID,
	}

	query, err := s.db.PrepareNamed(
		`SELECT * 
		FROM challenges
		WHERE id = :id
		AND deleted = 0`)
	if err != nil {
		log.Fatal(err)
	}

	err = query.Get(&challenge, params)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Println("Failed to retrieve challenge", err)
		return nil, err
	}

	return &challenge, nil
}

// FindChallenges returns a list of challenges.
func (s *ChallengeService) FindChallenges(trackID string) ([]models.Challenge, error) {
	challenges := make([]models.Challenge, 0)

	query, err := s.db.Preparex(
		`SELECT * 
		FROM challenges
		WHERE deleted = 0`)
	if err != nil {
		log.Fatal(err)
	}

	err = query.Select(&challenges)
	if err == sql.ErrNoRows {
		return challenges, nil
	}

	if err != nil {
		log.Println("Failed to retrieve challenges", err)
		return nil, err
	}

	return challenges, nil
}
