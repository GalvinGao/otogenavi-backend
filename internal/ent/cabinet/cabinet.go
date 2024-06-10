// Code generated by ent, DO NOT EDIT.

package cabinet

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the cabinet type in the database.
	Label = "cabinet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "cabinet_id"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldGameID holds the string denoting the game_id field in the database.
	FieldGameID = "game_id"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// LocationFieldID holds the string denoting the ID field of the Location.
	LocationFieldID = "location_id"
	// GameFieldID holds the string denoting the ID field of the Game.
	GameFieldID = "game_id"
	// Table holds the table name of the cabinet in the database.
	Table = "cabinets"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "cabinets"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location_id"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "cabinets"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_id"
)

// Columns holds all SQL columns for cabinet fields.
var Columns = []string{
	FieldID,
	FieldCount,
	FieldLocationID,
	FieldGameID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Cabinet queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByGameID orders the results by the game_id field.
func ByGameID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameID, opts...).ToFunc()
}

// ByLocationField orders the results by location field.
func ByLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationStep(), sql.OrderByField(field, opts...))
	}
}

// ByGameField orders the results by game field.
func ByGameField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGameStep(), sql.OrderByField(field, opts...))
	}
}
func newLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationInverseTable, LocationFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, LocationTable, LocationColumn),
	)
}
func newGameStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GameInverseTable, GameFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, GameTable, GameColumn),
	)
}
