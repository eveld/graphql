package models

import (
	"fmt"
	"io"
	"strconv"
)

// Tab describes a tab on a challenge.
type Tab struct {
	ID       string  `yaml:"-"                  json:"id"                 db:"id"`
	Title    string  `yaml:"title"              json:"title"              db:"title"`
	Type     TabType `yaml:"type"               json:"type"               db:"type"`
	Hostname string  `yaml:"hostname,omitempty" json:"hostname,omitempty" db:"hostname"`
	Path     string  `yaml:"path,omitempty"     json:"path,omitempty"     db:"path"`
	Port     int     `yaml:"port,omitempty"     json:"port,omitempty"     db:"port"`
	URL      string  `yaml:"url,omitempty"      json:"url,omitempty"      db:"url"`
	Target   string  `yaml:"-"                  json:"target"             db:"-"`
	Index    int     `yaml:"-"                  json:"index"              db:"index"`
}

// TabType describes all possible tab types.
type TabType string

const (
	// TabTypeTerminal is a terminal tab
	TabTypeTerminal TabType = "TERMINAL"
	// TabTypeEditor is an editor tab.
	TabTypeEditor TabType = "EDITOR"
	// TabTypeService is a service tab.
	TabTypeService TabType = "SERVICE"
	// TabTypeExternal is an external tab.
	TabTypeExternal TabType = "EXTERNAL"
)

// IsValid checks if the passed tab type is valid.
func (e TabType) IsValid() bool {
	switch e {
	case TabTypeTerminal, TabTypeEditor, TabTypeService, TabTypeExternal:
		return true
	}
	return false
}

// String converts the tab type to string.
func (e TabType) String() string {
	return string(e)
}

// UnmarshalGQL unmarhals the graphql enum to a tab type.
func (e *TabType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TabType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TabType", str)
	}
	return nil
}

// MarshalGQL marhals the tab type to a graphql enum.
func (e TabType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
