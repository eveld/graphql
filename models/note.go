package models

// Note describes a note track item.
type Note struct {
	ID      string `yaml:"id"         json:"id"         db:"id"`
	Content string `yaml:"content"    json:"content"    db:"content"`
	Deleted int64  `yaml:"-"          json:"-"          db:"deleted"`
}

// IsTrackItem allows us to query a union of track items.
func (Note) IsTrackItem() {}
