package models

import (
	"fmt"
	"io"
	"strconv"
)

// Role describes all possible roles.
type Role string

const (
	// RoleAdmin is an admin.
	RoleAdmin Role = "ADMIN"
	// RoleUser is a normal user.
	RoleUser Role = "USER"
)

// IsValid checks if the passed role is valid.
func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

// String converts the role to string.
func (e Role) String() string {
	return string(e)
}

// UnmarshalGQL unmarhals the graphql enum to a role.
func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

// MarshalGQL marhals the role to a graphql enum.
func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
