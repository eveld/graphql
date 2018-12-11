package models

// Track describes a track a user can start.
type Track struct {
	ID          string      `yaml:"id"             json:"id"               db:"id"`
	Slug        string      `yaml:"slug"           json:"slug"             db:"slug"`
	Title       string      `yaml:"title"          json:"title"            db:"title"`
	Teaser      string      `yaml:"teaser"         json:"teaser"           db:"teaser"`
	Description string      `yaml:"description"    json:"description"      db:"description"`
	Items       []TrackItem `yaml:"items"          json:"items"            db:"items"`
	Deleted     int64       `yaml:"-"              json:"-"                db:"deleted"`
}

// TrackItem describes an item that can be a Challenge, Video or Note.
type TrackItem interface {
	IsTrackItem()
}

// NewTrack holds all fields needed to create a new track.
type NewTrack struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Teaser      string `json:"teaser"`
	Description string `json:"description"`
}
