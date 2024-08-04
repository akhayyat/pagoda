// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// OryID applies equality check predicate on the "ory_id" field. It's identical to OryIDEQ.
func OryID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOryID, v))
}

// Admin applies equality check predicate on the "admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAdmin, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// OryIDEQ applies the EQ predicate on the "ory_id" field.
func OryIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOryID, v))
}

// OryIDNEQ applies the NEQ predicate on the "ory_id" field.
func OryIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOryID, v))
}

// OryIDIn applies the In predicate on the "ory_id" field.
func OryIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldOryID, vs...))
}

// OryIDNotIn applies the NotIn predicate on the "ory_id" field.
func OryIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOryID, vs...))
}

// OryIDGT applies the GT predicate on the "ory_id" field.
func OryIDGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldOryID, v))
}

// OryIDGTE applies the GTE predicate on the "ory_id" field.
func OryIDGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOryID, v))
}

// OryIDLT applies the LT predicate on the "ory_id" field.
func OryIDLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldOryID, v))
}

// OryIDLTE applies the LTE predicate on the "ory_id" field.
func OryIDLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOryID, v))
}

// UILanguageEQ applies the EQ predicate on the "ui_language" field.
func UILanguageEQ(v UILanguage) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUILanguage, v))
}

// UILanguageNEQ applies the NEQ predicate on the "ui_language" field.
func UILanguageNEQ(v UILanguage) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUILanguage, v))
}

// UILanguageIn applies the In predicate on the "ui_language" field.
func UILanguageIn(vs ...UILanguage) predicate.User {
	return predicate.User(sql.FieldIn(FieldUILanguage, vs...))
}

// UILanguageNotIn applies the NotIn predicate on the "ui_language" field.
func UILanguageNotIn(vs ...UILanguage) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUILanguage, vs...))
}

// AdminEQ applies the EQ predicate on the "admin" field.
func AdminEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAdmin, v))
}

// AdminNEQ applies the NEQ predicate on the "admin" field.
func AdminNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAdmin, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
