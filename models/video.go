package models

// Video describes a video track item.
type Video struct {
	ID      string `json:"id"         db:"id"`
	URL     string `json:"url"        db:"url"`
	Deleted int64  `json:"-"          db:"deleted"`
}

// IsTrackItem allows us to query a union of track items.
func (Video) IsTrackItem() {}
