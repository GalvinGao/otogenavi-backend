// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/GalvinGao/otogenavi-backend/internal/x/postgis"
)

type CoordinateWithin struct {
	Center         *postgis.PointS `json:"center,omitempty"`
	RadiusInMeters *float64        `json:"radiusInMeters,omitempty"`
}
