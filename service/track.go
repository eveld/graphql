package service

import (
	"database/sql"
	"log"

	"github.com/eveld/graphql/models"
	"github.com/jmoiron/sqlx"
)

// TrackService handles interaction with tracks.
type TrackService struct {
	db *sqlx.DB
}

// NewTrackService creates a new TrackService.
func NewTrackService(db *sqlx.DB) *TrackService {
	return &TrackService{
		db: db,
	}
}

// CreateTrack creates a new track.
func (s *TrackService) CreateTrack(input models.NewTrack) (*models.Track, error) {
	query, err := s.db.PrepareNamed(
		`INSERT INTO tracks(
			slug,
			title,
			teaser,
			description
		) 
		VALUES(
			:slug,
			:title,
			:teaser,
			:description
		)
		RETURNING *`)
	if err != nil {
		log.Fatal(err)
	}

	var track models.Track
	err = query.Get(&track, input)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Println("Failed to create track", err)
		return nil, err
	}

	return &track, nil
}

// GetTrack returns a specific track by ID.
func (s *TrackService) GetTrack(trackID string) (*models.Track, error) {
	var track models.Track

	params := map[string]interface{}{
		"id": trackID,
	}

	query, err := s.db.PrepareNamed(
		`SELECT * 
		FROM tracks
		WHERE id = :id
		AND deleted = 0`)
	if err != nil {
		log.Fatal(err)
	}

	err = query.Get(&track, params)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Println("Failed to retrieve track", err)
		return nil, err
	}

	return &track, nil
}

// FindTracks returns a list of tracks.
func (s *TrackService) FindTracks() ([]models.Track, error) {
	tracks := make([]models.Track, 0)

	query, err := s.db.Preparex(
		`SELECT * 
		FROM tracks
		WHERE deleted = 0`)
	if err != nil {
		log.Fatal(err)
	}

	err = query.Select(&tracks)
	if err == sql.ErrNoRows {
		return tracks, nil
	}

	if err != nil {
		log.Println("Failed to retrieve tracks", err)
		return nil, err
	}

	return tracks, nil
}
