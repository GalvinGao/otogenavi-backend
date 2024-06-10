// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/GalvinGao/otogenavi-backend/internal/ent/cabinet"
	"github.com/GalvinGao/otogenavi-backend/internal/ent/game"
	"github.com/GalvinGao/otogenavi-backend/internal/ent/location"
)

// Cabinet is the model entity for the Cabinet schema.
type Cabinet struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// LocationID holds the value of the "location_id" field.
	LocationID string `json:"location_id,omitempty"`
	// GameID holds the value of the "game_id" field.
	GameID string `json:"game_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CabinetQuery when eager-loading is set.
	Edges        CabinetEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CabinetEdges holds the relations/edges for other nodes in the graph.
type CabinetEdges struct {
	// Location holds the value of the location edge.
	Location *Location `json:"location,omitempty"`
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CabinetEdges) LocationOrErr() (*Location, error) {
	if e.Location != nil {
		return e.Location, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: location.Label}
	}
	return nil, &NotLoadedError{edge: "location"}
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CabinetEdges) GameOrErr() (*Game, error) {
	if e.Game != nil {
		return e.Game, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: game.Label}
	}
	return nil, &NotLoadedError{edge: "game"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cabinet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cabinet.FieldCount:
			values[i] = new(sql.NullInt64)
		case cabinet.FieldID, cabinet.FieldLocationID, cabinet.FieldGameID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cabinet fields.
func (c *Cabinet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cabinet.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case cabinet.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				c.Count = int(value.Int64)
			}
		case cabinet.FieldLocationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value.Valid {
				c.LocationID = value.String
			}
		case cabinet.FieldGameID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				c.GameID = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cabinet.
// This includes values selected through modifiers, order, etc.
func (c *Cabinet) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryLocation queries the "location" edge of the Cabinet entity.
func (c *Cabinet) QueryLocation() *LocationQuery {
	return NewCabinetClient(c.config).QueryLocation(c)
}

// QueryGame queries the "game" edge of the Cabinet entity.
func (c *Cabinet) QueryGame() *GameQuery {
	return NewCabinetClient(c.config).QueryGame(c)
}

// Update returns a builder for updating this Cabinet.
// Note that you need to call Cabinet.Unwrap() before calling this method if this Cabinet
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cabinet) Update() *CabinetUpdateOne {
	return NewCabinetClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Cabinet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cabinet) Unwrap() *Cabinet {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cabinet is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cabinet) String() string {
	var builder strings.Builder
	builder.WriteString("Cabinet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", c.Count))
	builder.WriteString(", ")
	builder.WriteString("location_id=")
	builder.WriteString(c.LocationID)
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(c.GameID)
	builder.WriteByte(')')
	return builder.String()
}

// Cabinets is a parsable slice of Cabinet.
type Cabinets []*Cabinet