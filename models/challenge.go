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
	// Tabs       []Tab  `yaml:"tabs"             json:"-"               db:"-"`
	Deleted int64 `yaml:"-"                json:"-"                  db:"deleted"          checksum:"-"`
}

func (c *Challenge) Tabs() ([]Tab, error) {
	return []Tab{
		Tab{
			ID:    "x",
			Title: "title",
		},
	}, nil
}

// Tab describes a tab on a challenge.
type Tab struct {
	ID       string  `yaml:"-"                  json:"id"                 db:"id"               checksum:"-"`
	Title    string  `yaml:"title"              json:"title"              db:"title"`
	Type     TabType `yaml:"type"               json:"type"               db:"type"`
	Hostname string  `yaml:"hostname,omitempty" json:"hostname,omitempty" db:"hostname"`
	Path     string  `yaml:"path,omitempty"     json:"path,omitempty"     db:"path"`
	Port     int     `yaml:"port,omitempty"     json:"port,omitempty"     db:"port"`
	URL      string  `yaml:"url,omitempty"      json:"url,omitempty"      db:"url"`
	Target   string  `yaml:"-"                  json:"target"             db:"-"                checksum:"-"`
	Index    int     `yaml:"-"                  json:"index"              db:"index"            checksum:"-"`
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
