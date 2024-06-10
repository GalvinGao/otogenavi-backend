// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/GalvinGao/otogenavi-backend/internal/x/postgis"
)

// CreateLocationInput represents a mutation input for creating locations.
type CreateLocationInput struct {
	Name       string
	RawAddress *string
	Coordinate *postgis.PointS
	CabinetIDs []string
}

// Mutate applies the CreateLocationInput on the LocationMutation builder.
func (i *CreateLocationInput) Mutate(m *LocationMutation) {
	m.SetName(i.Name)
	if v := i.RawAddress; v != nil {
		m.SetRawAddress(*v)
	}
	if v := i.Coordinate; v != nil {
		m.SetCoordinate(v)
	}
	if v := i.CabinetIDs; len(v) > 0 {
		m.AddCabinetIDs(v...)
	}
}

// SetInput applies the change-set in the CreateLocationInput on the LocationCreate builder.
func (c *LocationCreate) SetInput(i CreateLocationInput) *LocationCreate {
	i.Mutate(c.Mutation())
	return c
}