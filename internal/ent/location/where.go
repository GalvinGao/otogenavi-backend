// Code generated by ent, DO NOT EDIT.

package location

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/GalvinGao/otogenavi-backend/internal/ent/predicate"
	"github.com/GalvinGao/otogenavi-backend/internal/x/postgis"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldName, v))
}

// DeduplicationKey applies equality check predicate on the "deduplication_key" field. It's identical to DeduplicationKeyEQ.
func DeduplicationKey(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldDeduplicationKey, v))
}

// RawAddress applies equality check predicate on the "raw_address" field. It's identical to RawAddressEQ.
func RawAddress(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldRawAddress, v))
}

// Coordinate applies equality check predicate on the "coordinate" field. It's identical to CoordinateEQ.
func Coordinate(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCoordinate, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldName, v))
}

// DeduplicationKeyEQ applies the EQ predicate on the "deduplication_key" field.
func DeduplicationKeyEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldDeduplicationKey, v))
}

// DeduplicationKeyNEQ applies the NEQ predicate on the "deduplication_key" field.
func DeduplicationKeyNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldDeduplicationKey, v))
}

// DeduplicationKeyIn applies the In predicate on the "deduplication_key" field.
func DeduplicationKeyIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldDeduplicationKey, vs...))
}

// DeduplicationKeyNotIn applies the NotIn predicate on the "deduplication_key" field.
func DeduplicationKeyNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldDeduplicationKey, vs...))
}

// DeduplicationKeyGT applies the GT predicate on the "deduplication_key" field.
func DeduplicationKeyGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldDeduplicationKey, v))
}

// DeduplicationKeyGTE applies the GTE predicate on the "deduplication_key" field.
func DeduplicationKeyGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldDeduplicationKey, v))
}

// DeduplicationKeyLT applies the LT predicate on the "deduplication_key" field.
func DeduplicationKeyLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldDeduplicationKey, v))
}

// DeduplicationKeyLTE applies the LTE predicate on the "deduplication_key" field.
func DeduplicationKeyLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldDeduplicationKey, v))
}

// DeduplicationKeyContains applies the Contains predicate on the "deduplication_key" field.
func DeduplicationKeyContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldDeduplicationKey, v))
}

// DeduplicationKeyHasPrefix applies the HasPrefix predicate on the "deduplication_key" field.
func DeduplicationKeyHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldDeduplicationKey, v))
}

// DeduplicationKeyHasSuffix applies the HasSuffix predicate on the "deduplication_key" field.
func DeduplicationKeyHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldDeduplicationKey, v))
}

// DeduplicationKeyEqualFold applies the EqualFold predicate on the "deduplication_key" field.
func DeduplicationKeyEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldDeduplicationKey, v))
}

// DeduplicationKeyContainsFold applies the ContainsFold predicate on the "deduplication_key" field.
func DeduplicationKeyContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldDeduplicationKey, v))
}

// RawAddressEQ applies the EQ predicate on the "raw_address" field.
func RawAddressEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldRawAddress, v))
}

// RawAddressNEQ applies the NEQ predicate on the "raw_address" field.
func RawAddressNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldRawAddress, v))
}

// RawAddressIn applies the In predicate on the "raw_address" field.
func RawAddressIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldRawAddress, vs...))
}

// RawAddressNotIn applies the NotIn predicate on the "raw_address" field.
func RawAddressNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldRawAddress, vs...))
}

// RawAddressGT applies the GT predicate on the "raw_address" field.
func RawAddressGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldRawAddress, v))
}

// RawAddressGTE applies the GTE predicate on the "raw_address" field.
func RawAddressGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldRawAddress, v))
}

// RawAddressLT applies the LT predicate on the "raw_address" field.
func RawAddressLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldRawAddress, v))
}

// RawAddressLTE applies the LTE predicate on the "raw_address" field.
func RawAddressLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldRawAddress, v))
}

// RawAddressContains applies the Contains predicate on the "raw_address" field.
func RawAddressContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldRawAddress, v))
}

// RawAddressHasPrefix applies the HasPrefix predicate on the "raw_address" field.
func RawAddressHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldRawAddress, v))
}

// RawAddressHasSuffix applies the HasSuffix predicate on the "raw_address" field.
func RawAddressHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldRawAddress, v))
}

// RawAddressIsNil applies the IsNil predicate on the "raw_address" field.
func RawAddressIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldRawAddress))
}

// RawAddressNotNil applies the NotNil predicate on the "raw_address" field.
func RawAddressNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldRawAddress))
}

// RawAddressEqualFold applies the EqualFold predicate on the "raw_address" field.
func RawAddressEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldRawAddress, v))
}

// RawAddressContainsFold applies the ContainsFold predicate on the "raw_address" field.
func RawAddressContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldRawAddress, v))
}

// CoordinateEQ applies the EQ predicate on the "coordinate" field.
func CoordinateEQ(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCoordinate, v))
}

// CoordinateNEQ applies the NEQ predicate on the "coordinate" field.
func CoordinateNEQ(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldCoordinate, v))
}

// CoordinateIn applies the In predicate on the "coordinate" field.
func CoordinateIn(vs ...*postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldCoordinate, vs...))
}

// CoordinateNotIn applies the NotIn predicate on the "coordinate" field.
func CoordinateNotIn(vs ...*postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldCoordinate, vs...))
}

// CoordinateGT applies the GT predicate on the "coordinate" field.
func CoordinateGT(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldCoordinate, v))
}

// CoordinateGTE applies the GTE predicate on the "coordinate" field.
func CoordinateGTE(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldCoordinate, v))
}

// CoordinateLT applies the LT predicate on the "coordinate" field.
func CoordinateLT(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldCoordinate, v))
}

// CoordinateLTE applies the LTE predicate on the "coordinate" field.
func CoordinateLTE(v *postgis.PointS) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldCoordinate, v))
}

// CoordinateIsNil applies the IsNil predicate on the "coordinate" field.
func CoordinateIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldCoordinate))
}

// CoordinateNotNil applies the NotNil predicate on the "coordinate" field.
func CoordinateNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldCoordinate))
}

// HasCabinets applies the HasEdge predicate on the "cabinets" edge.
func HasCabinets() predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CabinetsTable, CabinetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCabinetsWith applies the HasEdge predicate on the "cabinets" edge with a given conditions (other predicates).
func HasCabinetsWith(preds ...predicate.Cabinet) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		step := newCabinetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Location) predicate.Location {
	return predicate.Location(sql.NotPredicates(p))
}
