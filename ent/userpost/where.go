// Code generated by ent, DO NOT EDIT.

package userpost

import (
	"movie-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPost {
	return predicate.UserPost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPost {
	return predicate.UserPost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPost {
	return predicate.UserPost(sql.FieldLTE(FieldID, id))
}

// PostText applies equality check predicate on the "post_text" field. It's identical to PostTextEQ.
func PostText(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldEQ(FieldPostText, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldEQ(FieldCreateTime, v))
}

// PostTextEQ applies the EQ predicate on the "post_text" field.
func PostTextEQ(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldEQ(FieldPostText, v))
}

// PostTextNEQ applies the NEQ predicate on the "post_text" field.
func PostTextNEQ(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldNEQ(FieldPostText, v))
}

// PostTextIn applies the In predicate on the "post_text" field.
func PostTextIn(vs ...string) predicate.UserPost {
	return predicate.UserPost(sql.FieldIn(FieldPostText, vs...))
}

// PostTextNotIn applies the NotIn predicate on the "post_text" field.
func PostTextNotIn(vs ...string) predicate.UserPost {
	return predicate.UserPost(sql.FieldNotIn(FieldPostText, vs...))
}

// PostTextGT applies the GT predicate on the "post_text" field.
func PostTextGT(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldGT(FieldPostText, v))
}

// PostTextGTE applies the GTE predicate on the "post_text" field.
func PostTextGTE(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldGTE(FieldPostText, v))
}

// PostTextLT applies the LT predicate on the "post_text" field.
func PostTextLT(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldLT(FieldPostText, v))
}

// PostTextLTE applies the LTE predicate on the "post_text" field.
func PostTextLTE(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldLTE(FieldPostText, v))
}

// PostTextContains applies the Contains predicate on the "post_text" field.
func PostTextContains(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldContains(FieldPostText, v))
}

// PostTextHasPrefix applies the HasPrefix predicate on the "post_text" field.
func PostTextHasPrefix(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldHasPrefix(FieldPostText, v))
}

// PostTextHasSuffix applies the HasSuffix predicate on the "post_text" field.
func PostTextHasSuffix(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldHasSuffix(FieldPostText, v))
}

// PostTextEqualFold applies the EqualFold predicate on the "post_text" field.
func PostTextEqualFold(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldEqualFold(FieldPostText, v))
}

// PostTextContainsFold applies the ContainsFold predicate on the "post_text" field.
func PostTextContainsFold(v string) predicate.UserPost {
	return predicate.UserPost(sql.FieldContainsFold(FieldPostText, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserPost {
	return predicate.UserPost(sql.FieldLTE(FieldCreateTime, v))
}

// HasUserID applies the HasEdge predicate on the "user_id" edge.
func HasUserID() predicate.UserPost {
	return predicate.UserPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIDWith applies the HasEdge predicate on the "user_id" edge with a given conditions (other predicates).
func HasUserIDWith(preds ...predicate.User) predicate.UserPost {
	return predicate.UserPost(func(s *sql.Selector) {
		step := newUserIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPost) predicate.UserPost {
	return predicate.UserPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPost) predicate.UserPost {
	return predicate.UserPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPost) predicate.UserPost {
	return predicate.UserPost(func(s *sql.Selector) {
		p(s.Not())
	})
}