package models

// Challenge describes a challenge track item.
type Challenge struct {
	ID         string `yaml:"id"               json:"id"                 db:"id"`
	Slug       string `yaml:"slug"             json:"slug"               db:"slug"`
	Title      string `yaml:"title"            json:"title"              db:"title"`
	Teaser     string `yaml:"teaser"           json:"teaser"             db:"teaser"`
	Assignment string `yaml:"assignment"       json:"assignment"         db:"assignment"`
	Difficulty string `yaml:"difficulty"       json:"difficulty"         db:"difficulty"`
	TimeLimit  int    `yaml:"timelimit"        json:"timelimit"          db:"timelimit"`
	Tabs       []Tab  `yaml:"tabs"             json:"tabs"               db:"-"`
	Deleted    int64  `yaml:"-"                json:"-"                  db:"deleted"`
}

// IsTrackItem allows us to query a union of track items.
func (Challenge) IsTrackItem() {}

// NewChallenge holds all fields needed to create a new challenge.
type NewChallenge struct {
	Slug       string `yaml:"slug"             json:"slug"               db:"slug"`
	Title      string `yaml:"title"            json:"title"              db:"title"`
	Teaser     string `yaml:"teaser"           json:"teaser"             db:"teaser"`
	Assignment string `yaml:"assignment"       json:"assignment"         db:"assignment"`
	Difficulty string `yaml:"difficulty"       json:"difficulty"         db:"difficulty"`
	TimeLimit  int    `yaml:"timelimit"        json:"timelimit"          db:"timelimit"`
}
